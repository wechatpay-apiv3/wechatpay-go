package builder

import (
	"crypto/x509"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher/encryptors"
)

func BuildWechatPayEncryptorWithCertificateList(l []*x509.Certificate) *encryptors.WechatPayEncryptor {
	return encryptors.NewWechatPayEncryptor(core.NewCertificateMapWithList(l))
}

func BuildWechatPayEncryptorWithDownloaderMgr(
	mgr *cert.CertificateDownloaderMgr, mchID string,
) *encryptors.WechatPayEncryptor {
	return encryptors.NewWechatPayEncryptor(mgr.GetCertificateVisitor(mchID))
}
