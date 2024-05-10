package verifiers

import (
	"context"
	"crypto/rsa"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

// SHA256WithRSACombinedVerifier 数字签名验证器，组合了公钥和平台证书
type SHA256WithRSACombinedVerifier struct {
	publicKeyVerifier SHA256WithRSAPubkeyVerifier
	certVerifier      SHA256WithRSAVerifier
}

// Verify 验证签名，如果序列号和公钥一致则使用公钥验签，否则使用平台证书验签
func (v *SHA256WithRSACombinedVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	if serialNumber == v.publicKeyVerifier.keyID {
		return v.publicKeyVerifier.Verify(ctx, serialNumber, message, signature)
	}
	return v.certVerifier.Verify(ctx, serialNumber, message, signature)
}

// GetSerial 获取可验签的公钥序列号。该验签器只用在回调，所以获取序列号时返回错误
func (v *SHA256WithRSACombinedVerifier) GetSerial(ctx context.Context) (string, error) {
	return v.publicKeyVerifier.keyID, nil
}

// NewSHA256WithRSACombinedVerifier 用公钥和平台证书初始化验证器
func NewSHA256WithRSACombinedVerifier(
	getter core.CertificateGetter,
	keyID string,
	publicKey rsa.PublicKey) *SHA256WithRSACombinedVerifier {
	return &SHA256WithRSACombinedVerifier{
		*NewSHA256WithRSAPubkeyVerifier(keyID, publicKey),
		*NewSHA256WithRSAVerifier(getter),
	}
}
