package fileuploader

import (
	"context"
	"crypto/x509"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func TestVideoUploader_Upload(t *testing.T) {
	const (
		testMchID                   = ""
		testCertificateSerialNumber = ""
		testPrivateKey              = `-----BEGIN PRIVATE KEY-----
-----END PRIVATE KEY-----`
		testWechatCertificateStr   = `-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----`
	)
	ctx := context.Background()

	privateKey, _ := utils.LoadPrivateKey(testPrivateKey)
	wechatPayCert, _ := utils.LoadCertificate(testWechatCertificateStr)
	opts := []core.ClientOption{
		core.WithMerchantCredential(testMchID, testCertificateSerialNumber, privateKey),
		core.WithWechatPayValidator([]*x509.Certificate{wechatPayCert}),
	}
	client, _ := core.NewClient(ctx, opts...)

	file, err := os.Open("video.mp4")
	defer file.Close()
	assert.Nil(t, err)

	svc := VideoUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "video.mp4", consts.VideoMP4)
	assert.Nil(t, err)
	assert.NotEmpty(t, resp.MediaId)
	t.Log("Resp:", resp)
	t.Log("APIResult:", result)
	t.Log("Err:", err)
}
