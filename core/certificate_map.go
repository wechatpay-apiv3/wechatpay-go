package core

import (
	"context"
	"crypto/x509"

	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// CertificateMap 最简单的证书获取器——证书Map
type CertificateMap struct {
	m      map[string]*x509.Certificate
	newest string
}

// Reset 完整重设 CertificateMap 中存储的证书，并重新选择最新的证书
func (m *CertificateMap) Reset(newCertificates map[string]*x509.Certificate) {
	var (
		newestCert     *x509.Certificate
		newestSerialNo string
	)

	m.m = make(map[string]*x509.Certificate)

	for serialNo, cert := range newCertificates {
		m.m[serialNo] = cert
		if newestSerialNo == "" || newestCert == nil || cert.NotBefore.After(newestCert.NotBefore) {
			newestSerialNo = serialNo
			newestCert = cert
		}
	}

	m.newest = newestSerialNo
}

// Get 获取证书序列号对应的平台证书
func (m *CertificateMap) Get(_ context.Context, serialNo string) (*x509.Certificate, bool) {
	cert, ok := m.m[serialNo]
	return cert, ok
}

// GetAll 获取平台证书Map
func (m *CertificateMap) GetAll(_ context.Context) map[string]*x509.Certificate {
	ret := make(map[string]*x509.Certificate)

	for serialNo, cert := range m.m {
		ret[serialNo] = cert
	}

	return ret
}

// GetNewestSerial 获取最新的平台证书的证书序列号
func (m *CertificateMap) GetNewestSerial(_ context.Context) string {
	return m.newest
}

// NewCertificateMap 使用 证书序列号->证书 映射 初始化 CertificateMap
func NewCertificateMap(certificateMap map[string]*x509.Certificate) *CertificateMap {
	m := CertificateMap{}
	m.Reset(certificateMap)

	return &m
}

// NewCertificateMapWithList 使用 证书列表 初始化 CertificateMap
func NewCertificateMapWithList(l []*x509.Certificate) *CertificateMap {
	m := make(map[string]*x509.Certificate)
	for _, c := range l {
		m[utils.GetCertificateSerialNumber(*c)] = c
	}

	return NewCertificateMap(m)
}
