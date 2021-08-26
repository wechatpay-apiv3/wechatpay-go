# profitsharing/MerchantsApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**QueryMerchantRatio**](#querymerchantratio) | **Get** /v3/profitsharing/merchant-configs/{sub_mchid} | 查询最大分账比例API



## QueryMerchantRatio

> QueryMerchantRatioResponse QueryMerchantRatio(QueryMerchantRatioRequest)

查询最大分账比例API



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

	svc := profitsharing.MerchantsApiService{Client: client}
	resp, result, err := svc.QueryMerchantRatio(ctx,
		profitsharing.QueryMerchantRatioRequest{
			SubMchid: core.String("1900000109"),
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
**req** | [**QueryMerchantRatioRequest**](QueryMerchantRatioRequest.md) | API `profitsharing` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**QueryMerchantRatioResponse**](QueryMerchantRatioResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#profitsharingmerchantsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

