package validators

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

type wechatPayValidator struct {
	verifier auth.Verifier
}

func (v *wechatPayValidator) validateHTTPMessage(ctx context.Context, header http.Header, body []byte) error {
	if v.verifier == nil {
		return fmt.Errorf("you must init Validator with auth.Verifier")
	}

	if err := checkParameters(ctx, header, body); err != nil {
		return err
	}

	message, err := buildMessage(ctx, header, body)
	if err != nil {
		return err
	}

	serialNumber := header.Get(consts.WechatPaySerial)
	signature := header.Get(consts.WechatPaySignature)
	err = v.verifier.Verify(ctx, serialNumber, message, signature)
	if err != nil {
		return fmt.Errorf(
			"validate verify fail serial=%s request-id=%s err=%s",
			serialNumber, header.Get(consts.RequestID), err,
		)
	}
	return nil
}

func checkParameters(ctx context.Context, header http.Header, body []byte) error {
	// Suppressing warnings
	_ = ctx
	_ = body

	requestID := strings.TrimSpace(header.Get(consts.RequestID))
	if requestID == "" {
		return fmt.Errorf("empty %s", consts.RequestID)
	}

	if strings.TrimSpace(header.Get(consts.WechatPaySerial)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPaySerial, requestID)
	}

	if strings.TrimSpace(header.Get(consts.WechatPaySignature)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPaySignature, requestID)
	}

	if strings.TrimSpace(header.Get(consts.WechatPayTimestamp)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPayTimestamp, requestID)
	}

	if strings.TrimSpace(header.Get(consts.WechatPayNonce)) == "" {
		return fmt.Errorf("empty %s, request-id=[%s]", consts.WechatPayNonce, requestID)
	}

	timeStampStr := strings.TrimSpace(header.Get(consts.WechatPayTimestamp))
	timeStamp, err := strconv.Atoi(timeStampStr)
	if err != nil {
		return fmt.Errorf("invalid timestamp:[%s] request-id=[%s] err:[%v]", timeStampStr, requestID, err)
	}

	if math.Abs(float64(timeStamp)-float64(time.Now().Unix())) >= consts.FiveMinute {
		return fmt.Errorf("timestamp=[%d] expires, request-id=[%s]", timeStamp, requestID)
	}
	return nil
}

func buildMessage(ctx context.Context, header http.Header, body []byte) (string, error) {
	timeStamp := header.Get(consts.WechatPayTimestamp)
	nonce := header.Get(consts.WechatPayNonce)

	message := fmt.Sprintf("%s\n%s\n%s\n", timeStamp, nonce, string(body))
	return message, nil
}
