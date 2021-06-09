package downloader

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"sync"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert"
	"github.com/wechatpay-apiv3/wechatpay-go/utils/task"
)

const (
	// DefaultDownloadInterval 默认微信支付平台证书更新间隔
	DefaultDownloadInterval = 24 * time.Hour
)

type pseudoCertificateDownloader struct {
	mgr   *CertificateDownloaderMgr
	mchID string
}

func (o *pseudoCertificateDownloader) GetAll() map[string]*x509.Certificate {
	return o.mgr.GetCertificateMap(o.mchID)
}

func (o *pseudoCertificateDownloader) Get(serialNo string) (*x509.Certificate, bool) {
	return o.mgr.GetCertificate(o.mchID, serialNo)
}

func (o *pseudoCertificateDownloader) GetNewestSerial() string {
	return o.mgr.GetNewestCertificateSerial(o.mchID)
}

func (o *pseudoCertificateDownloader) ExportAll() map[string]string {
	return o.mgr.ExportCertificateMap(o.mchID)
}

func (o *pseudoCertificateDownloader) Export(serialNo string) (string, bool) {
	return o.mgr.ExportCertificate(o.mchID, serialNo)
}

// CertificateDownloaderMgr 证书下载器管理器
// 可挂载证书下载器 CertificateDownloader，会定时调用 CertificateDownloader 下载最新的证书
//
// CertificateDownloaderMgr 不会被 GoGC 自动回收，不再使用时应调用 Stop 方法，防止发生资源泄漏
type CertificateDownloaderMgr struct {
	task          *task.RepeatedTask
	downloaderMap map[string]*CertificateDownloader
	lock          sync.Mutex
}

func (o *CertificateDownloaderMgr) Stop() {
	o.lock.Lock()
	defer o.lock.Unlock()

	o.task.Stop()
}

func (o *CertificateDownloaderMgr) GetCertificate(mchID, serialNo string) (*x509.Certificate, bool) {
	o.lock.Lock()
	downloader, ok := o.downloaderMap[mchID]
	o.lock.Unlock()

	if !ok {
		return nil, false
	}

	return downloader.Get(serialNo)
}

func (o *CertificateDownloaderMgr) GetCertificateMap(mchID string) map[string]*x509.Certificate {
	o.lock.Lock()
	downloader, ok := o.downloaderMap[mchID]
	o.lock.Unlock()

	if !ok {
		return nil
	}
	return downloader.GetAll()
}

func (o *CertificateDownloaderMgr) GetNewestCertificateSerial(mchID string) string {
	o.lock.Lock()
	downloader, ok := o.downloaderMap[mchID]
	o.lock.Unlock()

	if !ok {
		return ""
	}
	return downloader.GetNewestSerial()
}

func (o *CertificateDownloaderMgr) ExportCertificate(mchID, serialNo string) (string, bool) {
	o.lock.Lock()
	downloader, ok := o.downloaderMap[mchID]
	o.lock.Unlock()

	if !ok {
		return "", false
	}

	return downloader.Export(serialNo)
}

func (o *CertificateDownloaderMgr) ExportCertificateMap(mchID string) map[string]string {
	o.lock.Lock()
	downloader, ok := o.downloaderMap[mchID]
	o.lock.Unlock()

	if !ok {
		return nil
	}
	return downloader.ExportAll()
}

func (o *CertificateDownloaderMgr) GetCertificateVisitor(mchID string) cert.CertificateVisitor {
	return &pseudoCertificateDownloader{mgr: o, mchID: mchID}
}

func (o *CertificateDownloaderMgr) getTickHandler() func(time.Time) {
	return func(time.Time) {
		o.DownloadCertificates()
	}
}

func (o *CertificateDownloaderMgr) DownloadCertificates() {
	tmpDownloaderMap := make(map[string]*CertificateDownloader)

	o.lock.Lock()
	for key, downloader := range o.downloaderMap {
		tmpDownloaderMap[key] = downloader
	}
	o.lock.Unlock()

	for _, downloader := range tmpDownloaderMap {
		_ = downloader.DownloadCertificates()
	}
}

func (o *CertificateDownloaderMgr) RegisterDownloaderWithPrivateKey(
	ctx context.Context, privateKey *rsa.PrivateKey,
	certificateSerialNo string, mchID string, mchAPIv3Key string,
) error {
	downloader, err := NewCertificateDownloader(ctx, mchID, privateKey, certificateSerialNo, mchAPIv3Key)
	if err != nil {
		return err
	}

	o.lock.Lock()
	defer o.lock.Unlock()

	o.downloaderMap[mchID] = downloader
	return nil
}

func (o *CertificateDownloaderMgr) RegisterDownloaderWithClient(
	client *core.Client, mchID string,
	mchAPIv3Key string,
) error {
	downloader, err := NewCertificateDownloaderWithClient(client, mchAPIv3Key)
	if err != nil {
		return err
	}

	o.lock.Lock()
	defer o.lock.Unlock()

	o.downloaderMap[mchID] = downloader
	return nil
}

func (o *CertificateDownloaderMgr) RegisterDownloaderWithCredential(
	credential auth.Credential, mchID string, mchAPIv3Key string,
) error {
	downloader, err := NewCertificateDownloaderWithCredential(credential, mchAPIv3Key)
	if err != nil {
		return err
	}

	o.lock.Lock()
	defer o.lock.Unlock()

	o.downloaderMap[mchID] = downloader
	return nil
}

func (o *CertificateDownloaderMgr) RemoveDownloader(mchID string) *CertificateDownloader {
	o.lock.Lock()
	defer o.lock.Unlock()

	downloader, ok := o.downloaderMap[mchID]
	if !ok {
		return nil
	}

	delete(o.downloaderMap, mchID)
	return downloader
}

// NewCertificateDownloaderMgr 以默认间隔 DefaultDownloadInterval 创建证书下载管理器
// 该管理器将以 DefaultDownloadInterval 的间隔定期调度所有 Downloader 进行证书下载。
// 证书管理器一旦创建即启动，使用完毕请调用 Stop() 防止发生资源泄漏
func NewCertificateDownloaderMgr() *CertificateDownloaderMgr {
	return NewCertificateDownloaderMgrWithInterval(DefaultDownloadInterval)
}

// NewCertificateDownloaderMgrWithInterval 创建一个空证书下载管理器（自定义更新间隔）
//
// 更新间隔最大不建议超过 2 天，以免错过平台证书平滑切换窗口；
// 同时亦不建议小于 1 小时，以避免过多请求导致浪费
func NewCertificateDownloaderMgrWithInterval(downloadInterval time.Duration) *CertificateDownloaderMgr {
	if downloadInterval <= 0 {
		downloadInterval = DefaultDownloadInterval
	}

	downloader := CertificateDownloaderMgr{
		downloaderMap: make(map[string]*CertificateDownloader),
	}
	downloader.task = task.NewRepeatedTask(downloadInterval, downloader.getTickHandler())
	downloader.task.Start()
	return &downloader
}
