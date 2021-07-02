package downloader_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestNewCertificateDownloaderWithClient(t *testing.T) {
	patches := mockDownloadServer(t)
	defer patches.Reset()

	ctx := context.Background()

	privateKey, err := utils.LoadPrivateKey(mockMchPrivateKey)
	require.NoError(t, err)
	opts := []core.ClientOption{
		option.WithMerchantCredential(mockMchID, mockMchCertificateSerial, privateKey),
		option.WithoutValidator(),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	d, err := downloader.NewCertificateDownloaderWithClient(ctx, client, mockAPIv3Key)
	require.NoError(t, err)

	assert.NotEmpty(t, d.GetAll(ctx))
	for serialNo, cert := range d.GetAll(ctx) {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}

	// call Download
	err = d.DownloadCertificates(ctx)
	require.NoError(t, err)

	// call Download Again
	err = d.DownloadCertificates(ctx)
	require.NoError(t, err)
}

func TestNewCertificateDownloader(t *testing.T) {
	patches := mockDownloadServer(t)
	defer patches.Reset()

	privateKey, err := utils.LoadPrivateKey(mockMchPrivateKey)
	require.NoError(t, err)

	ctx := context.Background()

	d, err := downloader.NewCertificateDownloader(
		context.Background(), mockMchID, privateKey, mockMchCertificateSerial, mockAPIv3Key,
	)
	require.NoError(t, err)

	assert.NotEmpty(t, d.GetAll(ctx))
	for serialNo, cert := range d.GetAll(ctx) {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}

	// call Download
	err = d.DownloadCertificates(ctx)
	require.NoError(t, err)

	// call Download Again
	err = d.DownloadCertificates(ctx)
	require.NoError(t, err)
}
