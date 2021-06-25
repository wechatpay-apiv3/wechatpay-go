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
		return serial, c.cipher(ctx, cipherTypeEncrypt, v)
	} else {
		return serial, c.cipher(ctx, cipherTypeEncrypt, reflect.ValueOf(in))
	}
}

// Decrypt 对结构中的敏感字段进行解密
func (c *WechatPayCipher) Decrypt(ctx context.Context, in interface{}) error {
	if v, ok := in.(reflect.Value); ok {
		return c.cipher(ctx, cipherTypeDecrypt, v)
	} else {
		return c.cipher(ctx, cipherTypeDecrypt, reflect.ValueOf(in))
	}
}

// cipher 递归进行加密/解密操作
func (c *WechatPayCipher) cipher(ctx context.Context, ty cipherType, v reflect.Value) error {
	var t = v.Type()

	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}

		t = t.Elem()
		v = v.Elem()
	}

	// Only Struct can be ciphered
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("only struct can be ciphered")
	}

	// if not settable, cannot do further process
	if !v.CanSet() {
		return fmt.Errorf("in-place cipher requires settable input, ptr for example")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		fieldValue := v.Field(i)

		if !fieldValue.CanInterface() {
			// ignore unexported fields
			return nil
		}

		if fieldType.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				continue
			}

			fieldType = fieldType.Elem()
			fieldValue = fieldValue.Elem()
		}

		if fieldType.Kind() == reflect.Array || fieldType.Kind() == reflect.Slice {
			elemType := fieldType.Elem()

			if c.isFieldRequireCipher(field, elemType) {
				for j := 0; j < fieldValue.Len(); j++ {
					elemValue := fieldValue.Index(j)
					if err := c.cipherField(ctx, ty, field, elemType, elemValue); err != nil {
						return err
					}
				}
			}
		} else if c.isFieldRequireCipher(field, fieldType) {
			if err := c.cipherField(ctx, ty, field, fieldType, fieldValue); err != nil {
				return err
			}
		}
	}

	return nil
}

// cipherField 对字段进行加密/解密操作，无需操作则跳过
func (c *WechatPayCipher) cipherField(ctx context.Context, ty cipherType, f reflect.StructField, t reflect.Type, v reflect.Value) error {
	if t.Kind() == reflect.Struct {
		if err := c.cipher(ctx, ty, v); err != nil {
			return err
		}
	} else if t.Kind() == reflect.String && f.Tag.Get(fieldTagEncryption) == encryptionTypeAPIV3 {
		var cipherText string
		var err error

		switch ty {
		case cipherTypeEncrypt:
			serial, ok := getEncryptSerial(ctx)
			if !ok {
				return fmt.Errorf("`getEncryptSerial` not provided in ctx")
			}
			cipherText, err = c.encryptor.Encrypt(ctx, serial, v.Interface().(string))
		case cipherTypeDecrypt:
			cipherText, err = c.decryptor.Decrypt(ctx, v.Interface().(string))
		default:
			return fmt.Errorf("invalid cipher type:%v", ty)
		}

		if err != nil {
			return err
		}
		v.SetString(cipherText)
	}

	return nil
}

// isFieldRequireCipher 判断该字段是否需要加密/解密
func (c *WechatPayCipher) isFieldRequireCipher(f reflect.StructField, t reflect.Type) bool {
	if t.Kind() == reflect.Struct {
		return true
	}

	if t.Kind() == reflect.String && f.Tag.Get(fieldTagEncryption) == encryptionTypeAPIV3 {
		return true
	}

	return false
}

func NewWechatPayCipher(encryptor cipher.Encryptor, decryptor cipher.Decryptor) *WechatPayCipher {
	return &WechatPayCipher{encryptor: encryptor, decryptor: decryptor}
}
