package downloader

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"sync"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
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

// GetAll 获取平台证书Map
func (d *pseudoCertificateDownloader) GetAll(ctx context.Context) map[string]*x509.Certificate {
	return d.mgr.GetCertificateMap(ctx, d.mchID)
}

// Get 获取证书序列号对应的平台证书
func (d *pseudoCertificateDownloader) Get(ctx context.Context, serialNo string) (*x509.Certificate, bool) {
	return d.mgr.GetCertificate(ctx, d.mchID, serialNo)
}

// GetNewestSerial 获取最新的平台证书的证书序列号
func (d *pseudoCertificateDownloader) GetNewestSerial(ctx context.Context) string {
	return d.mgr.GetNewestCertificateSerial(ctx, d.mchID)
}

// ExportAll 获取平台证书内容Map
func (d *pseudoCertificateDownloader) ExportAll(ctx context.Context) map[string]string {
	return d.mgr.ExportCertificateMap(ctx, d.mchID)
}

// Export 获取证书序列号对应的平台证书内容
func (d *pseudoCertificateDownloader) Export(ctx context.Context, serialNo string) (string, bool) {
	return d.mgr.ExportCertificate(ctx, d.mchID, serialNo)
}

// CertificateDownloaderMgr 证书下载器管理器
// 可挂载证书下载器 CertificateDownloader，会定时调用 CertificateDownloader 下载最新的证书
//
// CertificateDownloaderMgr 不会被 GoGC 自动回收，不再使用时应调用 Stop 方法，防止发生资源泄漏
type CertificateDownloaderMgr struct {
	ctx           context.Context
	task          *task.RepeatedTask
	downloaderMap map[string]*CertificateDownloader
	lock          sync.RWMutex
}

// Stop 停止 CertificateDownloaderMgr 的自动下载 Goroutine
// 当且仅当不再需要当前管理器自动下载后调用
// 一旦调用成功，当前管理器无法再次启动
func (mgr *CertificateDownloaderMgr) Stop() {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.task.Stop()
}

// GetCertificate 获取商户的某个平台证书
func (mgr *CertificateDownloaderMgr) GetCertificate(ctx context.Context, mchID, serialNo string) (
	*x509.Certificate, bool,
) {
	mgr.lock.RLock()
	downloader, ok := mgr.downloaderMap[mchID]
	mgr.lock.RUnlock()

	if !ok {
		return nil, false
	}

	return downloader.Get(ctx, serialNo)
}

// GetCertificateMap 获取商户的平台证书Map
func (mgr *CertificateDownloaderMgr) GetCertificateMap(ctx context.Context, mchID string) map[string]*x509.Certificate {
	mgr.lock.RLock()
	downloader, ok := mgr.downloaderMap[mchID]
	mgr.lock.RUnlock()

	if !ok {
		return nil
	}
	return downloader.GetAll(ctx)
}

// GetNewestCertificateSerial 获取商户的最新的平台证书序列号
func (mgr *CertificateDownloaderMgr) GetNewestCertificateSerial(ctx context.Context, mchID string) string {
	mgr.lock.RLock()
	downloader, ok := mgr.downloaderMap[mchID]
	mgr.lock.RUnlock()

	if !ok {
		return ""
	}
	return downloader.GetNewestSerial(ctx)
}

// ExportCertificate 获取商户的某个平台证书内容
func (mgr *CertificateDownloaderMgr) ExportCertificate(ctx context.Context, mchID, serialNo string) (string, bool) {
	mgr.lock.RLock()
	downloader, ok := mgr.downloaderMap[mchID]
	mgr.lock.RUnlock()

	if !ok {
		return "", false
	}

	return downloader.Export(ctx, serialNo)
}

// ExportCertificateMap 导出商户的平台证书内容Map
func (mgr *CertificateDownloaderMgr) ExportCertificateMap(ctx context.Context, mchID string) map[string]string {
	mgr.lock.RLock()
	downloader, ok := mgr.downloaderMap[mchID]
	mgr.lock.RUnlock()

	if !ok {
		return nil
	}
	return downloader.ExportAll(ctx)
}

