//Package auth 微信支付api v3 go http-client 签名校验相关接口
package auth

import "context"

// SignatureResult 签名结果
type SignatureResult struct {
	MchCertificateSerialNo string // 商户序列号
	Signature string // 签名
}

// Signer 签名生成器
type Signer interface {
	GetName() string                                          // 获取签名器的名称
	GetType() string                                          // 获取签名器的类型
	GetVersion() string                                       // 获取签名器的版本
	Sign(ctx context.Context, message string) (*SignatureResult, error) // 对信息进行签名
}
