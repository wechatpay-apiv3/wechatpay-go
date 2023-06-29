# retailstore/QualificationApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**LockQualification**](#lockqualification) | **Post** /v3/marketing/goods-subsidy-activity/qualification/lock | 锁定品牌加价购活动资格
[**UnlockQualification**](#unlockqualification) | **Post** /v3/marketing/goods-subsidy-activity/qualification/unlock | 解锁品牌加价购活动资格



## LockQualification

> LockQualificationResponse LockQualification(LockQualificationRequest)

锁定品牌加价购活动资格



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/retailstore"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func main() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := retailstore.QualificationApiService{Client: client}
	resp, result, err := svc.LockQualification(ctx,
		retailstore.LockQualificationRequest{
			OrderInformation: &retailstore.OrderInfo{
				PayerOpenid:     core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
				OutTradeNo:      core.String("1217752501201407033233368018"),
				TotalFee:        core.Int64(100),
				StoreId:         core.String("123"),
				StoreMerchantId: core.String("1230000109"),
			},
			QualificationIds: []string{"8495134018"},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call LockQualification err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
```

### 参数列表
参数名 | 参数类型 | 参数描述
------------- | ------------- | -------------
**ctx** | **context.Context** | Golang 上下文，可用于日志、请求取消、请求跟踪等功能|
**req** | [**LockQualificationRequest**](LockQualificationRequest.md) | API `retailstore` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**LockQualificationResponse**](LockQualificationResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#retailstorequalificationapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## UnlockQualification

> UnlockQualificationResponse UnlockQualification(UnlockQualificationRequest)

解锁品牌加价购活动资格



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/retailstore"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func main() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := retailstore.QualificationApiService{Client: client}
	resp, result, err := svc.UnlockQualification(ctx,
		retailstore.UnlockQualificationRequest{
			OrderInformation: &retailstore.OrderInfo{
				PayerOpenid:     core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
				OutTradeNo:      core.String("1217752501201407033233368018"),
				TotalFee:        core.Int64(100),
				StoreId:         core.String("123"),
				StoreMerchantId: core.String("1230000109"),
			},
			QualificationIds: []string{"8495134018"},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call UnlockQualification err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
```

### 参数列表
参数名 | 参数类型 | 参数描述
------------- | ------------- | -------------
**ctx** | **context.Context** | Golang 上下文，可用于日志、请求取消、请求跟踪等功能|
**req** | [**UnlockQualificationRequest**](UnlockQualificationRequest.md) | API `retailstore` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**UnlockQualificationResponse**](UnlockQualificationResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#retailstorequalificationapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

