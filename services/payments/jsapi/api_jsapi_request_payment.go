package jsapi

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// PrepayWithRequestPaymentResponse 预下单ID，并包含了调起支付的请求参数
type PrepayWithRequestPaymentResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"`
	// 应用ID
	Appid *string `json:"appId"`
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名方式
	SignType *string `json:"signType"`
	// 签名
	PaySign *string `json:"paySign"`
}

// PrepayWithRequestPayment Jsapi支付下单，并返回调起支付的请求参数
func (a *JsapiApiService) PrepayWithRequestPayment(
	ctx context.Context, req PrepayRequest) (resp *PrepayWithRequestPaymentResponse, result *core.APIResult, err error) {
	prepayResp, result, err := a.Prepay(ctx, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(PrepayWithRequestPaymentResponse)
	resp.PrepayId = prepayResp.PrepayId
	resp.SignType = core.String("RSA")
	resp.Appid = req.Appid
	resp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, nil, fmt.Errorf("generate request for payment err:%s", err.Error())
	}
	resp.NonceStr = core.String(nonce)
	resp.Package = core.String("prepay_id=" + *prepayResp.PrepayId)
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *resp.Appid, *resp.TimeStamp, *resp.NonceStr, *resp.Package)
	signatureResult, err := a.Client.Sign(ctx, message)
	if err != nil {
		return nil, nil, fmt.Errorf("generate sign for payment err:%s", err.Error())
	}
	resp.PaySign = core.String(signatureResult.Signature)
	return resp, result, nil
}

func (o PrepayWithRequestPaymentResponse) String() string {
	var ret string
	if o.PrepayId == nil {
		ret += "PrepayId:<nil>, "
	} else {
		ret += fmt.Sprintf("PrepayId:%v, ", *o.PrepayId)
	}
	if o.Appid == nil {
		ret += "Appid:<nil>, "
	} else {
		ret += fmt.Sprintf("Appid:%v, ", *o.Appid)
	}
	if o.TimeStamp == nil {
		ret += "TimeStamp:<nil>, "
	} else {
		ret += fmt.Sprintf("TimeStamp:%v, ", *o.TimeStamp)
	}
	if o.NonceStr == nil {
		ret += "NonceStr:<nil>, "
	} else {
		ret += fmt.Sprintf("NonceStr:%v, ", *o.NonceStr)
	}
	if o.Package == nil {
		ret += "Package:<nil>, "
	} else {
		ret += fmt.Sprintf("Package:%v, ", *o.Package)
	}
	if o.SignType == nil {
		ret += "SignType:<nil>, "
	} else {
		ret += fmt.Sprintf("SignType:%v, ", *o.SignType)
	}
	if o.PaySign == nil {
		ret += "PaySign:<nil>"
	} else {
		ret += fmt.Sprintf("PaySign:%v", *o.PaySign)
	}

	return fmt.Sprintf("PrepayResponse{%s}", ret)
}
