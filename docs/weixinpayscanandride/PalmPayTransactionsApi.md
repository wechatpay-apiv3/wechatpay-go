# weixinpayscanandride/PalmPayTransactionsApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**CreatePalmPayTransaction**](#createpalmpaytransaction) | **Post** /v3/transit-palm-pay/transactions | 直连商户扣费受理—刷掌地铁
[**QueryPalmPayTransaction**](#querypalmpaytransaction) | **Get** /v3/transit-palm-pay/transactions/{out_trade_no} | 直连商户查询订单—刷掌地铁



## CreatePalmPayTransaction

> PalmPayTransaction CreatePalmPayTransaction(CreatePalmPayTransactionRequest)

直连商户扣费受理—刷掌地铁



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/weixinpayscanandride"
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

	svc := weixinpayscanandride.PalmPayTransactionsApiService{Client: client}
	resp, result, err := svc.CreatePalmPayTransaction(ctx,
		weixinpayscanandride.CreatePalmPayTransactionRequest{
			Appid:       core.String("wxcbda96de0b165486"),
			Description: core.String("地铁扣费"),
			Attach:      core.String("深圳分店"),
			OutTradeNo:  core.String("20150806125346"),
			GoodsTag:    core.String("WXG"),
			PlanId:      core.String("123"),
			Openid:      core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
			NotifyUrl:   core.String("https://yoursite.com/wxpay.html"),
			TradeScene:  weixinpayscanandride.PALMPAYTRADESCENE_PALM_METRO.Ptr(),
			MetroInfo: &weixinpayscanandride.PalmMetroSceneInfo{
				StartTime:        core.String("2017-08-26T10:43:39+08:00"),
				EndTime:          core.String("2017-08-26T10:43:39+08:00"),
				StartStation:     core.String("西单"),
				EndStation:       core.String("天安门西"),
				DeductCredential: core.String("0885E8959C06102118D9A28848200028979B01"),
			},
			Amount: &weixinpayscanandride.OrderAmount{
				Total:    core.Int64(600),
				Currency: core.String("CNY"),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreatePalmPayTransaction err:%s", err)
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
**req** | [**CreatePalmPayTransactionRequest**](CreatePalmPayTransactionRequest.md) | API `weixinpayscanandride` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**PalmPayTransaction**](PalmPayTransaction.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#weixinpayscanandridepalmpaytransactionsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## QueryPalmPayTransaction

> PalmPayTransaction QueryPalmPayTransaction(QueryPalmPayTransactionRequest)

直连商户查询订单—刷掌地铁



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/weixinpayscanandride"
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

	svc := weixinpayscanandride.PalmPayTransactionsApiService{Client: client}
	resp, result, err := svc.QueryPalmPayTransaction(ctx,
		weixinpayscanandride.QueryPalmPayTransactionRequest{
			OutTradeNo: core.String("20150806125346"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryPalmPayTransaction err:%s", err)
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
**req** | [**QueryPalmPayTransactionRequest**](QueryPalmPayTransactionRequest.md) | API `weixinpayscanandride` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**PalmPayTransaction**](PalmPayTransaction.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#weixinpayscanandridepalmpaytransactionsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

