# profitsharing/OrdersApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**CreateOrder**](#createorder) | **Post** /v3/profitsharing/orders | 请求分账API
[**QueryOrder**](#queryorder) | **Get** /v3/profitsharing/orders/{out_order_no} | 查询分账结果API
[**UnfreezeOrder**](#unfreezeorder) | **Post** /v3/profitsharing/orders/unfreeze | 解冻剩余资金API



## CreateOrder

> OrdersEntity CreateOrder(CreateOrderRequest)

请求分账API



### 调用示例

```go
package main

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
)

func main() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.CreateOrder(ctx,
		profitsharing.CreateOrderRequest{
			Appid:      core.String("wx8888888888888888"),
			OutOrderNo: core.String("P20150806125346"),
			Receivers: []profitsharing.CreateOrderReceiver{profitsharing.CreateOrderReceiver{
				Account:     core.String("86693852"),
				Amount:      core.Int64(888),
				Description: core.String("分给商户A"),
				Name:        core.String("hu89ohu89ohu89o"),
				Type:        core.String("MERCHANT_ID"),
			}},
			SubAppid:        core.String("wx8888888888888889"),
			SubMchid:        core.String("1900000109"),
			TransactionId:   core.String("4208450740201411110007820472"),
			UnfreezeUnsplit: core.Bool(true),
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
**req** | [**CreateOrderRequest**](CreateOrderRequest.md) | API `profitsharing` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**OrdersEntity**](OrdersEntity.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#profitsharingordersapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## QueryOrder

> OrdersEntity QueryOrder(QueryOrderRequest)

查询分账结果API



### 调用示例

```go
package main

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
)

func main() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.QueryOrder(ctx,
		profitsharing.QueryOrderRequest{
			TransactionId: core.String("4208450740201411110007820472"),
			OutOrderNo:    core.String("P20150806125346"),
			SubMchid:      core.String("1900000109"),
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
**req** | [**QueryOrderRequest**](QueryOrderRequest.md) | API `profitsharing` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**OrdersEntity**](OrdersEntity.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#profitsharingordersapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## UnfreezeOrder

> OrdersEntity UnfreezeOrder(UnfreezeOrderRequest)

解冻剩余资金API



### 调用示例

```go
package main

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/profitsharing"
)

func main() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := profitsharing.OrdersApiService{Client: client}
	resp, result, err := svc.UnfreezeOrder(ctx,
		profitsharing.UnfreezeOrderRequest{
			Description:   core.String("解冻全部剩余资金"),
			OutOrderNo:    core.String("P20150806125346"),
			SubMchid:      core.String("1900000109"),
			TransactionId: core.String("4208450740201411110007820472"),
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
**req** | [**UnfreezeOrderRequest**](UnfreezeOrderRequest.md) | API `profitsharing` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**OrdersEntity**](OrdersEntity.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#profitsharingordersapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

