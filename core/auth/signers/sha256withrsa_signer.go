//Package signers 微信支付api v3 go http-client 签名生成器
package signers

import (
	"context"
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
)

// Sha256WithRSASigner Sha256WithRSA 签名器
type Sha256WithRSASigner struct {
	MchCertificateSerialNo string          // 商户证书序列号
	PrivateKey             *rsa.PrivateKey // 商户私钥
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
func (s *Sha256WithRSASigner) Sign(ctx context.Context, message string) (*auth.SignatureResult, error) {
	if s.PrivateKey == nil {
		return nil, fmt.Errorf("you must set privatekey to use Sha256WithRSASigner")
	}
	if strings.TrimSpace(s.MchCertificateSerialNo) == "" {
		return nil, fmt.Errorf("you must set mch certificate serial no to use Sha256WithRSASigner")
	}
	signature, err := Sha256WithRsa(message, s.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &auth.SignatureResult{MchCertificateSerialNo: s.MchCertificateSerialNo, Signature: signature}, nil
}
