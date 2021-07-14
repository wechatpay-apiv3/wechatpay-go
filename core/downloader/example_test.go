// Copyright 2021 Tencent Inc. All rights reserved.

package downloader_test

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
)

func ExampleNewCertificateDownloader_saveCert() {
	ctx := context.Background()

	var (
		mchID           string
		mchCertSerialNo string
		mchPrivateKey   *rsa.PrivateKey
		mchAPIv3Key     string
	)
	// 假设以上参数已初始化完成

	d, err := downloader.NewCertificateDownloader(ctx, mchID, mchPrivateKey, mchCertSerialNo, mchAPIv3Key)
	if err != nil {
		fmt.Println(err)
		return
	}

	for serialNo, certContent := range d.ExportAll(ctx) {
		// 将 certContent 写入文件 *.pem
		_, _ = serialNo, certContent
	}
}

func ExampleNewCertificateDownloaderMgr() {
	ctx := context.Background()
	mgr := downloader.NewCertificateDownloaderMgr(ctx)
	// CertificateDownloaderMgr 初始化完成，尚未注册任何 Downloader，不会进行任何证书下载

	var (
		mchID           string
		mchCertSerialNo string
		mchPrivateKey   *rsa.PrivateKey
		mchAPIv3Key     string
	)
	// 假设以上参数已初始化完成

	// 注册证书下载器
	if err := mgr.RegisterDownloaderWithPrivateKey(
		ctx, mchPrivateKey, mchCertSerialNo, mchID, mchAPIv3Key,
	); err == nil {
		fmt.Println(err)
		return
	}
	// 可以注册多个商户的证书下载器...

	// 获取证书访问器
	certVisitor := mgr.GetCertificateVisitor(mchID)

	// 使用 certVisitor 初始化 Validator 进行验签
	option.WithVerifier(verifiers.NewSHA256WithRSAVerifier(certVisitor))
}

func ExampleNewCertificateDownloaderMgr_useMgr() {
	var certDownloadMgr *downloader.CertificateDownloaderMgr
	// certDownloadMgr 已经初始化完成且注册了需要的 Downloader

	var (
		mchID           string
		mchCertSerialNo string
		mchPrivateKey   *rsa.PrivateKey
	)

	ctx := context.Background()
	client, err := core.NewClient(
		ctx,
		option.WithWechatPayAutoAuthCipherUsingDownloaderMgr(mchID, mchCertSerialNo, mchPrivateKey, certDownloadMgr),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	// 使用下载管理器初始化 Client 成功
	_ = client
}
