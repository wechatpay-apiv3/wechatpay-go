package core_test

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"testing"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"

	"github.com/stretchr/testify/assert"
)

const (
	testMchID                   = ""
	testCertificateSerialNumber = ""
	testPrivateKey              = `-----BEGIN PRIVATE KEY-----
-----END PRIVATE KEY-----`
	testWechatCertSerialNumber = ""
	testWechatCertificateStr   = `-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----`
	filePath  = ""
	fileName  = "picture.jpeg"
	postURL   = "https://api.mch.weixin.qq.com/v3/marketing/favor/users/oHkLxt_htg84TUEbzvlMwQzVDBqo/coupons"
	getURL    = "https://api.mch.weixin.qq.com/v3/certificates"
	uploadURL = "https://api.mch.weixin.qq.com/v3/merchant/media/upload"
)

var (
	privateKey           *rsa.PrivateKey
	wechatPayCertificate *x509.Certificate
	signer               auth.Signer
	verifier             auth.Verifier
	ctx                  context.Context
	err                  error
)

func init() {
	privateKey, err = utils.LoadPrivateKey(testPrivateKey)
	if err != nil {
		panic(fmt.Errorf("load private err:%s", err.Error()))
	}
	wechatPayCertificate, err = utils.LoadCertificate(testWechatCertificateStr)
	if err != nil {
		panic(fmt.Errorf("load certificate err:%s", err.Error()))
	}
	ctx = context.Background()
}

func TestGet(t *testing.T) {
	opts := []core.ClientOption{
		option.WithMerchantCredential(testMchID, testCertificateSerialNumber, privateKey),
		option.WithWechatPayCertificate([]*x509.Certificate{wechatPayCertificate}),
	}
	client, err := core.NewClient(ctx, opts...)
	assert.Nil(t, err)
	result, err := client.Get(ctx, getURL)
	assert.Nil(t, err)
	body, err := ioutil.ReadAll(result.Response.Body)
	assert.Nil(t, err)
	t.Log(string(body))
}

type testData struct {
	StockID           string `json:"stock_id"`
	StockCreatorMchID string `json:"stock_creator_mchid"`
	OutRequestNo      string `json:"out_request_no"`
	AppID             string `json:"appid"`
}

func TestPost(t *testing.T) {
	opts := []core.ClientOption{
		option.WithMerchantCredential(testMchID, testCertificateSerialNumber, privateKey),
		option.WithWechatPayCertificate([]*x509.Certificate{wechatPayCertificate}),
	}
	client, err := core.NewClient(ctx, opts...)
	assert.Nil(t, err)
	data := &testData{
		StockID:           "xxx",
		StockCreatorMchID: "xxx",
		OutRequestNo:      "xxx",
		AppID:             "xxx",
	}
	result, err := client.Post(ctx, postURL, data)
	assert.Nil(t, err)
	body, err := ioutil.ReadAll(result.Response.Body)
	assert.Nil(t, err)
	t.Log(string(body))
}

type meta struct {
	FileName string `json:"filename" binding:"required"` // 商户上传的媒体图片的名称，商户自定义，必须以JPG、BMP、PNG为后缀。
	Sha256   string `json:"sha256" binding:"required"`   // 图片文件的文件摘要，即对图片文件的二进制内容进行sha256计算得到的值。
}

func TestClient_Upload(t *testing.T) {
	// 如果你有自定义的Signer或者Verifier
	signer = &signers.SHA256WithRSASigner{
		MchID:               testMchID,
		PrivateKey:          privateKey,
		CertificateSerialNo: testCertificateSerialNumber,
	}

	verifier = verifiers.NewSHA256WithRSAVerifier(
		core.NewCertificateMap(
			map[string]*x509.Certificate{testWechatCertSerialNumber: wechatPayCertificate},
		),
	)

	client, err := core.NewClient(ctx, option.WithSigner(signer), option.WithVerifier(verifier))
	assert.Nil(t, err)
	pictureByes, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	// 计算文件序列化后的sha256
	h := sha256.New()
	_, err = h.Write(pictureByes)
	assert.Nil(t, err)
	metaObject := &meta{}
	pictureSha256 := h.Sum(nil)
	metaObject.FileName = fileName
	metaObject.Sha256 = fmt.Sprintf("%x", string(pictureSha256))
	metaByte, _ := json.Marshal(metaObject)
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)
	err = core.CreateFormField(writer, "meta", "application/json", metaByte)
	assert.Nil(t, err)
	err = core.CreateFormFile(writer, fileName, "image/jpg", pictureByes)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)
	result, err := client.Upload(ctx, uploadURL, string(metaByte), reqBody.String(), writer.FormDataContentType())
	assert.Nil(t, err)
	if result.Response.Body != nil {
		defer result.Response.Body.Close()
	}
	body, err := ioutil.ReadAll(result.Response.Body)
	assert.Nil(t, err)
	t.Log(string(body))
}
