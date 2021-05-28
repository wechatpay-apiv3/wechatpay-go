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
	ctx := context.Background()

	privateKey, _ := utils.LoadPrivateKey(consts.MchPrivateKey)
	wechatPayCert, _ := utils.LoadCertificate(consts.WechatPayCertificate)
	opts := []core.ClientOption{
		core.WithMerchantCredential(consts.MchID, consts.SerialNo, privateKey),
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
