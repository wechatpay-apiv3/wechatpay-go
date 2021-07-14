// Copyright 2021 Tencent Inc. All rights reserved.

package downloader

import (
	"context"
	"sync"
)

var (
	mgrInstance *CertificateDownloaderMgr
	mgrLock     sync.RWMutex
)

// MgrInstance 获取 CertificateDownloaderMgr 默认单例，将在首次调用本方法后初始化
// 本单例旨在伴随整个进程生命周期持续运行，请不要调用其 Stop 方法，否则可能影响平台证书的自动更新
//
// 如果你希望自行管理 Mgr 的生命周期，请使用 NewCertificateDownloaderMgr 方法创建额外的Mgr
func MgrInstance() *CertificateDownloaderMgr {
	// 首次访问使用读锁
	mgrLock.RLock()
	if mgrInstance != nil {
		defer mgrLock.RUnlock()
		return mgrInstance
	}
	mgrLock.RUnlock()

	// 确认不存在后切换为写锁，由于 Go 没有读锁升级写锁的能力，因此解锁并重新捕获后，需要再次检查是否存在
	mgrLock.Lock()
	defer mgrLock.Unlock()

	if mgrInstance == nil {
		mgrInstance = NewCertificateDownloaderMgr(context.Background())
	}
	return mgrInstance
}
