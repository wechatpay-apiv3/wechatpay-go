//Package validators 微信支付api v3 go http-client 回包校验器
package validators

import (
	"bytes"
	"context"
	"encoding/base64"
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

// WechatPayValidator 回包校验器
type WechatPayValidator struct {
	Verifier auth.Verifier // 验证器
}

// Validate 使用验证器对回包进行校验
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
	signature, err := base64.StdEncoding.DecodeString(strings.TrimSpace(response.Header.Get(consts.WechatPaySignature)))
	if err != nil {
		return fmt.Errorf("base64 decode string wechat pay signature err:%s", err.Error())
	}
	err = validator.Verifier.Verify(ctx, serialNumber, message, string(signature))
	if err != nil {
		return fmt.Errorf("validate verify fail serial=%s request-id=%s err=%s", serialNumber,
			strings.TrimSpace(response.Header.Get(consts.WechatPaySerial)), err)
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
