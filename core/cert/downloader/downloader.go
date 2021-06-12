package downloader

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"sync"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/certificate_map"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// isSameCertificateMap Check if two CertificateMaps stores same certificates.
// Normally, checking serial number set is enough.
func isSameCertificateMap(l, r map[string]*x509.Certificate) bool {
	if l == nil && r == nil {
		return true
	}

	if len(l) != len(r) {
		return false
	}

	for serialNo := range l {
		if _, ok := r[serialNo]; !ok {
			return false
		}
	}

	return true
}

type CertificateDownloader struct {
	certContents map[string]string                    // 证书文本内容，用于导出
	certificates certificate_map.CertificateMap       // 证书实例
	apiSvc       *certificates.CertificatesApiService // 平台证书下载API
	mchAPIv3Key  string                               // 商户APIv3密钥
	lock         sync.RWMutex
}

func (o *CertificateDownloader) Get(serialNo string) (*x509.Certificate, bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()

	return o.certificates.Get(serialNo)
}

func (o *CertificateDownloader) GetAll() map[string]*x509.Certificate {
	o.lock.RLock()
	defer o.lock.RUnlock()

	return o.certificates.GetAll()
}

func (o *CertificateDownloader) GetNewestSerial() string {
	o.lock.RLock()
	defer o.lock.RUnlock()

	return o.certificates.GetNewestSerial()
}

func (o *CertificateDownloader) Export(serialNo string) (string, bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()

	content, ok := o.certContents[serialNo]
	return content, ok
}

func (o *CertificateDownloader) ExportAll() map[string]string {
	o.lock.RLock()
	defer o.lock.RUnlock()

	ret := make(map[string]string)
	for serialNo, content := range o.certContents {
		ret[serialNo] = content
	}

	return ret
}

func (o *CertificateDownloader) decryptCertificate(encryptCertificate *certificates.EncryptCertificate) (
	string, error,
) {
	plaintext, err := utils.DecryptAES256GCM(
		o.mchAPIv3Key, *encryptCertificate.AssociatedData,
		*encryptCertificate.Nonce, *encryptCertificate.Ciphertext,
	)
	if err != nil {
		return "", fmt.Errorf("decrypt downloaded certificate failed: %v", err)
	}

	return plaintext, nil
}

func (o *CertificateDownloader) updateCertificates(
	certContents map[string]string, certificates map[string]*x509.Certificate,
) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if isSameCertificateMap(o.certificates.GetAll(), certificates) {
		return
	}

	o.certContents = certContents
	o.certificates.Reset(certificates)
	o.apiSvc.Client = core.NewClientWithValidator(
		o.apiSvc.Client,
		&validators.WechatPayValidator{
			Verifier: verifiers.NewSHA256WithRSAVerifier(o),
		},
	)
}

func (o *CertificateDownloader) DownloadCertificates() error {
	resp, _, err := o.apiSvc.DownloadCertificates(context.Background())
	if err != nil {
		return err
	}

	rawCertContentMap := make(map[string]string)
	certificateMap := make(map[string]*x509.Certificate)
	for _, rawCertificate := range resp.Data {
		certContent, err := o.decryptCertificate(rawCertificate.EncryptCertificate)
		if err != nil {
			return err
		}

		certificate, err := utils.LoadCertificate(certContent)
		if err != nil {
			return fmt.Errorf("parse downlaoded certificate failed: %v, certcontent:%v", err, certContent)
		}

		serialNo := utils.GetCertificateSerialNumber(*certificate)

		rawCertContentMap[serialNo] = certContent
		certificateMap[serialNo] = certificate
	}

	if len(certificateMap) == 0 {
		return fmt.Errorf("no certificate downloaded")
	}

	o.updateCertificates(rawCertContentMap, certificateMap)
	return nil
}

func NewCertificateDownloader(
	ctx context.Context, mchID string, privateKey *rsa.PrivateKey,
	certificateSerialNo string, mchAPIv3Key string,
) (*CertificateDownloader, error) {
	client, err := core.NewClient(
		ctx, core.WithMerchantCredential(mchID, certificateSerialNo, privateKey), core.WithoutValidator(),
	)
	if err != nil {
		return nil, fmt.Errorf("create downloader failed, create client err:%v", err)
	}

	return NewCertificateDownloaderWithClient(client, mchAPIv3Key)
}

func NewCertificateDownloaderWithClient(client *core.Client, mchAPIv3Key string) (*CertificateDownloader, error) {
	downloader := CertificateDownloader{
		apiSvc:      &certificates.CertificatesApiService{Client: client},
		mchAPIv3Key: mchAPIv3Key,
	}

	if err := downloader.DownloadCertificates(); err != nil {
		return nil, err
	}

	return &downloader, nil
}
