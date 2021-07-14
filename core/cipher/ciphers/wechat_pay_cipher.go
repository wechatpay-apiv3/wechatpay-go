// Copyright 2021 Tencent Inc. All rights reserved.

package ciphers

import (
	"context"
	"fmt"
	"reflect"

	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher"
)

type cipherType string

const (
	cipherTypeEncrypt cipherType = "encrypt"
	cipherTypeDecrypt cipherType = "decrypt"
)

const (
	fieldTagEncryption  = "encryption"
	encryptionTypeAPIV3 = "EM_APIV3"
)

// fieldCipherFuncType 用于对特定类型字段进行加/解密的方法类型
type fieldCipherFuncType func(*WechatPayCipher, context.Context, cipherType, reflect.StructField, reflect.Value) error

// fieldCipherMap 对不同类型字段进行加/解密的方法字典
var fieldCipherMap map[reflect.Kind]fieldCipherFuncType

func init() {
	// 初始化加/解密方法字典，使用 init 初始化而不是直接在声明时初始化的原因是为了避免初始化循环依赖
	fieldCipherMap = map[reflect.Kind]fieldCipherFuncType{
		reflect.Struct: (*WechatPayCipher).cipherStructField,
		reflect.Array:  (*WechatPayCipher).cipherArrayField,
		reflect.Slice:  (*WechatPayCipher).cipherArrayField,
		reflect.String: (*WechatPayCipher).cipherStringField,
	}
}

// WechatPayCipher 提供微信支付敏感信息加解密功能
//
// 为了保证通信过程中敏感信息字段（如用户的住址、银行卡号、手机号码等）的机密性，微信支付API v3要求：
//  1. 商户对上送的敏感信息字段进行加密
//  2. 微信支付对下行的敏感信息字段进行加密
// 详见：https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/min-gan-xin-xi-jia-mi
type WechatPayCipher struct {
	encryptor cipher.Encryptor
	decryptor cipher.Decryptor
}

// Encrypt 对结构中的敏感字段进行加密
func (c *WechatPayCipher) Encrypt(ctx context.Context, in interface{}) (string, error) {
	serial, err := c.encryptor.SelectCertificate(ctx)
	if err != nil {
		return "", err
	}

	ctx = setEncryptSerial(ctx, serial)
	if v, ok := in.(reflect.Value); ok {
		err = c.cipher(ctx, cipherTypeEncrypt, v)
	} else {
		err = c.cipher(ctx, cipherTypeEncrypt, reflect.ValueOf(in))
	}
	if err != nil {
		return "", fmt.Errorf("encrypt struct failed: %w", err)
	}

	return serial, nil
}

// Decrypt 对结构中的敏感字段进行解密
func (c *WechatPayCipher) Decrypt(ctx context.Context, in interface{}) error {
	var err error
	if v, ok := in.(reflect.Value); ok {
		err = c.cipher(ctx, cipherTypeDecrypt, v)
	} else {
		err = c.cipher(ctx, cipherTypeDecrypt, reflect.ValueOf(in))
	}

	if err != nil {
		return fmt.Errorf("decrypt struct failed: %w", err)
	}

	return nil
}

// cipher 执行加/解密的入口函数
func (c *WechatPayCipher) cipher(ctx context.Context, ty cipherType, v reflect.Value) error {
	var isNil bool
	if v, isNil = derefPtrValue(v); isNil {
		// No cipher required for nil ptr
		return nil
	}

	if !v.CanSet() {
		return fmt.Errorf("in-place cipher requires settable input, ptr for example")
	}

	if v.Type().Kind() != reflect.Struct {
		return fmt.Errorf("only struct can be ciphered")
	}

	return c.cipherStruct(ctx, ty, v)
}

// cipherStruct 递归进行Struct的加/解密操作
func (c *WechatPayCipher) cipherStruct(ctx context.Context, ty cipherType, v reflect.Value) error {
	var t = v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if err := c.cipherField(ctx, ty, field, fieldValue); err != nil {
			return err
		}
	}

	return nil
}

// derefPtrValue 将 Ptr 类型的 Value 解引用，直到获得非指针内容。
//
// 如果输入的 Value 不是 Ptr，则返回其本身
// 如果输入的 Value 最终指向 Nil，则返回最终指向 Nil 的 Value 对象，且 isNil 为 true
func derefPtrValue(inValue reflect.Value) (outValue reflect.Value, isNil bool) {
	v := inValue

	for v.Type().Kind() == reflect.Ptr {
		if v.IsNil() {
			return v, true
		}
		v = v.Elem()
	}

	return v, false
}

// cipherField 对字段进行加/解密
func (c *WechatPayCipher) cipherField(
	ctx context.Context, ty cipherType, field reflect.StructField, fieldValue reflect.Value,
) error {
	if !fieldValue.CanInterface() {
		// Skip Unexported Field
		return nil
	}

	var isNil bool
	if fieldValue, isNil = derefPtrValue(fieldValue); isNil {
		// Skip Field with no data
		return nil
	}

	if fieldCipherFunc, ok := fieldCipherMap[fieldValue.Type().Kind()]; ok {
		return fieldCipherFunc(c, ctx, ty, field, fieldValue)
	}

	return nil
}

// cipherStructField 对Struct类型的字段进行加/解密
func (c *WechatPayCipher) cipherStructField(
	ctx context.Context, ty cipherType, field reflect.StructField, fieldValue reflect.Value,
) error {
	_ = field
	return c.cipherStruct(ctx, ty, fieldValue)
}

// cipherArrayField 对Array/Slice类型的字段进行加/解密
func (c *WechatPayCipher) cipherArrayField(
	ctx context.Context, ty cipherType, field reflect.StructField, fieldValue reflect.Value,
) error {
	elemType := fieldValue.Type().Elem()
	if _, ok := fieldCipherMap[elemType.Kind()]; !ok {
		// Field Element Type Requires no encryption, skip
		return nil
	}

	for j := 0; j < fieldValue.Len(); j++ {
		elemValue := fieldValue.Index(j)
		if err := c.cipherField(ctx, ty, field, elemValue); err != nil {
			return err
		}
	}
	return nil
}

// cipherStringField 对String类型的字段进行加/解密
func (c *WechatPayCipher) cipherStringField(
	ctx context.Context, ty cipherType, field reflect.StructField, fieldValue reflect.Value,
) error {
	if field.Tag.Get(fieldTagEncryption) != encryptionTypeAPIV3 {
		return nil
	}

	var cipherText string
	var err error

	switch ty {
	case cipherTypeEncrypt:
		serial, ok := getEncryptSerial(ctx)
		if !ok {
			// 前置逻辑已经设置了 EncryptSerial，这里正常来讲不会进入
			return fmt.Errorf("`%s` not provided in ctx(should not happen)", contextKeyEncryptSerial)
		}
		cipherText, err = c.encryptor.Encrypt(ctx, serial, fieldValue.Interface().(string))
	case cipherTypeDecrypt:
		cipherText, err = c.decryptor.Decrypt(ctx, fieldValue.Interface().(string))
	default:
		// 前置逻辑不会设置其他类型，这里正常来讲不会进入
		return fmt.Errorf("invalid cipher type:%v(should not happen)", ty)
	}

	if err != nil {
		return err
	}
	fieldValue.SetString(cipherText)
	return nil
}

// NewWechatPayCipher 使用 cipher.Encryptor + cipher.Decryptor 构建一个 WechatPayCipher
func NewWechatPayCipher(encryptor cipher.Encryptor, decryptor cipher.Decryptor) *WechatPayCipher {
	return &WechatPayCipher{encryptor: encryptor, decryptor: decryptor}
}
