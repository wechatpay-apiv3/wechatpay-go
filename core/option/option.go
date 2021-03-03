//Package option 微信支付api v3 go http-client Option选项
package option

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"math/big"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/setting"
)

// ClientOption  一个ClientOption可以作为微信支付api v3 client的配置
type ClientOption interface {
	Apply(settings *setting.DialSettings)
}


// WithMerchant 通过商户号、商户证书序列号、私钥构建一个默认的credential的ClientOption，用于生成http request header 中authorization信息
func WithMerchant(mchID, certificateSerialNo string, privateKey *rsa.PrivateKey) ClientOption{
	credential := &credentials.WechatPayCredentials{
		Signer:              &signers.Sha256WithRSASigner{PrivateKey: privateKey},
		MchID:               mchID,
		CertificateSerialNo: certificateSerialNo,
	}
	return withCredential{credential: credential}
}


// WithCredential 返回一个指定credential的ClientOption，用于生成http request header 中authorization信息
func WithCredential(credential auth.Credential) ClientOption {
	return withCredential{credential: credential}
}

type withCredential struct{ credential auth.Credential }

// Apply 将配置添加到DialSettings中
func (w withCredential) Apply(o *setting.DialSettings) {
	o.Credential = w.credential
}

// WithWechatPay 设置微信支付平台证书信息，返回一个指定validator的ClientOption，用于校验http response header
func WithWechatPay(certificateList []*x509.Certificate) ClientOption{
	certificates := map[string]*x509.Certificate{}
	for _, certificate := range certificateList {
		if serialNumberBigInt, ok := new(big.Int).SetString(certificate.SerialNumber.String(), 0); ok {
			serialNumberStr := fmt.Sprintf("%X", serialNumberBigInt)
			certificates[serialNumberStr] = certificate
		}
	}
	validator := &validators.WechatPayValidator{
		Verifier: &verifiers.WechatPayVerifier{
			Certificates: certificates,
		},
	}
	return withValidator{validator}
}

// WithCredential 返回一个指定validator的ClientOption，用于校验http response header
func WithValidator(validator auth.Validator) ClientOption {
	return withValidator{validator}
}

type withValidator struct{ Validator auth.Validator }

// Apply 将配置添加到DialSettings中
func (w withValidator) Apply(o *setting.DialSettings) {
	o.Validator = w.Validator
}

// WithHTTPClient 返回一个指定网络通信为HttpClient的ClientOption，指定后使用用户自动创建的的http.client，如果用户不创建，则帮助用户
// 创建一个默认的http.client
func WithHTTPClient(client *http.Client) ClientOption {
	return withHTTPClient{client}
}

type withHTTPClient struct{ client *http.Client }

// Apply 将配置添加到DialSettings中
func (w withHTTPClient) Apply(o *setting.DialSettings) {
	o.HTTPClient = w.client
}

// WithTimeout 返回一个指定超时时间的ClientOption
func WithTimeout(timeout time.Duration) ClientOption {
	return withTimeout(timeout)
}

type withTimeout time.Duration

// Apply 将配置添加到DialSettings中
func (w withTimeout) Apply(o *setting.DialSettings) {
	o.Timeout = time.Duration(w)
}
