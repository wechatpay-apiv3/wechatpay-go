package downloader_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestAutoCertificateDownloader(t *testing.T) {
	ctx := context.Background()

	mgr := downloader.NewCertificateDownloaderMgrWithInterval(ctx, 5*time.Second)
	require.NotNil(t, mgr)
	defer mgr.Stop()

	privateKey, err := utils.LoadPrivateKey(consts.MchPrivateKey)
	require.NoError(t, err)
	opts := []core.ClientOption{
		option.WithMerchantCredential(consts.MchID, consts.SerialNo, privateKey),
		option.WithoutValidator(),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	err = mgr.RegisterDownloaderWithClient(ctx, client, consts.MchID, consts.MchAPIv3Key)
	require.NoError(t, err)

	err = mgr.RegisterDownloaderWithPrivateKey(ctx, privateKey, consts.SerialNo, consts.MchID, consts.MchAPIv3Key)
	require.NoError(t, err)

	provider := mgr.GetCertificateVisitor(consts.MchID)

	assert.NotEmpty(t, provider.GetAll(ctx))
	for serialNo, cert := range provider.GetAll(ctx) {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}

	time.Sleep(11 * time.Second)

	mgr.RemoveDownloader(ctx, consts.MchID)
	assert.Empty(t, provider.GetAll(ctx))
}
