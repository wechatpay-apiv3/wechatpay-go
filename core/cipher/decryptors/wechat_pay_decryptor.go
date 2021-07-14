// Copyright 2021 Tencent Inc. All rights reserved.

package decryptors

import (
	"context"
	"crypto/rsa"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// WechatPayDecryptor 微信支付字符串解密器
type WechatPayDecryptor struct {
	// 商户私钥
	privateKey *rsa.PrivateKey
}

// Decrypt 使用商户私钥对字符串进行解密
func (d *WechatPayDecryptor) Decrypt(_ context.Context, ciphertext string) (plaintext string, err error) {
	if ciphertext == "" {
		return "", nil
	}
	return utils.DecryptOAEP(ciphertext, d.privateKey)
}

// NewWechatPayDecryptor 使用商户私钥初始化一个 WechatPayDecryptor
func NewWechatPayDecryptor(privateKey *rsa.PrivateKey) *WechatPayDecryptor {
	return &WechatPayDecryptor{privateKey: privateKey}
}
