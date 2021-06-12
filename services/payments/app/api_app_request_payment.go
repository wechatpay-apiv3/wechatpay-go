package app

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
	PrepayId *string `json:"prepayId"`
	// 商户号
	PartnerId *string `json:"partnerId"`
	// 时间戳
	TimeStamp *string `json:"timeStamp"`
	// 随机字符串
	NonceStr *string `json:"nonceStr"`
	// 订单详情扩展字符串
	Package *string `json:"package"`
	// 签名
	Sign *string `json:"sign"`
}

func (a *AppApiService) PrepayWithRequestPayment(
	ctx context.Context,
	req PrepayRequest,
) (resp *PrepayWithRequestPaymentResponse, result *core.APIResult, err error) {
	prepayResp, result, err := a.Prepay(ctx, req)
	if err != nil {
		return nil, result, err
	}

	resp = new(PrepayWithRequestPaymentResponse)
	resp.PrepayId = prepayResp.PrepayId
	resp.TimeStamp = core.String(strconv.FormatInt(time.Now().Unix(), 10))
	nonce, err := utils.GenerateNonce()
	if err != nil {
		return nil, nil, fmt.Errorf("generate request for payment err:%s", err.Error())
	}
	resp.NonceStr = core.String(nonce)
	resp.Package = core.String("Sign=WXPay")
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n", *req.Appid, *resp.TimeStamp, *resp.NonceStr, *prepayResp.PrepayId)
	signatureResult, err := a.Client.Sign(ctx, message)
	if err != nil {
		return nil, nil, fmt.Errorf("generate sign for payment err:%s", err.Error())
	}
	resp.Sign = core.String(signatureResult.Signature)
	resp.PartnerId = core.String(signatureResult.MchID)
	return resp, result, nil
}

func (o PrepayWithRequestPaymentResponse) String() string {
	var ret string
	if o.PrepayId == nil {
		ret += "PrepayId:<nil>, "
	} else {
		ret += fmt.Sprintf("PrepayId:%v, ", *o.PrepayId)
	}
	if o.PartnerId == nil {
		ret += "PartnerId:<nil>, "
	} else {
		ret += fmt.Sprintf("PartnerId:%v, ", *o.PartnerId)
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
	if o.Sign == nil {
		ret += "Sign:<nil>"
	} else {
		ret += fmt.Sprintf("Sign:%v", *o.Sign)
	}

	return fmt.Sprintf("PrepayResponse{%s}", ret)
}
