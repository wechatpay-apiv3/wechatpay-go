// Copyright 2024 Tencent Inc. All rights reserved.

package encryptors

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// WechatPayPubKeyEncryptor 微信支付字符串加密器，使用微信支付公钥
type WechatPayPubKeyEncryptor struct {
	// 微信支付公钥
	publicKey rsa.PublicKey
	// 公钥 ID
	keyID string
}

// NewWechatPayPubKeyEncryptor 新建一个 WechatPayPubKeyEncryptor
func NewWechatPayPubKeyEncryptor(keyID string, publicKey rsa.PublicKey) *WechatPayPubKeyEncryptor {
	return &WechatPayPubKeyEncryptor{publicKey: publicKey, keyID: keyID}
}

// SelectCertificate 选择合适的微信支付平台证书用于加密
// 返回公钥对应的 KeyId 作为证书序列号
func (e *WechatPayPubKeyEncryptor) SelectCertificate(ctx context.Context) (serial string, err error) {
	return e.keyID, nil
}

// Encrypt 对字符串加密
func (e *WechatPayPubKeyEncryptor) Encrypt(ctx context.Context, serial, plaintext string) (ciphertext string, err error) {
	if serial != e.keyID {
		return "", fmt.Errorf("serial %v not match key-id %v", serial, e.keyID)
	}

	// 不需要对空串进行加密
	if plaintext == "" {
		return "", nil
	}

	return utils.EncryptOAEPWithPublicKey(plaintext, &e.publicKey)
}
