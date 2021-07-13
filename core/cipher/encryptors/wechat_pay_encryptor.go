// Copyright 2021 Tencent Inc. All rights reserved.

package encryptors

import (
	"context"
	"fmt"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// WechatPayEncryptor 微信支付字符串加密器
type WechatPayEncryptor struct {
	// 微信支付平台证书提供器
	certGetter core.CertificateGetter
}

// NewWechatPayEncryptor 新建一个 WechatPayEncryptor
func NewWechatPayEncryptor(certProvider core.CertificateGetter) *WechatPayEncryptor {
	return &WechatPayEncryptor{certGetter: certProvider}
}

// SelectCertificate 选择合适的微信支付平台证书用于加密
// 根据微信支付平台证书文档说明，应优先使用最新的证书（即启用时间最晚）
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/jie-kou-wen-dang/ping-tai-zheng-shu#zhu-yi-shi-xiang
func (e *WechatPayEncryptor) SelectCertificate(ctx context.Context) (serial string, err error) {
	newestSerial := e.certGetter.GetNewestSerial(ctx)
	if newestSerial == "" {
		return "", fmt.Errorf("no certificate for encryption")
	}

	return newestSerial, nil
}

// Encrypt 对字符串加密
func (e *WechatPayEncryptor) Encrypt(ctx context.Context, serial, plaintext string) (ciphertext string, err error) {
	cert, ok := e.certGetter.Get(ctx, serial)

	if !ok {
		return plaintext, fmt.Errorf("cert for EncryptSerial(%v) not found", serial)
	}

	// 不需要对空串进行加密
	if plaintext == "" {
		return "", nil
	}

	return utils.EncryptOAEPWithCertificate(plaintext, cert)
}
