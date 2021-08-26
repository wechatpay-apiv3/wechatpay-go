# profitsharing/ReceiversApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**AddReceiver**](#addreceiver) | **Post** /v3/profitsharing/receivers/add | 添加分账接收方API
[**DeleteReceiver**](#deletereceiver) | **Post** /v3/profitsharing/receivers/delete | 删除分账接收方API



## AddReceiver

> AddReceiverResponse AddReceiver(AddReceiverRequest)

添加分账接收方API



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

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, result, err := svc.AddReceiver(ctx,
		profitsharing.AddReceiverRequest{
			Account:        core.String("86693852"),
			Appid:          core.String("wx8888888888888888"),
			CustomRelation: core.String("代理商"),
			Name:           core.String("hu89ohu89ohu89o"),
			RelationType:   profitsharing.RECEIVERRELATIONTYPE_SERVICE_PROVIDER.Ptr(),
			SubAppid:       core.String("wx8888888888888889"),
			SubMchid:       core.String("1900000109"),
			Type:           profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
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
**req** | [**AddReceiverRequest**](AddReceiverRequest.md) | API `profitsharing` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**AddReceiverResponse**](AddReceiverResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#profitsharingreceiversapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## DeleteReceiver

> DeleteReceiverResponse DeleteReceiver(DeleteReceiverRequest)

删除分账接收方API



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

	svc := profitsharing.ReceiversApiService{Client: client}
	resp, result, err := svc.DeleteReceiver(ctx,
		profitsharing.DeleteReceiverRequest{
			Account:  core.String("86693852"),
			Appid:    core.String("wx8888888888888888"),
			SubAppid: core.String("wx8888888888888889"),
			SubMchid: core.String("1900000109"),
			Type:     profitsharing.RECEIVERTYPE_MERCHANT_ID.Ptr(),
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
**req** | [**DeleteReceiverRequest**](DeleteReceiverRequest.md) | API `profitsharing` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**DeleteReceiverResponse**](DeleteReceiverResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#profitsharingreceiversapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

