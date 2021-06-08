package downloader

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"sync"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
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
	certificates map[string]*x509.Certificate
	apiSvc       *certificates.CertificatesApiService
	mchAPIv3Key  string
	lock         sync.RWMutex
}

func (o *CertificateDownloader) GetCertificate(serialNo string) (*x509.Certificate, bool) {
	o.lock.RLock()
	defer o.lock.RUnlock()

	cert, ok := o.certificates[serialNo]
	return cert, ok
}

func (o *CertificateDownloader) GetCertificateMap() map[string]*x509.Certificate {
	o.lock.RLock()
	defer o.lock.RUnlock()

	return o.certificates
}

func (o *CertificateDownloader) decryptCertificate(encryptCertificate *certificates.EncryptCertificate) (
	*x509.Certificate, error) {
	plaintext, err := utils.DecryptAES256GCM(o.mchAPIv3Key, *encryptCertificate.AssociatedData,
		*encryptCertificate.Nonce, *encryptCertificate.Ciphertext)
	if err != nil {
		return nil, fmt.Errorf("decrypt downloaded certificate failed: %v", err)
	}

	certificate, err := utils.LoadCertificate(plaintext)
	if err != nil {
		return nil, fmt.Errorf("load downloaded certificate failed: %v", err)
	}

	return certificate, nil
}

func (o *CertificateDownloader) updateCertificates(certificates map[string]*x509.Certificate) {
	o.lock.Lock()
	defer o.lock.Unlock()
	if isSameCertificateMap(o.certificates, certificates) {
		return
	}

	o.certificates = certificates
	o.apiSvc.Client = core.NewClientWithValidator(o.apiSvc.Client,
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

	certificateMap := make(map[string]*x509.Certificate)
	for _, rawCertificate := range resp.Data {
		certificate, err := o.decryptCertificate(rawCertificate.EncryptCertificate)
		if err != nil {
			return err
		}

		certificateMap[utils.GetCertificateSerialNumber(*certificate)] = certificate
	}

	if len(certificateMap) == 0 {
		return fmt.Errorf("no certificate downloaded")
	}

	o.updateCertificates(certificateMap)
	return nil
}

func NewCertificateDownloader(ctx context.Context, mchID string, privateKey *rsa.PrivateKey,
	certificateSerialNo string, mchAPIv3Key string) (*CertificateDownloader, error) {

	client, err := core.NewClient(ctx, core.WithMerchantCredential(mchID, certificateSerialNo, privateKey))
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

func NewCertificateDownloaderWithCredential(credential auth.Credential, mchAPIv3Key string) (*CertificateDownloader, error) {
	client, err := core.NewClient(context.Background(), core.WithCredential(credential))
	if err != nil {
		return nil, fmt.Errorf("create downloader failed, create client err:%v", err)
	}

	return NewCertificateDownloaderWithClient(client, mchAPIv3Key)
}
