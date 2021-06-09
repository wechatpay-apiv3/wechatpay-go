package certificate_visitors

import "crypto/x509"

// SimpleGetter 最简单的证书提供器 —— 证书Map
type SimpleGetter map[string]*x509.Certificate

func (p SimpleGetter) Get(serialNo string) (*x509.Certificate, bool) {
	certificate, ok := p[serialNo]
	return certificate, ok
}

func (p SimpleGetter) GetAll() map[string]*x509.Certificate {
	return p
}

func NewSimpleGetter(m map[string]*x509.Certificate) SimpleGetter {
	return m
}
