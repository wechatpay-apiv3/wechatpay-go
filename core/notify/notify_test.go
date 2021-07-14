// Copyright 2021 Tencent Inc. All rights reserved.

package notify

import (
	"bytes"
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"

	"github.com/agiledragon/gomonkey"
)

func Test_getRequestBody(t *testing.T) {
	body := "fake req body"
	bodyBuf := &bytes.Buffer{}
	bodyBuf.WriteString(body)

	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", bodyBuf)

	bodyBytes, err := getRequestBody(req)
	require.NoError(t, err)
	assert.Equal(t, body, string(bodyBytes))

	// Read Two times
	bodyBytes, err = getRequestBody(req)
	require.NoError(t, err)
	assert.Equal(t, body, string(bodyBytes))

	// Read Three times
	bodyBytes, err = getRequestBody(req)
	require.NoError(t, err)
	assert.Equal(t, body, string(bodyBytes))
}

type contentType struct {
	Mchid           *string    `json:"mchid"`
	Appid           *string    `json:"appid"`
	CreateTime      *time.Time `json:"create_time"`
	OutContractCode *string    `json:"out_contract_code"`
}

func (o contentType) String() string {
	ret := ""

	if o.Mchid == nil {
		ret += "Mchid:<nil>,"
	} else {
		ret += fmt.Sprintf("Mchid:%v,", *o.Mchid)
	}

	if o.Appid == nil {
		ret += "Appid:<nil>,"
	} else {
		ret += fmt.Sprintf("Appid:%v,", *o.Appid)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>,"
	} else {
		ret += fmt.Sprintf("CreateTime:%v,", *o.CreateTime)
	}

	if o.OutContractCode == nil {
		ret += "OutContractCode:<nil>,"
	} else {
		ret += fmt.Sprintf("OutContractCode:%v,", *o.OutContractCode)
	}

	return fmt.Sprintf("contentType{%s}", ret)
}

