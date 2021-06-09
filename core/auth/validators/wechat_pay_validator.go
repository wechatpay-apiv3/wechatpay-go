// Package validators 微信支付 API v3 Go SDK 应答报文验证器
package validators

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

// WechatPayValidator 微信支付 API v3 默认应答报文验证器
type WechatPayValidator struct {
	Verifier auth.Verifier // 数字签名验证器
}

// Validate 使用验证器对微信支付应答报文进行验证
func (validator *WechatPayValidator) Validate(ctx context.Context, response *http.Response) error {
	if validator.Verifier == nil {
		return fmt.Errorf("you must init WechatPayValidator with auth.Verifier")
	}
	err := validateParameters(response)
	if err != nil {
		return err
	}
	message, err := buildMessage(response)
	if err != nil {
		return err
	}
	serialNumber := strings.TrimSpace(response.Header.Get(consts.WechatPaySerial))
	signature := response.Header.Get(consts.WechatPaySignature)
	err = validator.Verifier.Verify(ctx, serialNumber, message, signature)
	if err != nil {
		return fmt.Errorf("validate verify fail serial=%s request-id=%s err=%s", serialNumber,
			strings.TrimSpace(response.Header.Get(consts.RequestID)), err)
	}
	return nil
}

func validateParameters(response *http.Response) (err error) {
	requestID := strings.TrimSpace(response.Header.Get(consts.RequestID))
	if requestID == "" {
		return fmt.Errorf("empty %s", consts.RequestID)
	}
	if strings.TrimSpace(response.Header.Get(consts.WechatPaySerial)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPaySerial, requestID)
	}
	if strings.TrimSpace(response.Header.Get(consts.WechatPaySignature)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPaySignature, requestID)
	}
	if strings.TrimSpace(response.Header.Get(consts.WechatPayTimestamp)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPayTimestamp, requestID)
	}
	if strings.TrimSpace(response.Header.Get(consts.WechatPayNonce)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPayNonce, requestID)
	}
	timeStampStr := strings.TrimSpace(response.Header.Get(consts.WechatPayTimestamp))
	timeStamp, err := strconv.Atoi(timeStampStr)
	if err != nil {
		return fmt.Errorf("invalid timestamp:[%s] request-id=[%s] err:[%v]", timeStampStr, requestID, err)
	}
	if math.Abs(float64(timeStamp)-float64(time.Now().Unix())) >= consts.FiveMinute {
		return fmt.Errorf("timestamp=[%d] expires, request-id=[%s]", timeStamp, requestID)
	}
	return nil
}

func buildMessage(response *http.Response) (message string, err error) {
	timeStamp := strings.TrimSpace(response.Header.Get(consts.WechatPayTimestamp))
	nonce := strings.TrimSpace(response.Header.Get(consts.WechatPayNonce))
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("read response body err:[%s]", err.Error())
	}
	body := string(bodyBytes)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	message = fmt.Sprintf("%s\n%s\n%s\n", timeStamp, nonce, body)
	return message, nil
}
