// Copyright 2024
//
// 付款码支付
//
// 付款码支付API
//
// API version: 1.2.3

// 文档地址： https://pay.weixin.qq.com/docs/merchant/apis/code-payment-v3/direct/code-pay.html
package code

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
	nethttp "net/http"
	neturl "net/url"
	"strings"
)

type CodeApiService services.Service

// 撤销订单
// 支付交易返回失败或支付系统超时，调用该接口撤销交易。如果此订单用户支付失败，微信支付系统会将此订单关闭；如果用户支付成功，微信支付系统会将此订单资金退还给用户。
func (a *CodeApiService) CloseOrder(ctx context.Context, req CloseOrderRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.Appid == nil {
		return nil, fmt.Errorf("field `OutTradeNo` is required and must be specified in CloseOrderRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/out-trade-no/{out_trade_no}/reverse"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutTradeNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &CloseRequest{
		Mchid: req.Mchid,
		Appid: req.Appid,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// Prepay 付款码支付下单
//
// 收银员使用扫码设备读取微信用户付款码以后，二维码或条码信息会传送至商户收银台，由商户收银台或者商户后台调用该接口发起支付。
func (a *CodeApiService) Prepay(ctx context.Context, req PrepayRequest) (resp *PrepayResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/pay/transactions/codepay"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(PrepayResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
