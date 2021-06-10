package certificate_map

import (
	"context"
	"crypto/x509"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// CertificateMap 最简单的证书获取器——证书Map
type CertificateMap struct {
	m map[string]*x509.Certificate
	newest string
}

func (m *CertificateMap) Reset(newCertificates map[string]*x509.Certificate) {
	var (
		newestCert *x509.Certificate
		newestSerialNo string
	)

	m.m = make(map[string]*x509.Certificate)

	for serialNo, cert := range newCertificates {
		m.m[serialNo] = cert
		if newestSerialNo == "" || newestCert == nil || cert.NotBefore.After(newestCert.NotBefore){
			newestSerialNo = serialNo
			newestCert = cert
		}
	}

	m.newest = newestSerialNo
}

func (m *CertificateMap) Get(_ context.Context, serialNo string) (*x509.Certificate, bool) {
	cert, ok := m.m[serialNo]
	return cert, ok
}

func (m *CertificateMap) GetAll(_ context.Context) map[string]*x509.Certificate {
	ret := make(map[string]*x509.Certificate)

	for serialNo, cert := range m.m {
		ret[serialNo] = cert
	}

	return ret
}

func (m *CertificateMap) GetNewestSerial(_ context.Context) string {
	return m.newest
}

func NewCertificateMap(certificateMap map[string]*x509.Certificate) *CertificateMap {
	m := CertificateMap{}
	m.Reset(certificateMap)

	return &m
}

func NewCertificateMapWithList(l []*x509.Certificate) *CertificateMap {
	m := make(map[string]*x509.Certificate)
	for _, c := range l {
		m[utils.GetCertificateSerialNumber(*c)] = c
	}

	return NewCertificateMap(m)
}
