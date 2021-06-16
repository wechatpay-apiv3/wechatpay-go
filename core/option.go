package core

import (
	"crypto/rsa"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/ciphers"
)

// ClientOption 微信支付 API v3 HTTPClient core.Client 初始化参数
type ClientOption interface {
	Apply(settings *DialSettings) error
}

// ErrorOption 错误初始化参数，用于返回错误
type ErrorOption struct{ Error error }

// Apply 返回初始化错误
func (w ErrorOption) Apply(o *DialSettings) error {
	return w.Error
}

type withSignerOption struct {
	signer auth.Signer
}

func (w withSignerOption) Apply(o *DialSettings) error {
	o.Signer = w.signer
	return nil
}

// WithSigner 返回一个指定signer的ClientOption
func WithSigner(signer auth.Signer) ClientOption {
	return withSignerOption{signer}
}

// WithMerchantCredential 通过商户号、商户证书序列号、商户私钥构建一对 Credential/Signer，用于生成请求头中的 Authorization 信息
func WithMerchantCredential(mchID, certificateSerialNo string, privateKey *rsa.PrivateKey) ClientOption {
	signer := &signers.SHA256WithRSASigner{
		MchID:               mchID,
		PrivateKey:          privateKey,
		CertificateSerialNo: certificateSerialNo,
	}
	return WithSigner(signer)
}

type withValidatorOption struct{ Validator auth.Validator }

// Apply 将配置添加到DialSettings中
func (w withValidatorOption) Apply(o *DialSettings) error {
	o.Validator = w.Validator
	return nil
}

// WithVerifier 返回一个指定verifier的ClientOption，用于校验http response header
func WithVerifier(verifier auth.Verifier) ClientOption {
	validator := &validators.WechatPayValidator{
		Verifier: verifier,
	}
	return withValidatorOption{validator}
}

// WithWechatPayCertificate 设置微信支付平台证书信息，返回一个指定validator的ClientOption，用于校验http response header
func WithWechatPayCertificate(certificateList []*x509.Certificate) ClientOption {
	verifier := verifiers.NewSHA256WithRSAVerifier(cert.NewCertificateMapWithList(certificateList))
	return WithVerifier(verifier)
}

// WithoutValidator 返回一个指定validator的ClientOption，不进行验签 用于下载证书和下载账单等不需要进行验签的接口
func WithoutValidator() ClientOption {
	return withValidatorOption{&validators.NullValidator{}}
}

type withHTTPClientOption struct{ client *http.Client }

// Apply 将配置添加到DialSettings中
func (w withHTTPClientOption) Apply(o *DialSettings) error {
	o.HTTPClient = w.client
	return nil
}

// WithHTTPClient 返回一个指定网络通信为HttpClient的ClientOption，指定后使用用户自动创建的的http.client，如果用户不创建，则帮助用户
// 创建一个默认的http.client
func WithHTTPClient(client *http.Client) ClientOption {
	return withHTTPClientOption{client}
}

type withTimeoutOption time.Duration

// Apply 将配置添加到DialSettings中
func (w withTimeoutOption) Apply(o *DialSettings) error {
	o.Timeout = time.Duration(w)
	return nil
}

// WithTimeout 返回一个指定超时时间的ClientOption
func WithTimeout(timeout time.Duration) ClientOption {
	return withTimeoutOption(timeout)
}

type withHeaderOption struct{ header http.Header }

// Apply 将配置添加到DialSettings中
func (w withHeaderOption) Apply(o *DialSettings) error {
	o.Header = w.header
	return nil
}

// WithHeader 返回一个为http client设置额外header信息的ClientOption
func WithHeader(header http.Header) ClientOption {
	return withHeaderOption{header: header}
}

type withCipherOption struct{ cipher cipher.Cipher }

// Apply 将配置添加到DialSettings中
func (w withCipherOption) Apply(o *DialSettings) error {
	o.Cipher = w.cipher
	return nil
}

// WithWechatPayCipher 返回一个为 Client 设置 WechatPayCipher 的 ClientOption
func WithWechatPayCipher(encryptor cipher.Encryptor, decryptor cipher.Decryptor) ClientOption {
	return withCipherOption{ciphers.NewWechatPayCipher(encryptor, decryptor)}
}
