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

func TestFileUploader_Upload(t *testing.T) {
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

	svc := baseFileUploader{Client: client}
	result, err := svc.upload(ctx,
		"/v3/merchant/media/upload",
		file,
		"picture.jpg",
		consts.ImageJPG,
		map[string]interface{}{"sub_mch_code": 123})
	assert.Nil(t, err)
	assert.NotNil(t, result.Response)
	t.Log("APIResult:", result)
	t.Log("HttpResponse:", *result.Response)
	t.Log("Err:", err)
}