func TestHandler_ParseNotifyRequest(t *testing.T) {
	patch := gomonkey.ApplyFunc(
		time.Now, func() time.Time {
			return time.Unix(1624523846, 0)
		},
	)
	defer patch.Reset()

	const (
		mchAPIv3Key          = "testMchAPIv3Key0"
		wechatPayCertificate = `-----BEGIN CERTIFICATE-----
MIIDVzCCAj+gAwIBAgIJANfOWdH1ItcBMA0GCSqGSIb3DQEBCwUAMEIxCzAJBgNV
BAYTAlhYMRUwEwYDVQQHDAxEZWZhdWx0IENpdHkxHDAaBgNVBAoME0RlZmF1bHQg
Q29tcGFueSBMdGQwHhcNMjEwNDI3MDg1NTIzWhcNMzEwNDI1MDg1NTIzWjBCMQsw
CQYDVQQGEwJYWDEVMBMGA1UEBwwMRGVmYXVsdCBDaXR5MRwwGgYDVQQKDBNEZWZh
dWx0IENvbXBhbnkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
2VCTd91fnUn73Xy9DLvt/V62TVxRTEEstVdeRaZ3B3leO0pldE806mXO4RwdHXag
HQ4vGeZN0yqm++rDsGK+U3AH7kejyD2pXshNP9Cq5YwbptiLGtjcquw4HNxJQUOm
DeJf2vg6byms9RUipiq4SzbJKqJFlUpbuIPDpSpWz10PYmyCNeDGUUK65E5h2B83
4uxl1zNLYQCrkdBzb8oUxwYeP5a2DNxmjL5lsJML7DGr5znsevnoqGRwTm9fxCGf
y8wus7hwKz6clt3Whmmda7UAdb1c08hEQFVRbF14AR73xbnd8N0obCWJPCbzMCtk
aSef4FdEEgEXJiw0VAJT8wIDAQABo1AwTjAdBgNVHQ4EFgQUT1c7nd/SUO76HSoZ
umNUJv1R5PwwHwYDVR0jBBgwFoAUT1c7nd/SUO76HSoZumNUJv1R5PwwDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAfTjxKRQMzNB/U6ZoCUS+BSNfa2Oh
0plMN6ZuzwiVVZwg1jywvv5yv04koS7Pd4i9E4gt9ZBUQXlpq+A3oOCEEHNRR6b2
kyazGRM7s0OP5X21WrbpSmKmU6K7hkfx30yYs08LVs/Q8DIhvaj1FCFeJzUCzYn/
fHMq4tsbKO0dKAeydPM/nrUZBmaYQVKMVOORGLFjFKVO7JV6Kq/R86ouhjEPgJOe
2xulNBUcjicqtZlBdEh/PWCYP2SpGVDclKm8jeo175T3EVAkdKzzmfpxtMmnMlmq
cTJOU9TxuGvNASMtjj7pYIerTx+xgZDXEVBWFW9PjJ0TV06tCRsgSHItgg==
-----END CERTIFICATE-----`
		data = "{" +
			"\"mchid\":\"1234567890\"," +
			"\"appid\":\"054aa7d7a2a54ab5898df65bd96f001c\"," +
			"\"create_time\":\"2020-06-30T12:12:00+08:00\"," +
			"\"out_contract_code\":\"21640bdbd08e473e828f3206a2741c6e\"" +
			"}"
	)

	headers := map[string]string{
		"Content-Type":        "application/json",
		"Wechatpay-Nonce":     "EcZ9Cmy4Xyx1i6RlJQzLcCyEqDa26NBz",
		"Wechatpay-Timestamp": "1624523846",
		"Wechatpay-Serial":    "D7CE59D1F522D701",
		"Wechatpay-Signature": "tJHIiIS9eB2hAYstmAmbbD3ZE5LiIm/Ug5tuL4fC0YOFRWIHV39UFIZXC0e9Wl6lBu6sKvkqDkzpqzBsVHyXF" +
			"lbYZTOQrVdG4b6LfTnK4mikv9++ixJMd3vTf2yCqvBkh98zs3Ds5zsYQakzbcwhmw4fMJs4nPLws28H0UW9FjDR//rxELLwXvV1VEA1I" +
			"BLX70xptjL8hrfUjEE8kkry6yNJTHZRU8CAc7qHli2Ng1V1qb9ARbK8A3ThmFmPQvRGrapI/jS2laKKgYUmfdEdkNO6B2Cke5e8VTxY4" +
			"06ArAmQ90GAihDwIcb16TQMnzCMBoutnwZKNiKRACrFmtxw2Q==",
	}
	body := "{" +
		"\"id\":\"3119dfba-e649-5eec-ab1e-3412bc4d2e17\"," +
		"\"create_time\":\"2021-06-24T16:37:26+08:00\"," +
		"\"resource_type\":\"encrypt-resource\"," +
		"\"event_type\":\"PAYSCORE.USER_OPEN_SERVICE\"," +
		"\"summary\":\"签约成功\"," +
		"\"resource\":{" +
		"\"original_type\":\"payscore\"," +
		"\"algorithm\":\"AEAD_AES_256_GCM\"," +
		"\"ciphertext\":\"YDS3lKPaC4Y52Gf3uhft5qUBlIa8b428AWTtTauHQfQrRw+X1WpiuHIDy0vo1Vd6VEq67aVyqPdDYMkRVSDaZL3iZt" +
		"tevRMOoPKMifozg6XPWjIZumks/GpT48lI4NizyeaqLBokNebthah3o1H76qSlO9NkDjp9bzmKLEYYH9TEklFUpsvPqOTOcgSLgh21YJXYR" +
		"7dEBXFgRLiNIKRgO5JdXh1hccRUAlyVWxE54PXpnQ==\"," +
		"\"associated_data\":\"payscore\"," +
		"\"nonce\":\"Kj7QIyUiYx1q\"}" +
		"}"

	bodyBuf := &bytes.Buffer{}
	bodyBuf.WriteString(body)

	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1", bodyBuf)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	cert, err := utils.LoadCertificate(wechatPayCertificate)
	require.NoError(t, err)

	handler := NewNotifyHandler(
		mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(core.NewCertificateMapWithList([]*x509.Certificate{cert})),
	)

	content := new(contentType)

	notifyReq, err := handler.ParseNotifyRequest(context.Background(), req, content)
	require.NoError(t, err)

	assert.Equal(t, "3119dfba-e649-5eec-ab1e-3412bc4d2e17", notifyReq.ID)
	assert.Equal(t, "2021-06-24T16:37:26+08:00", notifyReq.CreateTime.Format(time.RFC3339))
	assert.Equal(t, "encrypt-resource", notifyReq.ResourceType)
	assert.Equal(t, "PAYSCORE.USER_OPEN_SERVICE", notifyReq.EventType)
	assert.Equal(t, "签约成功", notifyReq.Summary)

	assert.Equal(t, "payscore", notifyReq.Resource.AssociatedData)
	assert.Equal(t, "AEAD_AES_256_GCM", notifyReq.Resource.Algorithm)
	assert.Equal(t, "payscore", notifyReq.Resource.OriginalType)

	assert.Equal(t, data, notifyReq.Resource.Plaintext)

	assert.Equal(t, "1234567890", *content.Mchid)
	assert.Equal(t, "054aa7d7a2a54ab5898df65bd96f001c", *content.Appid)
	assert.Equal(t, "21640bdbd08e473e828f3206a2741c6e", *content.OutContractCode)
	createTime, _ := time.Parse(time.RFC3339, "2020-06-30T12:12:00+08:00")
	assert.Zero(t, content.CreateTime.Sub(createTime))
}
