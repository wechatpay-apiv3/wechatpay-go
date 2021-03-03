//Package signers 微信支付api v3 go http-client 签名生成器
package signers

import (
	"context"
	"crypto/rsa"
	"fmt"
)

// Sha256WithRSASigner Sha256WithRSA 签名器
type Sha256WithRSASigner struct {
	PrivateKey *rsa.PrivateKey // 私钥
}

// GetName 获取签名器的名称
func (s *Sha256WithRSASigner) GetName() string {
	return "SHA256withRSA"
}

// 获取签名器的类型
func (s *Sha256WithRSASigner) GetType() string {
	return "PRIVATEKEY"
}

// 获取签名器的版本
func (s *Sha256WithRSASigner) GetVersion() string {
	return "1.0"
}

// 对信息使用Sha256WithRsa的方式进行签名
func (s *Sha256WithRSASigner) Sign(ctx context.Context, message string) (signature string, err error) {
	if s.PrivateKey == nil {
		return "", fmt.Errorf("you must set privatekey to use Sha256WithRSASigner")
	}
	return Sha256WithRsa(message, s.PrivateKey)
}
