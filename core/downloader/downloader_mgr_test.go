// Copyright 2021 Tencent Inc. All rights reserved.

package downloader_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestAutoCertificateDownloader(t *testing.T) {
	patches := mockDownloadServer(t)
	defer patches.Reset()

	ctx := context.Background()

	mgr := downloader.NewCertificateDownloaderMgrWithInterval(ctx, 5*time.Second)
	require.NotNil(t, mgr)
	defer mgr.Stop()

	privateKey, err := utils.LoadPrivateKey(testingKey(mockMchPrivateKey))
	require.NoError(t, err)
	opts := []core.ClientOption{
		option.WithMerchantCredential(mockMchID, mockMchCertificateSerial, privateKey),
		option.WithoutValidator(),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	err = mgr.RegisterDownloaderWithClient(ctx, client, mockMchID, mockAPIv3Key)
	require.NoError(t, err)

	err = mgr.RegisterDownloaderWithPrivateKey(ctx, privateKey, mockMchCertificateSerial, mockMchID, mockAPIv3Key)
	require.NoError(t, err)

	provider := mgr.GetCertificateVisitor(mockMchID)

	assert.NotEmpty(t, provider.GetAll(ctx))
	for serialNo, cert := range provider.GetAll(ctx) {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}

	time.Sleep(11 * time.Second)

	mgr.RemoveDownloader(ctx, mockMchID)
	assert.Empty(t, provider.GetAll(ctx))
}
