package downloader_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestCertificateDownloader(t *testing.T) {
	ctx := context.Background()

	privateKey, err := utils.LoadPrivateKey(consts.MchPrivateKey)
	require.NoError(t, err)
	opts := []core.ClientOption{
		core.WithMerchantCredential(consts.MchID, consts.SerialNo, privateKey),
		core.WithoutValidator(),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	d, err := downloader.NewCertificateDownloaderWithClient(client, consts.MchAPIv3Key)
	require.NoError(t, err)

	assert.NotEmpty(t, d.GetCertificateMap())
	for serialNo, cert := range d.GetCertificateMap() {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}
}
