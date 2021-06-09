package builder

import (
	"crypto/x509"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/certificate_providers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// BuildWechatPayValidator 使用 平台证书列表 初始化 WechatPayValidator
func BuildWechatPayValidator(certificateList []*x509.Certificate) *validators.WechatPayValidator {
	certificates := map[string]*x509.Certificate{}
	for _, certificate := range certificateList {
		serialNo := utils.GetCertificateSerialNumber(*certificate)
		certificates[serialNo] = certificate
	}
	validator := &validators.WechatPayValidator{
		Verifier: verifiers.NewSHA256WithRSAVerifier(certificate_providers.NewSimpleCertificateProvider(certificates)),
	}
	return validator
}

// BuildWechatPayValidatorWithCertProvider 使用 平台证书提供器 初始化 WechatPayValidator
func BuildWechatPayValidatorWithCertProvider(certProvider cert.CertificateProvider) *validators.WechatPayValidator {
	return &validators.WechatPayValidator{
		Verifier: verifiers.NewSHA256WithRSAVerifier(certProvider),
	}
}

// BuildWechatPayValidatorWithCertificateDownloaderMgr 使用平台证书下载管理器构建 WechatPayValidator
// 平台证书下载管理器会自动下载最新的平台证书，保证验签功能持续有效
func BuildWechatPayValidatorWithCertificateDownloaderMgr(
	mgr *downloader.CertificateDownloaderMgr,
	mchID string,
) *validators.WechatPayValidator {
	return BuildWechatPayValidatorWithCertProvider(mgr.GetCertificateVisitor(mchID))
}
