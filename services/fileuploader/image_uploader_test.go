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

func TestImageUploader_Upload(t *testing.T) {
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

	file, err := os.Open("picture.jpg")
	defer file.Close()
	assert.Nil(t, err)

	svc := ImageUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "picture.jpg", consts.ImageJPG)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.MediaId)
	t.Log("Resp: ", resp)
	t.Log("APIResult: ", result)
	t.Log("Err: ", err)
}
