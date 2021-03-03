package core

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

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
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
	GetURL    = "https://api.mch.weixin.qq.com/v3/certificates"
	uploadURL = "https://api.mch.weixin.qq.com/v3/merchant/media/upload"
)

var (
	privateKey           *rsa.PrivateKey
	wechatPayCertificate *x509.Certificate
	credential           auth.Credential
	validator            auth.Validator
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
	opts := []option.ClientOption{
		option.WithMerchant(testMchID, testCertificateSerialNumber, privateKey),
		option.WithWechatPay([]*x509.Certificate{wechatPayCertificate}),
	}
	client, err := NewClient(ctx, opts...)
	assert.Nil(t, err)
	response, err := client.Get(ctx, GetURL)
	assert.Nil(t, err)
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)
	t.Log(string(body))
	assert.Nil(t, CheckResponse(response))
}

type testData struct {
	StockID           string `json:"stock_id"`
	StockCreatorMchID string `json:"stock_creator_mchid"`
	OutRequestNo      string `json:"out_request_no"`
	AppID             string `json:"appid"`
}

func TestPost(t *testing.T) {
	opts := []option.ClientOption{
		option.WithMerchant(testMchID, testCertificateSerialNumber, privateKey),
		option.WithWechatPay([]*x509.Certificate{wechatPayCertificate}),
	}
	client, err := NewClient(ctx, opts...)
	assert.Nil(t, err)
	data := &testData{
		StockID:           "xxx",
		StockCreatorMchID: "xxx",
		OutRequestNo:      "xxx",
		AppID:             "xxx",
	}
	response, err := client.Post(ctx, postURL, data)
	assert.Nil(t, err)
	assert.NotNil(t, CheckResponse(response))
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)
	t.Log(string(body))
}

type meta struct {
	FileName string `json:"filename" binding:"required"` // 商户上传的媒体图片的名称，商户自定义，必须以JPG、BMP、PNG为后缀。
	Sha256   string `json:"sha256" binding:"required"`   // 图片文件的文件摘要，即对图片文件的二进制内容进行sha256计算得到的值。
}

func TestClient_Upload(t *testing.T) {
	// 如果你有自定义的Signer或者Verfifer
	credential = &credentials.WechatPayCredentials{
		Signer:              &signers.Sha256WithRSASigner{PrivateKey: privateKey},
		MchID:               testMchID,
		CertificateSerialNo: testCertificateSerialNumber,
	}
	validator = &validators.WechatPayValidator{
		Verifier: &verifiers.WechatPayVerifier{
			Certificates: map[string]*x509.Certificate{
				testWechatCertSerialNumber: wechatPayCertificate,
			},
		},
	}
	client, err := NewClient(ctx, option.WithCredential(credential), option.WithValidator(validator))
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
	err = CreateFormField(writer, "meta", "application/json", metaByte)
	assert.Nil(t, err)
	err = CreateFormFile(writer, fileName, "image/jpg", pictureByes)
	assert.Nil(t, err)
	err = writer.Close()
	assert.Nil(t, err)
	response, err := client.Upload(ctx, uploadURL, string(metaByte), reqBody.String(), writer.FormDataContentType())
	assert.Nil(t, err)
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)
	t.Log(string(body))
}
