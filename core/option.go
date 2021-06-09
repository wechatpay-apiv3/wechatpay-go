package core

import (
	"crypto/rsa"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/certificate_map"
)

// ClientOption  一个ClientOption可以作为微信支付api v3 client的配置
type ClientOption interface {
	Apply(settings *dialSettings)
}

type withCredentialOption struct{ credential auth.Credential }

// Apply 将配置添加到DialSettings中
func (w withCredentialOption) Apply(o *dialSettings) {
	o.Credential = w.credential
}

// WithCredential 返回一个指定credential的ClientOption，用于生成http request header 中authorization信息
func WithCredential(credential auth.Credential) ClientOption {
	return withCredentialOption{credential: credential}
}

// WithMerchantCredential 通过商户号、商户证书序列号、商户私钥构建一个默认的 Credential，用于生成请求头中的 Authorization 信息
func WithMerchantCredential(mchID, certificateSerialNo string, privateKey *rsa.PrivateKey) ClientOption {
	credential := &credentials.WechatPayCredentials{
		Signer: &signers.SHA256WithRSASigner{PrivateKey: privateKey, CertificateSerialNo: certificateSerialNo},
		MchID:  mchID,
	}
	return WithCredential(credential)
}

type withValidatorOption struct{ Validator auth.Validator }

// Apply 将配置添加到DialSettings中
func (w withValidatorOption) Apply(o *dialSettings) {
	o.Validator = w.Validator
}

// WithValidator 返回一个指定validator的ClientOption，用于校验http response header
func WithValidator(validator auth.Validator) ClientOption {
	return withValidatorOption{validator}
}

// WithWechatPayValidator 设置微信支付平台证书信息，返回一个指定validator的ClientOption，用于校验http response header
func WithWechatPayValidator(certificateList []*x509.Certificate) ClientOption {
	validator := &validators.WechatPayValidator{
		Verifier: verifiers.NewSHA256WithRSAVerifier(certificate_map.NewCertificateMapWithList(certificateList)),
	}
	return withValidatorOption{validator}
}

// WithoutValidator 返回一个指定validator的ClientOption，不进行验签 用于下载证书和下载账单等不需要进行验签的接口
func WithoutValidator() ClientOption {
	return withValidatorOption{&validators.NullValidator{}}
}

type withHTTPClientOption struct{ client *http.Client }

// Apply 将配置添加到DialSettings中
func (w withHTTPClientOption) Apply(o *dialSettings) {
	o.HTTPClient = w.client
}

// WithHTTPClient 返回一个指定网络通信为HttpClient的ClientOption，指定后使用用户自动创建的的http.client，如果用户不创建，则帮助用户
// 创建一个默认的http.client
func WithHTTPClient(client *http.Client) ClientOption {
	return withHTTPClientOption{client}
}

type withTimeoutOption time.Duration

// Apply 将配置添加到DialSettings中
func (w withTimeoutOption) Apply(o *dialSettings) {
	o.Timeout = time.Duration(w)
}

// WithTimeout 返回一个指定超时时间的ClientOption
func WithTimeout(timeout time.Duration) ClientOption {
	return withTimeoutOption(timeout)
}

type withHeaderOption struct{ header http.Header }

// Apply 将配置添加到DialSettings中
func (w withHeaderOption) Apply(o *dialSettings) {
	o.Header = w.header
}

// WithHeader 返回一个为http client设置额外header信息的ClientOption
func WithHeader(header http.Header) ClientOption {
	return withHeaderOption{header: header}
}
