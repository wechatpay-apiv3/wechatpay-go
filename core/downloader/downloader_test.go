package downloader_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestNewCertificateDownloaderWithClient(t *testing.T) {
	ctx := context.Background()

	privateKey, err := utils.LoadPrivateKey(consts.MchPrivateKey)
	require.NoError(t, err)
	opts := []core.ClientOption{
		option.WithMerchantCredential(consts.MchID, consts.SerialNo, privateKey),
		option.WithoutValidator(),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	d, err := downloader.NewCertificateDownloaderWithClient(ctx, client, consts.MchAPIv3Key)
	require.NoError(t, err)

	assert.NotEmpty(t, d.GetAll(ctx))
	for serialNo, cert := range d.GetAll(ctx) {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}
}

func TestNewCertificateDownloader(t *testing.T) {
	privateKey, err := utils.LoadPrivateKey(consts.MchPrivateKey)
	require.NoError(t, err)

	ctx := context.Background()

	d, err := downloader.NewCertificateDownloader(
		context.Background(), consts.MchID, privateKey, consts.SerialNo, consts.MchAPIv3Key,
	)
	require.NoError(t, err)

	assert.NotEmpty(t, d.GetAll(ctx))
	for serialNo, cert := range d.GetAll(ctx) {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}
}
