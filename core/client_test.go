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
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/signers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

const (
	testMchID                   = "example-mchid"
	testCertificateSerialNumber = "example-sn"
	testPrivateKey              = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCkxOav8p5RFFmN\n7hjLGrNtXPYgCd0Zuvxabv+IWl1HVkWi/1iVqac+XwKH/ZeCYDURqZ6P0iiq8NBd\npygoeJiQM+qzaTV3alNXdLcpcaQSNqJ16a7Z2Co5LgOkBJHIF1qUWf+BpAtyLqo5\niGUQ47w3IWQHtfpW7RrECiNsI7kGKuSc4U2JX/gxrFG7ugpRA9Gp1eF0/wBMWSKV\nmKveKERvAueaTKEujpN4lctoO1wsMW93nuFNH1gtHPYmkaZaJS88GEp0VYJIcOpX\n2HlVPPiWx+0KAndcqMLVQ+qTk21tivpjuqPDstxcT9cXn/3CzSEkYrKkMLpjpSl/\nIn+amHddAgMBAAECggEADtQRlsAU82MLdDR7UrwCbdMx60w387raPyFCKflH78WZ\n2sN0K3PrMzfFuItf+UHDROWo+XSGaGvntKX4fTvtLv0dICxVvXt6KKK+YSJzC5iT\nIl13eO91TVQQy9AFdqZzZmp7DiW/SfVdKHRX9B8qryN4JyF/eBc6k23+JhtI6X8J\nrPeXGcw/Muo2cvoZ6oarRvzKU7pDivZADavAsgGwi+QwZUtrpUhDbJxWxoasCrWz\n6X24D+JbKndf/AeyHC2mqoAwTYdQgkkPHLJWGsqAt/GHtHLmZHIPc0dXfksP5omX\nOIii34YNud1j2/X0ryxoRBKPGIolV5Tyyh9PX75UAQKBgQDa/oPFUNsBo+NkfE4r\nMry+mnAeBM06vZh65acbHu3prqzn7tzfQR5rF9nIEUjNkach+ZeTnUxKfX4TwfiL\nibbM1Epb+yhEpSlA9HhY1AGhdNKpFbq63oVzS+lwpmLcXFLHhOYw8GrnYH9lcE+E\nYCqFcuk4t/y3rU/8Y5GhZKZ53QKBgQDAnKk6CEgnSkXfZLmzVtoM/5hnI0GEDOUo\nV35alqvgdJtCiPs4C03snYLVHqHjAknGzLGONAQ6h9au2qwcHy1qH/noq151u92b\nbCrKmghnm2SoIgCaZ7i2scWm6NM9Da60H662WxjaKcZMnUClm+G+Irl9m5cm1i3V\n56ZU63nbgQKBgHyAlFO6mzg8f4via+J9TvciADngyvjpT2YXaECv/dyL9TtK/oFi\nmTOTdLocsYJFm3piVv2SQQxcejArZ+2U1rtuufO/P25/Y4vNMRp3NZIgQ5/jfay9\n06rv7oCf57aWOm26LdCG7pAquWLnTh3ZOnNyGAup9mBKhR3dUa8q9MZ1AoGATZGJ\n0VYugKw3sXymEKRkkiGJJdgb9WsgCnwZ5a+SLoWnVUdHLM3YpvbUDrIUbhCo14ft\n5Z/rKAs2mRp1f6nKp1eTVHFXTEDJQWNxZEBeLCN3iQKQjZ5B1EmJmOtgztCoz9+G\ng+fx/UIfmxElTMyXP/RKEVzMpZZRxThSUxa174ECgYEAsviAmgBskM2ibtrA8tNW\njHdgut0xAtIJIIHlYmtksWgAWD4cPtCg7HPurXqBKxMogH/ZsZc6/5PpOIRYBNl0\nEebH0MZ/yiEOrmgsFJ1gKWk3fx8/yLQBlhn32AIhj7wmcFwzi/4hcwihHRCjS7t5\nQhVpKswxQyxqeIdQw0CgKZY=\n-----END PRIVATE KEY-----"

	testWechatPrivateKeyStr  = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDzbMaiMbCzJ11sZ2r7/XisolGu1pVpWvnv1PVOAWZZAWr/WcNNuLni8ddJkIs8NdTz+iAHtb9XMNG4hj7d10cy8QE6QG8YUef1fGb47Wee3DjJRk8N9lWyPDAy9AW70yYWItl/05XgkGt2eJuoU5CcZ+Cy0u2nAXxEGs2Z4Fg/100Ylcyq4GrimjngIyUFLnowOcZltJUQSw/Iu63V0BEh9PMNnXhKkwS2xfkA2nZRzSgzpMvX74v8F7Zf+HkyrHHzwY/7YUWg3pSj3GD0xKsJOwzz0MLlS8uIdLj1lKGzzt1ROgwe0sM5LL5XMfmbjDhcVBmyxQI80WiNaC281tVhAgMBAAECggEBAJ134Wrs0Ayky2ej4u5OAvFSM5rxj0fPJV3DGkiy2R18sFWtII03kXBA1+7rxVZW0IJfbLbwGG3z08cVeLeTWqiWhR/ErNlDqtT/+7DOCrkWZtm1VNCIaNla3Ccp+keNiNbLBn4NRqg1ZH8H+FHEdQjonc+waTIe4N9Bo30GRrBMeMbAgN8mwQhZ+6R23j3GsrJOViFpPRgGhih4aEAORxU+DWl22vklxO8lnDzSwfnXvNDvapnPaA8VXekkThxABq3p/ggv5MI1QPYl8BU5PWd6AJlPs/u2nzFGmCgVufHMjdszWYd3hbj5EmhlL1VMs4uxhCC8OM+ypbnx0CmBB5ECgYEA+W3KJciw4qG6XqWgjK9hkgiZrO8z3tu6tQ6ge28f3BxYEbGUab8bKKzacbYODRiRCM5oKcrXfTqP6IbqUoESiuoz0CUPbShp7k00wXKd7BH8LoIECDbctz77NG0KBE31OGEJw4hm762M14V9nU0KDHGuud1CH1fqivTGG0g2HiUCgYEA+dZ+e5iSvin+omXkr3Vhwf3kutX+GNkKm5LrWZNmWybOKw/K1YFvESpfl1b6YgjA/qUXj0tbOumFQw6e/OLfxIvK/dtkd+7pbSC4T9w8rH7zgjYdJ030Nyv3UGfHDbES+z9MDMo5h+3RKsLN2bp1JcXp6vht9CiXDYm1df3O340CgYB118Qg08+WU1iM7O2MajPL3dpVFPJJwUBV2GJDzv2bbZzCR0baKxr2vau6+4tp7ohfQ718uUPT+34QGuXMMwUCsqHmHgxKw0RA/SMGnlM0PE8L3gtvohPnU481dqq72+UWTOpjAie35yPak0wErGgp9u/ZCkr6Kfw6yGhsbVJ8LQKBgEBLxS1FrK4n3JIqqtnE2a21C4JRxBzc7m/vNYZN+s+GgxRt8gNUViMSxpsKFVHZcuGV1yRXflkA8/y37I6kTHYmi80dAxQidgxRmV1kDnFOEpj2GDafRzRTqkgVDRMm+P2T4pyABqJGv8fDbnqUE8Xu0y5XVOS69XTUddCxyuWZAoGAbpF6JOh6B7OV4XRTDm98Z8OPYmYd9JQ4xt8bqsG9LdzvhU/PI4zwIaKDqZ8vzCI+r8TOrC6SBEfjAe6o2FEExmFWTjBAVCp+Qvnz+Pj7d+WP3kCX/B62IZckVhdV3a2frMTBPvAh8XdENdvsu4DsWJBCA54GLU3wdUa/FO0RUsA=\n-----END PRIVATE KEY-----\n"
	testWechatCertificateStr = "-----BEGIN CERTIFICATE-----\nMIID1TCCAr2gAwIBAgIRAJ8qZJYAQUwUheimQ8sQNZMwDQYJKoZIhvcNAQELBQAw\nXjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU\nZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w\nHhcNMjEwNjI3MTMwNTM2WhcNMjIwNjI3MTMwNTM2WjAkMQswCQYDVQQGEwJDTjEV\nMBMGA1UEAxMMd2VjaGF0cGF5LWdvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEA82zGojGwsyddbGdq+/14rKJRrtaVaVr579T1TgFmWQFq/1nDTbi54vHX\nSZCLPDXU8/ogB7W/VzDRuIY+3ddHMvEBOkBvGFHn9Xxm+O1nntw4yUZPDfZVsjww\nMvQFu9MmFiLZf9OV4JBrdnibqFOQnGfgstLtpwF8RBrNmeBYP9dNGJXMquBq4po5\n4CMlBS56MDnGZbSVEEsPyLut1dARIfTzDZ14SpMEtsX5ANp2Uc0oM6TL1++L/Be2\nX/h5Mqxx88GP+2FFoN6Uo9xg9MSrCTsM89DC5UvLiHS49ZShs87dUToMHtLDOSy+\nVzH5m4w4XFQZssUCPNFojWgtvNbVYQIDAQABo4HHMIHEMA4GA1UdDwEB/wQEAwIH\ngDATBgNVHSUEDDAKBggrBgEFBQcDAzAfBgNVHSMEGDAWgBQogSYF0TQaP8FzD7uT\nzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGGFWh0dHA6Ly9vY3Nw\nLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15c3NsLmNvbS9teXNz\nbHRlc3Ryc2EuY3J0MBcGA1UdEQQQMA6CDHdlY2hhdHBheS1nbzANBgkqhkiG9w0B\nAQsFAAOCAQEAjh4oxMcJqsVaN5/aA+4+NSfV9wR4uzTVtAyL/dApymZn6Wjknd85\nDltekcTflNP84bDiFEE3Ls3RYatjRx9pWeW7QbpdYvfDtWuxL5dhzRYtUO83z8wT\n+/sceeyNOQAWGD6Gt7Aw7yb7bIZ5slcZYepqdKSHyMnn06CCNRZtDVTuQYRqnmoh\nCaK5RNe4lYM/hncMgddE/DugTxzh5NUMpAY4xAsqOofkVmptX3trVZVILPglJ6nQ\n5dALCpp2UCuxikwFdEpvvGIC2qZQv5jmemFLCDIQZ227GUZ/EcbuTtAdQYHUnGJT\nvrFBSDe4FbKwljUmrccD/LkR9FmPn6gWKA==\n-----END CERTIFICATE-----\n"

	filePath  = ""
	fileName  = "picture.jpeg"
	postURL   = "https://api.mch.weixin.qq.com/v3/marketing/favor/users/oHkLxt_htg84TUEbzvlMwQzVDBqo/coupons"
	uploadURL = "https://api.mch.weixin.qq.com/v3/merchant/media/upload"

	wechatPaySignFormat = "%s\n%s\n%s\n"
	responseBody        = "hello client"
)

