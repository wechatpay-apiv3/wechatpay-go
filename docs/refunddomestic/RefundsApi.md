# refunddomestic/RefundsApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**Create**](#create) | **Post** /v3/refund/domestic/refunds | 退款申请
[**QueryByOutRefundNo**](#querybyoutrefundno) | **Get** /v3/refund/domestic/refunds/{out_refund_no} | 查询单笔退款（通过商户退款单号）



## Create

> Refund Create(CreateRequest)

退款申请



### 调用示例

```go
package main

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

func main() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := refunddomestic.RefundsApiService{Client: client}
	resp, result, err := svc.Create(ctx,
		refunddomestic.CreateRequest{
			SubMchid:      core.String("1900000109"),
			TransactionId: core.String("1217752501201407033233368018"),
			OutTradeNo:    core.String("1217752501201407033233368018"),
			OutRefundNo:   core.String("1217752501201407033233368018"),
			Reason:        core.String("商品已售完"),
			NotifyUrl:     core.String("https://weixin.qq.com"),
			FundsAccount:  refunddomestic.REQFUNDSACCOUNT_AVAILABLE.Ptr(),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				From: []refunddomestic.FundsFromItem{refunddomestic.FundsFromItem{
					Account: refunddomestic.ACCOUNT_AVAILABLE.Ptr(),
					Amount:  core.Int64(444),
				}},
				Refund: core.Int64(888),
				Total:  core.Int64(888),
			},
			GoodsDetail: []refunddomestic.GoodsDetail{refunddomestic.GoodsDetail{
				GoodsName:        core.String("iPhone6s 16G"),
				MerchantGoodsId:  core.String("1217752501201407033233368018"),
				RefundAmount:     core.Int64(528800),
				RefundQuantity:   core.Int64(1),
				UnitPrice:        core.Int64(528800),
				WechatpayGoodsId: core.String("1001"),
			}},
		},
	)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}
```

### 参数列表
参数名 | 参数类型 | 参数描述
------------- | ------------- | -------------
**ctx** | **context.Context** | Golang 上下文，可用于日志、请求取消、请求跟踪等功能|
**req** | [**CreateRequest**](CreateRequest.md) | API `refunddomestic` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**Refund**](Refund.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#refunddomesticrefundsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## QueryByOutRefundNo

> Refund QueryByOutRefundNo(QueryByOutRefundNoRequest)

查询单笔退款（通过商户退款单号）



### 调用示例

```go
package main

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

func main() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := refunddomestic.RefundsApiService{Client: client}
	resp, result, err := svc.QueryByOutRefundNo(ctx,
		refunddomestic.QueryByOutRefundNoRequest{
			OutRefundNo: core.String("1217752501201407033233368018"),
			SubMchid:    core.String("1900000109"),
		},
	)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}
```

### 参数列表
参数名 | 参数类型 | 参数描述
------------- | ------------- | -------------
**ctx** | **context.Context** | Golang 上下文，可用于日志、请求取消、请求跟踪等功能|
**req** | [**QueryByOutRefundNoRequest**](QueryByOutRefundNoRequest.md) | API `refunddomestic` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**Refund**](Refund.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#refunddomesticrefundsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

