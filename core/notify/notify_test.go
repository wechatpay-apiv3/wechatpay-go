package notify

import (
	"bytes"
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"

	"github.com/agiledragon/gomonkey"
)

func Test_getRequestBody(t *testing.T) {
	body := "fake req body"
	bodyBuf := &bytes.Buffer{}
	bodyBuf.WriteString(body)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", bodyBuf)
	require.NoError(t, err)

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
	ContractId      *string    `json:"contract_id"`
	Mchid           *string    `json:"mchid"`
	Appid           *string    `json:"appid"`
	Openid          *string    `json:"openid"`
	PlanId          *string    `json:"plan_id"`
	ContractStatus  *string    `json:"contract_status"`
	CreateTime      *time.Time `json:"create_time"`
	OutContractCode *string    `json:"out_contract_code"`
}

func (o contentType) String() string {
	ret := ""

	if o.ContractId == nil {
		ret += "ContractId:<nil>,"
	} else {
		ret += fmt.Sprintf("ContractId:%v,", *o.ContractId)
	}

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

	if o.Openid == nil {
		ret += "Openid:<nil>,"
	} else {
		ret += fmt.Sprintf("Openid:%v,", *o.Openid)
	}

	if o.PlanId == nil {
		ret += "PlanId:<nil>,"
	} else {
		ret += fmt.Sprintf("PlanId:%v,", *o.PlanId)
	}

	if o.ContractStatus == nil {
		ret += "ContractStatus:<nil>,"
	} else {
		ret += fmt.Sprintf("ContractStatus:%v,", *o.ContractStatus)
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
	patch := gomonkey.ApplyFunc(time.Now, func() time.Time {
		return time.Unix(1624523846, 0)
	})
	defer patch.Reset()

	headers := map[string]string{
		"Content-Type":        "application/json",
		"Request-Id":          "0885F2CF8606108F0518E29E944820F10B28E24A",
		"Wechatpay-Nonce":     "EcZ9Cmy4Xyx1i6RlJQzLcCyEqDa26NBz",
		"Wechatpay-Timestamp": "1624523846",
		"Wechatpay-Serial":    "714E57301CE0FD6267DA19B19E08CEE431E6D943",
		"Wechatpay-Signature": "E/RMTykUevOG5ZMjk7z+pm/EpVGM9VbmTlhVm+7R4cu/wd6Dtk2jKgzgBSAwL6wL8484lz9V5Xi8n2FO6oDAstwy/YOBpHdB7PH/Za/WUN8Op2CGpxdZNepn+PLWor/TKhSm9cb0owg1m2yNvEdHmtrDe3hD8IHJAWRovdBybc+41YLJuFSh5gRD7OtLyYsj3w7Zrac+CARdr8iF0QTvLhTv7rN45Tv59g/n9TpNxSnvj4nHb74L6EhjCwN17rnq9IR7wxfnOtJHVbDaCOm5xy4pWA+d3uvJrF/SfE0FmBhn1KQTlKfveGAiatSXl6qrIwWmjaiqcsTZy7ZEP1M/Wg==",
	}
	body := `{"id":"3119dfba-e649-5eec-ab1e-3412bc4d2e17","create_time":"2021-06-24T16:37:26+08:00","resource_type":"encrypt-resource","event_type":"PAYSCORE.USER_OPEN_SERVICE","summary":"签约成功","resource":{"original_type":"payscore","algorithm":"AEAD_AES_256_GCM","ciphertext":"857DUp4FqJhj2AxF63p333erLhhcG2glNoAKeP3ArRC3RbZyL4INrdmqzduerALpJAJ7mF0rYN0gbQBEYc89N3wvExHJsnre80AhBQydIF5IhX5xuzjlw682dC5aP317H2UyDOCobQHSJsTyIoty/2AY30q/Cc7yz4TBVgNAgMMLFWAsrq+T/oguNQxmxKglPgxTEXM5N42UdLPYwOmy0EZZqOT5Lt5ppwnSJrJPX7qGha5Ba27VmOgNf5u15WVlEzgY1ryba4ijLmQCix70Pgnq45emUT8iGMrXRpJcM9vnMxFjgabI9prlH0lmruEZd0o4FFMbRIESGfPYvRyq7qDMp+WiRXnUOkso8Dw4F9GBNTK1XgukPPgPLblHro2O9Z+7HUfizh/9KDa/7JwdhJX58r5NmV4wZEyootmjOFA4jNetfEvUccXLpDO82/pUZ5IhSw==","associated_data":"payscore","nonce":"Kj7QIyUiYx1q"}}`

	bodyBuf := &bytes.Buffer{}
	bodyBuf.WriteString(body)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", bodyBuf)
	require.NoError(t, err)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	cert, err := utils.LoadCertificate(consts.WechatPayCertificate)
	require.NoError(t, err)

	handler := NewNotifyHandler(consts.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(core.NewCertificateMapWithList([]*x509.Certificate{cert})))

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

	t.Log(content)
	t.Log(notifyReq.Resource.Plaintext)
}