var (
	privateKey           *rsa.PrivateKey
	wechatPayPrivateKey  *rsa.PrivateKey
	wechatPayCertificate *x509.Certificate
	signer               auth.Signer
	verifier             auth.Verifier
	ctx                  context.Context
	ts                   *httptest.Server
)

func init() {
	ctx = context.Background()

	var err error
	privateKey, err = utils.LoadPrivateKey(testPrivateKey)
	wechatPayCertificate, err = utils.LoadCertificate(testWechatCertificateStr)
	wechatPayPrivateKey, err = utils.LoadPrivateKey(testWechatPrivateKeyStr)
	if err != nil {
		panic(fmt.Errorf("generate merchant key err:%s", err.Error()))
	}
}

func writeResponse(w http.ResponseWriter) {
	writeSignature(w, responseBody)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, responseBody)
}

func writeSignature(w http.ResponseWriter, body string) {
	w.Header().Set("Request-Id", "0")
	w.Header().Set("Wechatpay-Serial", utils.GetCertificateSerialNumber(*wechatPayCertificate))

	nonce := "this-is-a-nonce"
	w.Header().Set("Wechatpay-Nonce", nonce)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	w.Header().Set("Wechatpay-Timestamp", timestamp)

	signature, _ := utils.SignSHA256WithRSA(
		fmt.Sprintf(wechatPaySignFormat, timestamp, nonce, body), wechatPayPrivateKey)
	w.Header().Set("Wechatpay-Signature", signature)
}

func TestGet(t *testing.T) {
	opts := []core.ClientOption{
		option.WithMerchantCredential(testMchID, testCertificateSerialNumber, privateKey),
		option.WithWechatPayCertificate([]*x509.Certificate{wechatPayCertificate}),
	}
	client, err := core.NewClient(ctx, opts...)
	assert.Nil(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Header.Get("Authorization"))
		writeResponse(w)
	}))
	defer ts.Close()

	result, err := client.Get(ctx, ts.URL)
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
			map[string]*x509.Certificate{utils.GetCertificateSerialNumber(*wechatPayCertificate): wechatPayCertificate},
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
