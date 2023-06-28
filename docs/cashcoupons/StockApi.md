# cashcoupons/StockApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**CreateCouponStock**](#createcouponstock) | **Post** /v3/marketing/favor/coupon-stocks | 创建代金券批次
[**CreateNocashSingleitemStock**](#createnocashsingleitemstock) | **Post** /v3/marketing/favor/nocash-singleitem-stocks | 创建无资金流单品批次
[**CreateNocashStock**](#createnocashstock) | **Post** /v3/marketing/favor/nocash-stocks | 创建无资金流全场批次
[**ListAvailableMerchants**](#listavailablemerchants) | **Get** /v3/marketing/favor/stocks/{stock_id}/merchants | 查询代金券可用商户
[**ListAvailableSingleitems**](#listavailablesingleitems) | **Get** /v3/marketing/favor/stocks/{stock_id}/items | 查询可核销商品编码
[**ListStocks**](#liststocks) | **Get** /v3/marketing/favor/stocks | 条件查询批次列表
[**ModifyAvailableMerchants**](#modifyavailablemerchants) | **Post** /v3/marketing/favor/stocks/{stock_id}/modify-merchants | 修改批次可用商户号
[**ModifyAvailableSingleitems**](#modifyavailablesingleitems) | **Post** /v3/marketing/favor/stocks/{stock_id}/modify-goods-id | 修改可核销商品编码
[**ModifyBudget**](#modifybudget) | **Post** /v3/marketing/favor/stocks/{stock_id}/max-amount-modify | 修改批次预算金额
[**ModifyStock**](#modifystock) | **Patch** /v3/marketing/favor/stocks/{stock_id} | 修改批次基本信息
[**PauseStock**](#pausestock) | **Post** /v3/marketing/favor/stocks/{stock_id}/pause | 暂停批次
[**QueryStock**](#querystock) | **Get** /v3/marketing/favor/stocks/{stock_id} | 查询批次详情
[**RefundFlow**](#refundflow) | **Get** /v3/marketing/favor/stocks/{stock_id}/refund-flow | 下载批次退款明细
[**RestartStock**](#restartstock) | **Post** /v3/marketing/favor/stocks/{stock_id}/restart | 重启批次
[**StartStock**](#startstock) | **Post** /v3/marketing/favor/stocks/{stock_id}/start | 激活开启批次
[**StopStock**](#stopstock) | **Post** /v3/marketing/favor/stocks/{stock_id}/stop | 终止批次
[**UseFlow**](#useflow) | **Get** /v3/marketing/favor/stocks/{stock_id}/use-flow | 下载批次核销明细



## CreateCouponStock

> CreateCouponStockResponse CreateCouponStock(CreateCouponStockRequest)

创建代金券批次



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.CreateCouponStock(ctx,
		cashcoupons.CreateCouponStockRequest{
			StockName:          core.String("微信支付代金券批次"),
			Comment:            core.String("零售批次"),
			BelongMerchant:     core.String("98568865"),
			AvailableBeginTime: core.String("2015-05-20T13:29:35.120+08:00"),
			AvailableEndTime:   core.String("2015-05-20T13:29:35.120+08:00"),
			StockUseRule: &cashcoupons.StockRule{
				MaxCoupons:         core.Int64(100),
				MaxAmount:          core.Int64(5000),
				MaxAmountByDay:     core.Int64(400),
				MaxCouponsPerUser:  core.Int64(3),
				NaturalPersonLimit: core.Bool(false),
				PreventApiAbuse:    core.Bool(false),
			},
			PatternInfo: &cashcoupons.PatternInfo{
				Description:      core.String("微信支付营销代金券"),
				MerchantLogo:     core.String("CDN地址"),
				MerchantName:     core.String("微信支付"),
				BackgroundColor:  cashcoupons.BACKGROUNDCOLOR_COLOR010.Ptr(),
				CouponImage:      core.String("图片cdn地址"),
				JumpTarget:       cashcoupons.JUMPTARGET_PAYMENT_CODE.Ptr(),
				MiniProgramAppid: core.String("wx23232232323"),
				MiniProgramPath:  core.String("/path/index/index"),
			},
			CouponUseRule: &cashcoupons.CouponRule{
				CouponAvailableTime: &cashcoupons.FavorAvailableTime{
					FixAvailableTime: &cashcoupons.FixedAvailableTime{
						AvailableWeekDay: []int64{int64(1)},
						BeginTime:        core.Int64(0),
						EndTime:          core.Int64(3600),
					},
					SecondDayAvailable:        core.Bool(false),
					AvailableTimeAfterReceive: core.Int64(1440),
				},
				FixedNormalCoupon: &cashcoupons.FixedValueStockMsg{
					CouponAmount:       core.Int64(100),
					TransactionMinimum: core.Int64(100),
				},
				GoodsTag:           []string{"123321"},
				TradeType:          []cashcoupons.TradeType{cashcoupons.TRADETYPE_MICROAPP},
				CombineUse:         core.Bool(false),
				AvailableItems:     []string{"123321"},
				UnavailableItems:   []string{"789987"},
				AvailableMerchants: []string{"9856000"},
				LimitCard: &cashcoupons.CardLimitation{
					Name: core.String("精粹白金"),
					Bin:  []string{"62542688"},
				},
				LimitPay: []string{"BCZ_DEBIT"},
			},
			NoCash:       core.Bool(false),
			StockType:    core.String("NORMAL"),
			OutRequestNo: core.String("example_out_request_no"),
			ExtInfo:      core.String("{'exinfo1':'1234','exinfo2':'3456'}"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateCouponStock err:%s", err)
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
**req** | [**CreateCouponStockRequest**](CreateCouponStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**CreateCouponStockResponse**](CreateCouponStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## CreateNocashSingleitemStock

> CreateNocashSingleitemStockResponse CreateNocashSingleitemStock(CreateNocashSingleitemStockRequest)

创建无资金流单品批次



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.CreateNocashSingleitemStock(ctx,
		cashcoupons.CreateNocashSingleitemStockRequest{
			StockName:          core.String("微信支付批次"),
			AvailableMchid:     []string{"9865001"},
			AvailableBeginTime: core.String("2015-05-20T13:29:35.120+08:00"),
			AvailableEndTime:   core.String("2015-05-20T13:29:35.120+08:00"),
			NaturalPersonLimit: core.Bool(false),
			PreventApiAbuse:    core.Bool(false),
			FavorAvailableTime: &cashcoupons.FavorAvailableTime{
				FixAvailableTime: &cashcoupons.FixedAvailableTime{
					AvailableWeekDay: []int64{int64(1)},
					BeginTime:        core.Int64(0),
					EndTime:          core.Int64(3600),
				},
				SecondDayAvailable:        core.Bool(false),
				AvailableTimeAfterReceive: core.Int64(1440),
			},
			Description: core.String("微信支付代金券"),
			GoodsIdList: []string{"54312"},
			StockUseRule: &cashcoupons.StockUseRule{
				MaxCoupons:     core.Int64(100),
				MaxAmount:      core.Int64(5000),
				MaxAmountByDay: core.Int64(400),
				FixedNormalCoupon: &cashcoupons.FixedValueStockMsg{
					CouponAmount:       core.Int64(100),
					TransactionMinimum: core.Int64(100),
				},
				MaxCouponsPerUser: core.Int64(3),
				CouponType:        core.String("NORMAL"),
				GoodsTag:          []string{"123456"},
				TradeType:         []cashcoupons.TradeType{cashcoupons.TRADETYPE_MICROAPP},
				CombineUse:        core.Bool(true),
			},
			SingleitemFavor: &cashcoupons.CutTypeMsg{
				SinglePriceMax: core.Int64(100),
				CutToPrice:     core.Int64(1),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateNocashSingleitemStock err:%s", err)
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
**req** | [**CreateNocashSingleitemStockRequest**](CreateNocashSingleitemStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**CreateNocashSingleitemStockResponse**](CreateNocashSingleitemStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## CreateNocashStock

> CreateNocashStockResponse CreateNocashStock(CreateNocashStockRequest)

创建无资金流全场批次



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.CreateNocashStock(ctx,
		cashcoupons.CreateNocashStockRequest{
			StockName:          core.String("微信支付批次"),
			AvailableMerchants: []string{"9865001"},
			AvailableBeginTime: core.String("2015-05-20T13:29:35.120+08:00"),
			AvailableEndTime:   core.String("2015-05-20T13:29:35.120+08:00"),
			NaturalPersonLimit: core.Bool(false),
			PreventApiAbuse:    core.Bool(false),
			FavorAvailableTime: &cashcoupons.FavorAvailableTime{
				FixAvailableTime: &cashcoupons.FixedAvailableTime{
					AvailableWeekDay: []int64{int64(1)},
					BeginTime:        core.Int64(0),
					EndTime:          core.Int64(3600),
				},
				SecondDayAvailable:        core.Bool(false),
				AvailableTimeAfterReceive: core.Int64(1440),
			},
			Description: core.String("微信支付代金券"),
			StockUseRule: &cashcoupons.StockUseRule{
				MaxCoupons:     core.Int64(100),
				MaxAmount:      core.Int64(5000),
				MaxAmountByDay: core.Int64(400),
				FixedNormalCoupon: &cashcoupons.FixedValueStockMsg{
					CouponAmount:       core.Int64(100),
					TransactionMinimum: core.Int64(100),
				},
				MaxCouponsPerUser: core.Int64(3),
				CouponType:        core.String("NORMAL"),
				GoodsTag:          []string{"123456"},
				TradeType:         []cashcoupons.TradeType{cashcoupons.TRADETYPE_MICROAPP},
				CombineUse:        core.Bool(true),
			},
			StockType: core.String("NORMAL"),
			DiscountCutMessage: &cashcoupons.DiscountCutMsg{
				TransactionMinimum: core.Int64(10),
				RandomCutMessage: []cashcoupons.RandomCut{cashcoupons.RandomCut{
					RuleMaxAmount:  core.Int64(1),
					CutAmountBegin: core.Int64(12),
					CutAmountEnd:   core.Int64(24),
				}},
				DiscountMessage: &cashcoupons.DiscountMsg{
					DiscountAmountMax:  core.Int64(100),
					DiscountPercent:    core.Int64(88),
					TransactionMinimum: core.Int64(100),
				},
				FixCutAmount: core.Int64(12),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateNocashStock err:%s", err)
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
**req** | [**CreateNocashStockRequest**](CreateNocashStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**CreateNocashStockResponse**](CreateNocashStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ListAvailableMerchants

> AvailableMerchantCollection ListAvailableMerchants(ListAvailableMerchantsRequest)

查询代金券可用商户



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.ListAvailableMerchants(ctx,
		cashcoupons.ListAvailableMerchantsRequest{
			Offset:            core.Int64(10),
			Limit:             core.Int64(10),
			StockCreatorMchid: core.String("9865000"),
			StockId:           core.String("9865000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListAvailableMerchants err:%s", err)
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
**req** | [**ListAvailableMerchantsRequest**](ListAvailableMerchantsRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**AvailableMerchantCollection**](AvailableMerchantCollection.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ListAvailableSingleitems

> AvailableSingleitemCollection ListAvailableSingleitems(ListAvailableSingleitemsRequest)

查询可核销商品编码



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.ListAvailableSingleitems(ctx,
		cashcoupons.ListAvailableSingleitemsRequest{
			Offset:            core.Int64(10),
			Limit:             core.Int64(10),
			StockCreatorMchid: core.String("9865000"),
			StockId:           core.String("9865000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListAvailableSingleitems err:%s", err)
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
**req** | [**ListAvailableSingleitemsRequest**](ListAvailableSingleitemsRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**AvailableSingleitemCollection**](AvailableSingleitemCollection.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ListStocks

> StockCollection ListStocks(ListStocksRequest)

条件查询批次列表



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.ListStocks(ctx,
		cashcoupons.ListStocksRequest{
			Offset:            core.Int64(1),
			Limit:             core.Int64(8),
			StockCreatorMchid: core.String("9856888"),
			CreateStartTime:   core.String("2015-05-20T13:29:35.120+08:00"),
			CreateEndTime:     core.String("2015-05-20T13:29:35.120+08:00"),
			Status:            core.String("paused"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListStocks err:%s", err)
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
**req** | [**ListStocksRequest**](ListStocksRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**StockCollection**](StockCollection.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ModifyAvailableMerchants

> ModifyAvailableMerchantResponse ModifyAvailableMerchants(ModifyAvailableMerchantsRequest)

修改批次可用商户号



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.ModifyAvailableMerchants(ctx,
		cashcoupons.ModifyAvailableMerchantsRequest{
			StockId:           core.String("9865000"),
			StockCreatorMchid: core.String("9865000"),
			AddMchidList:      []string{"9865002"},
			DeleteMchidList:   []string{"9865004"},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ModifyAvailableMerchants err:%s", err)
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
**req** | [**ModifyAvailableMerchantsRequest**](ModifyAvailableMerchantsRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ModifyAvailableMerchantResponse**](ModifyAvailableMerchantResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ModifyAvailableSingleitems

> ModifyAvailableSingleitemResponse ModifyAvailableSingleitems(ModifyAvailableSingleitemsRequest)

修改可核销商品编码



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.ModifyAvailableSingleitems(ctx,
		cashcoupons.ModifyAvailableSingleitemsRequest{
			StockId:           core.String("9865000"),
			StockCreatorMchid: core.String("9865000"),
			AddGoodsIdList:    []string{"34565"},
			DeleteGoodsIdList: []string{"123545"},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ModifyAvailableSingleitems err:%s", err)
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
**req** | [**ModifyAvailableSingleitemsRequest**](ModifyAvailableSingleitemsRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ModifyAvailableSingleitemResponse**](ModifyAvailableSingleitemResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ModifyBudget

> ModifyStockBudgetResponse ModifyBudget(ModifyBudgetRequest)

修改批次预算金额



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.ModifyBudget(ctx,
		cashcoupons.ModifyBudgetRequest{
			StockId:           core.String("9856000"),
			StockCreatorMchid: core.String("9856000"),
			TargetMaxAmount:   core.Int64(1),
			CurrentMaxAmount:  core.Int64(1),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ModifyBudget err:%s", err)
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
**req** | [**ModifyBudgetRequest**](ModifyBudgetRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ModifyStockBudgetResponse**](ModifyStockBudgetResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ModifyStock

> ModifyStockResponse ModifyStock(ModifyStockRequest)

修改批次基本信息



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.ModifyStock(ctx,
		cashcoupons.ModifyStockRequest{
			StockId:           core.String("8965000"),
			StockCreatorMchid: core.String("9865000"),
			MaxAmountByDay:    core.Int64(1000),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ModifyStock err:%s", err)
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
**req** | [**ModifyStockRequest**](ModifyStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ModifyStockResponse**](ModifyStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## PauseStock

> PauseStockResponse PauseStock(PauseStockRequest)

暂停批次



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.PauseStock(ctx,
		cashcoupons.PauseStockRequest{
			StockId:           core.String("实例值"),
			StockCreatorMchid: core.String("8965000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call PauseStock err:%s", err)
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
**req** | [**PauseStockRequest**](PauseStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**PauseStockResponse**](PauseStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## QueryStock

> Stock QueryStock(QueryStockRequest)

查询批次详情



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.QueryStock(ctx,
		cashcoupons.QueryStockRequest{
			StockId:           core.String("9856888"),
			StockCreatorMchid: core.String("123456"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryStock err:%s", err)
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
**req** | [**QueryStockRequest**](QueryStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**Stock**](Stock.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## RefundFlow

> StockDownloadRefundFlowResponse RefundFlow(RefundFlowRequest)

下载批次退款明细



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.RefundFlow(ctx,
		cashcoupons.RefundFlowRequest{
			StockId: core.String("9865000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call RefundFlow err:%s", err)
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
**req** | [**RefundFlowRequest**](RefundFlowRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**StockDownloadRefundFlowResponse**](StockDownloadRefundFlowResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## RestartStock

> RestartStockResponse RestartStock(RestartStockRequest)

重启批次



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.RestartStock(ctx,
		cashcoupons.RestartStockRequest{
			StockId:           core.String("8965000"),
			StockCreatorMchid: core.String("9865000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call RestartStock err:%s", err)
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
**req** | [**RestartStockRequest**](RestartStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**RestartStockResponse**](RestartStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## StartStock

> StartStockResponse StartStock(StartStockRequest)

激活开启批次



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.StartStock(ctx,
		cashcoupons.StartStockRequest{
			StockId:           core.String("9856000"),
			StockCreatorMchid: core.String("8956000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call StartStock err:%s", err)
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
**req** | [**StartStockRequest**](StartStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**StartStockResponse**](StartStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## StopStock

> StopStockResponse StopStock(StopStockRequest)

终止批次



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.StopStock(ctx,
		cashcoupons.StopStockRequest{
			StockId:           core.String("8695000"),
			StockCreatorMchid: core.String("9865000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call StopStock err:%s", err)
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
**req** | [**StopStockRequest**](StopStockRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**StopStockResponse**](StopStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## UseFlow

> StockDownloadUseFlowResponse UseFlow(UseFlowRequest)

下载批次核销明细



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

	svc := cashcoupons.StockApiService{Client: client}
	resp, result, err := svc.UseFlow(ctx,
		cashcoupons.UseFlowRequest{
			StockId: core.String("9865000"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call UseFlow err:%s", err)
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
**req** | [**UseFlowRequest**](UseFlowRequest.md) | API `cashcoupons` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**StockDownloadUseFlowResponse**](StockDownloadUseFlowResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#cashcouponsstockapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

