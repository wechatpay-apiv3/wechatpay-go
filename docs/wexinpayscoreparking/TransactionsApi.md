# wexinpayscoreparking/TransactionsApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**CreateTransaction**](#createtransaction) | **Post** /v3/vehicle/transactions/parking | 扣费受理
[**QueryTransaction**](#querytransaction) | **Get** /v3/vehicle/transactions/out-trade-no/{out_trade_no} | 查询订单



## CreateTransaction

> Transaction CreateTransaction(CreateTransactionRequest)

扣费受理



### 调用示例

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/wexinpayscoreparking"
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

	svc := wexinpayscoreparking.TransactionsApiService{Client: client}
	resp, result, err := svc.CreateTransaction(ctx,
		wexinpayscoreparking.CreateTransactionRequest{
			Appid:         core.String("wxcbda96de0b165486"),
			SubAppid:      core.String("wxcbda96de0b165486"),
			SubMchid:      core.String("1900000109"),
			Description:   core.String("停车场扣费"),
			Attach:        core.String("深圳分店"),
			OutTradeNo:    core.String("20150806125346"),
			TradeScene:    core.String("PARKING"),
			GoodsTag:      core.String("WXG"),
			NotifyUrl:     core.String("https://yoursite.com/wxpay.html"),
			ProfitSharing: core.String("Y"),
			Amount: &wexinpayscoreparking.OrderAmount{
				Total:    core.Int64(888),
				Currency: core.String("CNY"),
			},
			ParkingInfo: &wexinpayscoreparking.ParkingTradeScene{
				ParkingId:        core.String("5K8264ILTKCH16CQ250"),
				PlateNumber:      core.String("粤B888888"),
				PlateColor:       wexinpayscoreparking.PLATECOLOR_BLUE.Ptr(),
				StartTime:        core.Time(time.Now()),
				EndTime:          core.Time(time.Now()),
				ParkingName:      core.String("欢乐海岸停车场"),
				ChargingDuration: core.Int64(3600),
				DeviceId:         core.String("12313"),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateTransaction err:%s", err)
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
**req** | [**CreateTransactionRequest**](CreateTransactionRequest.md) | API `wexinpayscoreparking` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**Transaction**](Transaction.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#wexinpayscoreparkingtransactionsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## QueryTransaction

> Transaction QueryTransaction(QueryTransactionRequest)

查询订单



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/wexinpayscoreparking"
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

	svc := wexinpayscoreparking.TransactionsApiService{Client: client}
	resp, result, err := svc.QueryTransaction(ctx,
		wexinpayscoreparking.QueryTransactionRequest{
			SubMchid:   core.String("1900000109"),
			OutTradeNo: core.String("20150806125346"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryTransaction err:%s", err)
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
**req** | [**QueryTransactionRequest**](QueryTransactionRequest.md) | API `wexinpayscoreparking` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**Transaction**](Transaction.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#wexinpayscoreparkingtransactionsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

