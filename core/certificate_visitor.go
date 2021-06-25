package core

import (
	"context"
	"crypto/x509"
)

// CertificateGetter 平台证书提供器
type CertificateGetter interface {
	// Get 获取证书序列号对应的平台证书
	Get(ctx context.Context, serialNo string) (*x509.Certificate, bool)
	// GetAll 获取平台证书Map
	GetAll(ctx context.Context) map[string]*x509.Certificate
	// GetNewestSerial 获取最新的平台证书的证书序列号
	GetNewestSerial(ctx context.Context) string
}

// CertificateExporter 平台证书导出器，可获取平台证书内容，
type CertificateExporter interface {
	// Export 获取证书序列号对应的平台证书内容
	Export(ctx context.Context, serialNo string) (string, bool)
	// ExportAll 获取平台证书内容Map
	ExportAll(ctx context.Context) map[string]string
}

// CertificateVisitor 证书访问器，集 CertificateGetter 与 CertificateExporter 于一体
type CertificateVisitor interface {
	CertificateGetter
	CertificateExporter
}
