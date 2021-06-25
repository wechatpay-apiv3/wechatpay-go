# certificates/CertificatesApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**DownloadCertificates**](#downloadcertificates) | **Get** /v3/certificates | 获取平台证书列表



## DownloadCertificates

> DownloadCertificatesResponse DownloadCertificates()

获取平台证书列表



### 调用示例

```go
package main

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
)

func main() {
	var (
		ctx    context.Context
		client *core.Client
	)
	// 假设已获得初始化后的 core.Client

	svc := certificates.CertificatesApiService{Client: client}
	resp, result, err = svc.DownloadCertificates(ctx)

	// TODO: 处理返回结果
	_, _, _ = resp, result, err
}
```

### 参数列表
参数名 | 参数类型 | 参数描述
------------- | ------------- | -------------
**ctx** | **context.Context** | Golang 上下文，可用于日志、请求取消、请求跟踪等功能|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**DownloadCertificatesResponse**](DownloadCertificatesResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#certificatescertificatesapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

