package cert

import (
	"crypto/x509"
)

// CertificateGetter 平台证书提供器
type CertificateGetter interface {
	// Get 获取证书序列号对应的平台证书
	Get(serialNo string) (*x509.Certificate, bool)
	// GetAll 获取平台证书Map
	GetAll() map[string]*x509.Certificate
}

// CertificateExporter 平台证书导出器，可获取平台证书内容，
type CertificateExporter interface {
	// Export 获取证书序列号对应的平台证书内容
	Export(serialNo string) (string, bool)
	// ExportAll 获取平台证书内容Map
	ExportAll() map[string]string
}

type CertificateVisitor interface {
	CertificateGetter
	CertificateExporter
}
