package builder

import (
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/downloader"
)

// BuildVerifierWithCertProvider 使用 平台证书提供器 初始化 Signer
func BuildVerifierWithCertProvider(certProvider cert.CertificateGetter) auth.Verifier {
	return verifiers.NewSHA256WithRSAVerifier(certProvider)
}

// BuildVerifierWithCertificateDownloaderMgr 使用平台证书下载管理器构建 WechatPayValidator
// 平台证书下载管理器会自动下载最新的平台证书，保证验签功能持续有效
func BuildVerifierWithCertificateDownloaderMgr(
	mgr *downloader.CertificateDownloaderMgr,
	mchID string,
) auth.Verifier {
	return BuildVerifierWithCertProvider(mgr.GetCertificateVisitor(mchID))
}
