package builder

import (
	"crypto/x509"

	"github.com/wechatpay-apiv3/wechatpay-go/core/cert"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/encryptors"
)

func BuildWechatPayEncryptorWithCertificateList(l []*x509.Certificate) *encryptors.WechatPayEncryptor {
	return encryptors.NewWechatPayEncryptor(cert.NewCertificateMapWithList(l))
}

func BuildWechatPayEncryptorWithDownloaderMgr(
	mgr *downloader.CertificateDownloaderMgr, mchID string,
) *encryptors.WechatPayEncryptor {
	return encryptors.NewWechatPayEncryptor(mgr.GetCertificateVisitor(mchID))
}
