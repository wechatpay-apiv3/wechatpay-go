package builder

import (
	"crypto/x509"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/encryptors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
)

func BuildWechatPayEncryptorWithCertificateList(l []*x509.Certificate) *encryptors.WechatPayEncryptor {
	return encryptors.NewWechatPayEncryptor(core.NewCertificateMapWithList(l))
}

func BuildWechatPayEncryptorWithDownloaderMgr(
	mgr *downloader.CertificateDownloaderMgr, mchID string,
) *encryptors.WechatPayEncryptor {
	return encryptors.NewWechatPayEncryptor(mgr.GetCertificateVisitor(mchID))
}
