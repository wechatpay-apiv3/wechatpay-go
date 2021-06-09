package downloader_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cert/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestAutoCertificateDownloader(t *testing.T) {
	mgr := downloader.NewCertificateDownloaderMgrWithInterval(5 * time.Second)
	require.NotNil(t, mgr)
	defer mgr.Stop()

	ctx := context.Background()

	privateKey, err := utils.LoadPrivateKey(consts.MchPrivateKey)
	require.NoError(t, err)
	opts := []core.ClientOption{
		core.WithMerchantCredential(consts.MchID, consts.SerialNo, privateKey),
		core.WithoutValidator(),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	err = mgr.RegisterDownloaderWithClient(client, consts.MchID, consts.MchAPIv3Key)
	require.NoError(t, err)

	provider := mgr.GetCertificateVisitor(consts.MchID)

	assert.NotEmpty(t, provider.GetAll())
	for serialNo, cert := range provider.GetAll() {
		assert.Equal(t, serialNo, utils.GetCertificateSerialNumber(*cert))
	}

	time.Sleep(11 * time.Second)

	mgr.RemoveDownloader(consts.MchID)
	assert.Empty(t, provider.GetAll())
}
