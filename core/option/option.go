package option

import (
	"crypto/rsa"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/ciphers"
)

// WithSigner 返回一个指定signer的ClientOption
func WithSigner(signer auth.Signer) core.ClientOption {
	return core.WithSignerOption{Signer: signer}
}

// WithMerchantCredential 通过商户号、商户证书序列号、商户私钥构建一对 Credential/Signer，用于生成请求头中的 Authorization 信息
func WithMerchantCredential(mchID, certificateSerialNo string, privateKey *rsa.PrivateKey) core.ClientOption {
	signer := &signers.SHA256WithRSASigner{
		MchID:               mchID,
		PrivateKey:          privateKey,
		CertificateSerialNo: certificateSerialNo,
	}
	return WithSigner(signer)
}

// WithVerifier 返回一个指定verifier的ClientOption，用于校验http response header
func WithVerifier(verifier auth.Verifier) core.ClientOption {
	validator := &validators.WechatPayValidator{
		Verifier: verifier,
	}
	return core.WithValidatorOption{Validator: validator}
}

// WithWechatPayCertificate 设置微信支付平台证书信息，返回一个指定validator的ClientOption，用于校验http response header
func WithWechatPayCertificate(certificateList []*x509.Certificate) core.ClientOption {
	verifier := verifiers.NewSHA256WithRSAVerifier(cert.NewCertificateMapWithList(certificateList))
	return WithVerifier(verifier)
}

// WithoutValidator 返回一个指定validator的ClientOption，不进行验签 用于下载证书和下载账单等不需要进行验签的接口
func WithoutValidator() core.ClientOption {
	return core.WithValidatorOption{Validator: &validators.NullValidator{}}
}

// WithHTTPClient 返回一个指定网络通信为HttpClient的ClientOption，指定后使用用户自动创建的的http.client，如果用户不创建，则帮助用户
// 创建一个默认的http.client
func WithHTTPClient(client *http.Client) core.ClientOption {
	return core.WithHTTPClientOption{Client: client}
}

// WithTimeout 返回一个指定超时时间的ClientOption
func WithTimeout(timeout time.Duration) core.ClientOption {
	return core.WithTimeoutOption(timeout)
}

// WithHeader 返回一个为http client设置额外header信息的ClientOption
func WithHeader(header http.Header) core.ClientOption {
	return core.WithHeaderOption{Header: header}
}

// WithWechatPayCipher 返回一个为 Client 设置 WechatPayCipher 的 core.ClientOption
func WithWechatPayCipher(encryptor cipher.Encryptor, decryptor cipher.Decryptor) core.ClientOption {
	return core.WithCipherOption{Cipher: ciphers.NewWechatPayCipher(encryptor, decryptor)}
}