// GetCertificateVisitor 获取某个商户的平台证书访问器
func (mgr *CertificateDownloaderMgr) GetCertificateVisitor(mchID string) core.CertificateVisitor {
	return &pseudoCertificateDownloader{mgr: mgr, mchID: mchID}
}

func (mgr *CertificateDownloaderMgr) getTickHandler() func(time.Time) {
	return func(time.Time) {
		mgr.DownloadCertificates(mgr.ctx)
	}
}

// DownloadCertificates 让所有已注册下载器均进行一次下载
func (mgr *CertificateDownloaderMgr) DownloadCertificates(ctx context.Context) {
	tmpDownloaderMap := make(map[string]*CertificateDownloader)

	mgr.lock.RLock()
	for key, downloader := range mgr.downloaderMap {
		tmpDownloaderMap[key] = downloader
	}
	mgr.lock.RUnlock()

	for _, downloader := range tmpDownloaderMap {
		_ = downloader.DownloadCertificates(ctx)
	}
}

// RegisterDownloaderWithPrivateKey 向 Mgr 注册商户的平台证书下载器
func (mgr *CertificateDownloaderMgr) RegisterDownloaderWithPrivateKey(
	ctx context.Context, privateKey *rsa.PrivateKey,
	certificateSerialNo string, mchID string, mchAPIv3Key string,
) error {
	downloader, err := NewCertificateDownloader(ctx, mchID, privateKey, certificateSerialNo, mchAPIv3Key)
	if err != nil {
		return err
	}

	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.downloaderMap[mchID] = downloader
	return nil
}

// RegisterDownloaderWithClient 向 Mgr 注册商户的平台证书下载器
func (mgr *CertificateDownloaderMgr) RegisterDownloaderWithClient(
	ctx context.Context, client *core.Client, mchID string, mchAPIv3Key string,
) error {
	downloader, err := NewCertificateDownloaderWithClient(ctx, client, mchAPIv3Key)
	if err != nil {
		return err
	}

	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.downloaderMap[mchID] = downloader
	return nil
}

// RemoveDownloader 移除商户的平台证书下载器
// 移除后从 GetCertificateVisitor 接口获得的对应商户的 CertificateVisitor 将会失效，
// 请确认不再需要该商户的证书后再行移除，如果下载器存在，本接口将会返回该下载器。
func (mgr *CertificateDownloaderMgr) RemoveDownloader(_ context.Context, mchID string) *CertificateDownloader {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	downloader, ok := mgr.downloaderMap[mchID]
	if !ok {
		return nil
	}

	delete(mgr.downloaderMap, mchID)
	return downloader
}

func (mgr *CertificateDownloaderMgr) HasDownloader(_ context.Context, mchID string) bool {
	mgr.lock.RLock()
	defer mgr.lock.RUnlock()

	_, ok := mgr.downloaderMap[mchID]
	return ok
}

// NewCertificateDownloaderMgr 以默认间隔 DefaultDownloadInterval 创建证书下载管理器
// 该管理器将以 DefaultDownloadInterval 的间隔定期调度所有 Downloader 进行证书下载。
// 证书管理器一旦创建即启动，使用完毕请调用 Stop() 防止发生资源泄漏
func NewCertificateDownloaderMgr(ctx context.Context) *CertificateDownloaderMgr {
	return NewCertificateDownloaderMgrWithInterval(ctx, DefaultDownloadInterval)
}

// NewCertificateDownloaderMgrWithInterval 创建一个空证书下载管理器（自定义更新间隔）
//
// 更新间隔最大不建议超过 2 天，以免错过平台证书平滑切换窗口；
// 同时亦不建议小于 1 小时，以避免过多请求导致浪费
func NewCertificateDownloaderMgrWithInterval(
	ctx context.Context, downloadInterval time.Duration,
) *CertificateDownloaderMgr {
	if downloadInterval <= 0 {
		downloadInterval = DefaultDownloadInterval
	}

	downloader := CertificateDownloaderMgr{
		ctx:           ctx,
		downloaderMap: make(map[string]*CertificateDownloader),
	}
	downloader.task = task.NewRepeatedTask(downloadInterval, downloader.getTickHandler())
	downloader.task.Start()
	return &downloader
}
