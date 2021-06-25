package option

import (
	"context"
	"crypto/rsa"
	"crypto/x509"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/ciphers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/decryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/encryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
)

type withAuthCipherOption struct{ settings core.DialSettings }

func (w withAuthCipherOption) Apply(o *core.DialSettings) error {
	o.Signer = w.settings.Signer
	o.Validator = w.settings.Validator
	o.Cipher = w.settings.Cipher
	return nil
}

// WithWechatPayAuthCipher 一键初始化 Client，使其具备「签名/验签/敏感字段加解密」能力
func WithWechatPayAuthCipher(
	mchID string, certificateSerialNo string, privateKey *rsa.PrivateKey, certificateList []*x509.Certificate,
) core.ClientOption {
	certGetter := core.NewCertificateMapWithList(certificateList)
	return withAuthCipherOption{
		settings: core.DialSettings{
			Signer: &signers.SHA256WithRSASigner{
				MchID:               mchID,
				PrivateKey:          privateKey,
				CertificateSerialNo: certificateSerialNo,
			},
			Validator: validators.NewWechatPayResponseValidator(verifiers.NewSHA256WithRSAVerifier(certGetter)),
			Cipher: ciphers.NewWechatPayCipher(
				encryptors.NewWechatPayEncryptor(certGetter),
				decryptors.NewWechatPayDecryptor(privateKey),
			),
		},
	}
}

// WithWechatPayAutoAuthCipher 一键初始化 Client，使其具备「签名/验签/敏感字段加解密」能力。
// 同时提供证书定时更新功能（因此需要提供 mchAPIv3Key 用于证书解密），不再需要本地提供平台证书
func WithWechatPayAutoAuthCipher(
	mchID string, certificateSerialNo string, privateKey *rsa.PrivateKey, mchAPIv3Key string,
) core.ClientOption {
	mgr := downloader.MgrInstance()

	if !mgr.HasDownloader(context.Background(), mchID) {
		err := mgr.RegisterDownloaderWithPrivateKey(
			context.Background(), privateKey, certificateSerialNo, mchID, mchAPIv3Key,
		)
		if err != nil {
			return core.ErrorOption{Error: err}
		}
	}

	return WithWechatPayAutoAuthCipherUsingDownloaderMgr(mchID, certificateSerialNo, privateKey, mgr)
}

// WithWechatPayAutoAuthCipherUsingDownloaderMgr 一键初始化 Client，使其具备「签名/验签/敏感字段加解密」能力。
// 需要使用者自行提供 CertificateDownloaderMgr 已实现平台证书的自动更新
//
// 【注意】本函数的能力与 WithWechatPayAutoAuthCipher 完全一致，除非有自行管理 CertificateDownloaderMgr 的需求，
// 否则推荐直接使用 WithWechatPayAutoAuthCipher
func WithWechatPayAutoAuthCipherUsingDownloaderMgr(
	mchID string, certificateSerialNo string, privateKey *rsa.PrivateKey, mgr *downloader.CertificateDownloaderMgr,
) core.ClientOption {
	certVisitor := mgr.GetCertificateVisitor(mchID)
	return withAuthCipherOption{
		settings: core.DialSettings{
			Signer: &signers.SHA256WithRSASigner{
				MchID:               mchID,
				CertificateSerialNo: certificateSerialNo,
				PrivateKey:          privateKey,
			},
			Validator: validators.NewWechatPayResponseValidator(verifiers.NewSHA256WithRSAVerifier(certVisitor)),
			Cipher: ciphers.NewWechatPayCipher(
				encryptors.NewWechatPayEncryptor(certVisitor),
				decryptors.NewWechatPayDecryptor(privateKey),
			),
		},
	}
}
