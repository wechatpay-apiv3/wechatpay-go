// Package auth 微信支付 API v3 Go SDK 安全验证相关接口
package auth

import "context"

// SignatureResult 数字签名结果
type SignatureResult struct {
	CertificateSerialNo string // 签名对应的证书序列号
	Signature           string // 签名内容
}

// Signer 数字签名生成器
type Signer interface {
	Sign(ctx context.Context, message string) (*SignatureResult, error) // 对信息进行签名
}
