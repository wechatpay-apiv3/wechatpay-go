package certificate_providers

import "crypto/x509"

// SimpleCertificateProvider 最简单的证书提供器 —— 证书Map
type SimpleCertificateProvider map[string]*x509.Certificate

func (p SimpleCertificateProvider) GetCertificate(serialNo string) (*x509.Certificate, bool) {
	certificate, ok := p[serialNo]
	return certificate, ok
}

func (p SimpleCertificateProvider) GetCertificateMap() map[string]*x509.Certificate {
	return p
}
