package cert

import (
	"crypto/x509"
)

// CertificateProvider 平台证书提供器
type CertificateProvider interface {
	// GetCertificate 获取证书序列号对应的平台证书
	GetCertificate(serialNo string) (*x509.Certificate, bool)
	// GetCertificateMap 获取平台证书Map
	GetCertificateMap() map[string]*x509.Certificate
}