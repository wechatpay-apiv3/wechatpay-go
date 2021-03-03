//Package verifiers 微信支付api v3 go http-client 验证器
package verifiers

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"strings"
)

// WechatPayVerifier 微信支付验证器
type WechatPayVerifier struct {
	Certificates map[string]*x509.Certificate // key 微信支付平台证书序列号 value 微信支付平台证书 （需要通过下载证书接口获得）
}

// Verify  对回包中的签名信息进行验证，验证信息是否被篡改、劫持等
func (verifier *WechatPayVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	err := checkParameter(ctx, serialNumber, message, signature)
	if err != nil {
		return err
	}
	if verifier.Certificates == nil {
		return fmt.Errorf("there is no certificate in wechat pay verifier")
	}
	certificate, ok := verifier.Certificates[serialNumber]
	if !ok {
		return fmt.Errorf("no serial number:%s corresponding certificate ", serialNumber)
	}
	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(certificate.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], []byte(signature))
	if err != nil {
		return fmt.Errorf("verifty signature with public key err:%s", err.Error())
	}
	return nil
}

func checkParameter(ctx context.Context, serialNumber, message, signature string) error {
	if ctx == nil {
		return fmt.Errorf("context is nil, verifier need input context.Context")
	}
	if strings.TrimSpace(serialNumber) == "" {
		return fmt.Errorf("serialNumber is empty, verifier need input serialNumber")
	}
	if strings.TrimSpace(message) == "" {
		return fmt.Errorf("message is empty, verifier need input message")
	}
	if strings.TrimSpace(signature) == "" {
		return fmt.Errorf("signature is empty, verifier need input signature")
	}
	return nil
}
