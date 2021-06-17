package certificates

import (
	"context"
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestCertificatesApiService_DownloadCertificates_WithoutValidator(t *testing.T) {
	ctx := context.Background()

	privateKey, err := utils.LoadPrivateKey(consts.MchPrivateKey)
	require.NoError(t, err)
	opts := []core.ClientOption{
		option.WithMerchantCredential(consts.MchID, consts.SerialNo, privateKey),
		option.WithoutValidator(),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	svc := CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	require.NoError(t, err)

	t.Log(resp)
	t.Log(result.Request)
	t.Log(result.Response)
}

func TestCertificatesApiService_DownloadCertificates_WithValidator(t *testing.T) {
	ctx := context.Background()

	privateKey, err := utils.LoadPrivateKey(consts.MchPrivateKey)
	require.NoError(t, err)
	wechatPayCert, err := utils.LoadCertificate(consts.WechatPayCertificate)
	require.NoError(t, err)
	opts := []core.ClientOption{
		option.WithMerchantCredential(consts.MchID, consts.SerialNo, privateKey),
		option.WithWechatPayCertificate([]*x509.Certificate{wechatPayCert}),
	}

	client, err := core.NewClient(ctx, opts...)
	require.NoError(t, err)

	svc := CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	require.NoError(t, err)

	t.Log(resp)
	t.Log(result.Request)
	t.Log(result.Response)
}
