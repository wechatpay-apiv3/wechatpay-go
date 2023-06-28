# cashcoupons/PlatformFavorApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**CreatePlatformFavor**](#createplatformfavor) | **Post** /v3/marketing/favor/platform-favor | 创建立减金批次
[**GetPlatformFavor**](#getplatformfavor) | **Get** /v3/marketing/favor/platform-favor/{stock_id} | 查询立减金批次信息



## CreatePlatformFavor

> PlatformFavorStockEntity CreatePlatformFavor(CreatePlatformFavorRequest)

创建立减金批次



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/cashcoupons"
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

	svc := cashcoupons.PlatformFavorApiService{Client: client}
	resp, result, err := svc.CreatePlatformFavor(ctx,
		cashcoupons.CreatePlatformFavorRequest{
			BelongMerchant: core.String("98568865"),
			SendRule: &cashcoupons.PlatformFavorSendRule{
				MaxCoupons:          core.Int64(9000),
				MaxAmount:           core.Int64(10000),
				MaxCouponsByDay:     core.Int64(100),
				MaxCouponsPerUser:   core.Int64(3),
				NaturalPersonLimit:  core.Bool(true),
				PreventApiAbuse:     core.Bool(false),
				DeductBalanceMethod: cashcoupons.DEDUCTBALANCEMETHOD_BATCH_DEDUCT.Ptr(),
			},
			UseRule: &cashcoupons.PlatformFavorUseRule{
				AvailableTime: &cashcoupons.PlatformFavorAvailableTime{
					AvailableBeginTime: core.String("2015-05-20T13:29:35.120+08:00"),
					AvailableEndTime:   core.String("2015-05-20T13:29:35.120+08:00"),
				},
				CouponAmount:       core.Int64(15),
				TransactionMinimum: core.Int64(100),
			},
			DisplayPattern: &cashcoupons.PlatformFavorDisplayPattern{
				Description: core.String("使用微信支付消费时自动抵扣。单笔消费满xx元以上（含xx元）可一次性全额使用立减金，不找零，不重复使用。生活缴费、理财等少数特定商户不可使用，有问题请致电95017"),
			},
			OutRequestNo: core.String("mch_create_202003080999"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreatePlatformFavor err:%s", err)
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
**req** | [**CreatePlatformFavorRequest**](CreatePlatformFavorRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**PlatformFavorStockEntity**](PlatformFavorStockEntity.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsplatformfavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## GetPlatformFavor

> PlatformFavorStockEntity GetPlatformFavor(GetPlatformFavorRequest)

查询立减金批次信息



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/cashcoupons"
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

	svc := cashcoupons.PlatformFavorApiService{Client: client}
	resp, result, err := svc.GetPlatformFavor(ctx,
		cashcoupons.GetPlatformFavorRequest{
			StockId:           core.String("1263456"),
			StockCreatorMchid: core.String("987654"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetPlatformFavor err:%s", err)
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
**req** | [**GetPlatformFavorRequest**](GetPlatformFavorRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**PlatformFavorStockEntity**](PlatformFavorStockEntity.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsplatformfavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

