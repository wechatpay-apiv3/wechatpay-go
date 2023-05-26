# merchantexclusivecoupon/BusiFavorApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**CouponCodeInfo**](#couponcodeinfo) | **Get** /v3/marketing/busifavor/stocks/{stock_id}/couponcodes/{coupon_code} | 查询预存code详情
[**CreateBusifavorStock**](#createbusifavorstock) | **Post** /v3/marketing/busifavor/stocks | 创建商家券
[**DeleteCouponCode**](#deletecouponcode) | **Delete** /v3/marketing/busifavor/stocks/{stock_id}/couponcodes/{coupon_code} | 删除预存code
[**ModifyBudget**](#modifybudget) | **Patch** /v3/marketing/busifavor/stocks/{stock_id}/budget | 修改批次预算
[**ModifyStockInfo**](#modifystockinfo) | **Patch** /v3/marketing/busifavor/stocks/{stock_id} | 修改商家券基本信息
[**QueryCouponCodeList**](#querycouponcodelist) | **Get** /v3/marketing/busifavor/stocks/{stock_id}/couponcodes | 查询预存code列表
[**QueryStock**](#querystock) | **Get** /v3/marketing/busifavor/stocks/{stock_id} | 查询商家券批次详情
[**UploadCouponCode**](#uploadcouponcode) | **Post** /v3/marketing/busifavor/stocks/{stock_id}/couponcodes | 上传预存code



## CouponCodeInfo

> CouponCodeInfoResponse CouponCodeInfo(CouponCodeInfoRequest)

查询预存code详情



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	resp, result, err := svc.CouponCodeInfo(ctx,
		merchantexclusivecoupon.CouponCodeInfoRequest{
			StockId:    core.String("98065001"),
			CouponCode: core.String("ABC1234567"),
			Appid:      core.String("wx1234567889999"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CouponCodeInfo err:%s", err)
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
**req** | [**CouponCodeInfoRequest**](CouponCodeInfoRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**CouponCodeInfoResponse**](CouponCodeInfoResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## CreateBusifavorStock

> CreateBusiFavorStockResponse CreateBusifavorStock(CreateBusifavorStockRequest)

创建商家券



### 调用示例

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	resp, result, err := svc.CreateBusifavorStock(ctx,
		merchantexclusivecoupon.CreateBusifavorStockRequest{
			StockName:      core.String("8月1日活动券"),
			BelongMerchant: core.String("10000098"),
			Comment:        core.String("活动使用"),
			GoodsName:      core.String("填写商家券可适用的商品或服务"),
			StockType:      merchantexclusivecoupon.BUSIFAVORSTOCKTYPE_NORMAL.Ptr(),
			CouponUseRule: &merchantexclusivecoupon.CouponUseRule{
				CouponAvailableTime: &merchantexclusivecoupon.FavorAvailableTime{
					AvailableBeginTime:       core.Time(time.Now()),
					AvailableEndTime:         core.Time(time.Now()),
					AvailableDayAfterReceive: core.Int64(3),
					AvailableWeek: &merchantexclusivecoupon.AvailableWeek{
						WeekDay: []int64{int64(1)},
						AvailableDayTime: []merchantexclusivecoupon.AvailableCurrentDayTime{merchantexclusivecoupon.AvailableCurrentDayTime{
							BeginTime: core.Int64(3600),
							EndTime:   core.Int64(86399),
						}},
					},
					IrregularyAvaliableTime: []merchantexclusivecoupon.IrregularAvailableTime{merchantexclusivecoupon.IrregularAvailableTime{
						BeginTime: core.Time(time.Now()),
						EndTime:   core.Time(time.Now()),
					}},
					WaitDaysAfterReceive: core.Int64(7),
				},
				FixedNormalCoupon: &merchantexclusivecoupon.FixedValueStockMsg{
					DiscountAmount:     core.Int64(5),
					TransactionMinimum: core.Int64(100),
				},
				DiscountCoupon: &merchantexclusivecoupon.DiscountMsg{
					DiscountPercent:    core.Int64(88),
					TransactionMinimum: core.Int64(100),
				},
				ExchangeCoupon: &merchantexclusivecoupon.ExchangeMsg{
					ExchangePrice:      core.Int64(100),
					TransactionMinimum: core.Int64(100),
				},
				UseMethod:         merchantexclusivecoupon.COUPONUSEMETHOD_OFF_LINE.Ptr(),
				MiniProgramsAppid: core.String("wx23232232323"),
				MiniProgramsPath:  core.String("/path/index/index"),
			},
			StockSendRule: &merchantexclusivecoupon.StockSendRule{
				MaxAmount:          core.Int64(100000),
				MaxCoupons:         core.Int64(100),
				MaxCouponsPerUser:  core.Int64(5),
				MaxAmountByDay:     core.Int64(1000),
				MaxCouponsByDay:    core.Int64(100),
				NaturalPersonLimit: core.Bool(false),
				PreventApiAbuse:    core.Bool(false),
				Transferable:       core.Bool(false),
				Shareable:          core.Bool(false),
			},
			OutRequestNo: core.String("100002322019090134234sfdf"),
			CustomEntrance: &merchantexclusivecoupon.CustomEntrance{
				MiniProgramsInfo: &merchantexclusivecoupon.MiniAppInfo{
					MiniProgramsAppid: core.String("wx234545656765876"),
					MiniProgramsPath:  core.String("/path/index/index"),
					EntranceWords:     core.String("欢迎选购"),
					GuidingWords:      core.String("获取更多优惠"),
				},
				Appid:           core.String("wx324345hgfhfghfg"),
				HallId:          core.String("233455656"),
				StoreId:         core.String("233554655"),
				CodeDisplayMode: merchantexclusivecoupon.CODEDISPLAYMODE_NOT_SHOW.Ptr(),
			},
			DisplayPatternInfo: &merchantexclusivecoupon.DisplayPatternInfo{
				Description:     core.String("xxx门店可用"),
				MerchantLogoUrl: core.String("https://xxx"),
				MerchantName:    core.String("微信支付"),
				BackgroundColor: core.String("xxxxx"),
				CouponImageUrl:  core.String("图片cdn地址"),
				FinderInfo: &merchantexclusivecoupon.FinderInfo{
					FinderId:                 core.String("sph6Rngt2T4RlUf"),
					FinderVideoId:            core.String("export/UzFfAgtgekIEAQAAAAAAb4MgnPInmAAAAAstQy6ubaLX4KHWvLEZgBPEwIEgVnk9HIP-zNPgMJofG6tpdGPJNg_ojtEjoT94"),
					FinderVideoCoverImageUrl: core.String("https://wxpaylogo.qpic.cn/xxx"),
				},
			},
			CouponCodeMode: merchantexclusivecoupon.COUPONCODEMODE_WECHATPAY_MODE.Ptr(),
			NotifyConfig: &merchantexclusivecoupon.NotifyConfig{
				NotifyAppid: core.String("wx23232232323"),
			},
			Subsidy: core.Bool(false),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateBusifavorStock err:%s", err)
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
**req** | [**CreateBusifavorStockRequest**](CreateBusifavorStockRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**CreateBusiFavorStockResponse**](CreateBusiFavorStockResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## DeleteCouponCode

> DeleteCouponCodeResponse DeleteCouponCode(DeleteCouponCodeRequest)

删除预存code



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	resp, result, err := svc.DeleteCouponCode(ctx,
		merchantexclusivecoupon.DeleteCouponCodeRequest{
			StockId:         core.String("98065001"),
			CouponCode:      core.String("ABC9588201"),
			DeleteRequestNo: core.String("100002322019090134234sfdf"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call DeleteCouponCode err:%s", err)
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
**req** | [**DeleteCouponCodeRequest**](DeleteCouponCodeRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**DeleteCouponCodeResponse**](DeleteCouponCodeResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ModifyBudget

> ModifyBudgetResponse ModifyBudget(ModifyBudgetRequest)

修改批次预算



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	resp, result, err := svc.ModifyBudget(ctx,
		merchantexclusivecoupon.ModifyBudgetRequest{
			StockId:                core.String("98065001"),
			TargetMaxCoupons:       core.Int64(3000),
			CurrentMaxCoupons:      core.Int64(500),
			TargetMaxCouponsByDay:  core.Int64(500),
			CurrentMaxCouponsByDay: core.Int64(300),
			ModifyBudgetRequestNo:  core.String("1002600620019090123143254436"),
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
**req** | [**ModifyBudgetRequest**](ModifyBudgetRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ModifyBudgetResponse**](ModifyBudgetResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ModifyStockInfo

> void ModifyStockInfo(ModifyStockInfoRequest)

修改商家券基本信息



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	result, err := svc.ModifyStockInfo(ctx,
		merchantexclusivecoupon.ModifyStockInfoRequest{
			StockId: core.String("101156451224"),
			CustomEntrance: &merchantexclusivecoupon.ModifyCustomEntrance{
				MiniProgramsInfo: &merchantexclusivecoupon.ModifyMiniAppInfo{
					MiniProgramsAppid: core.String("wx234545656765876"),
					MiniProgramsPath:  core.String("/path/index/index"),
					EntranceWords:     core.String("欢迎选购"),
					GuidingWords:      core.String("获取更多优惠"),
				},
				Appid:           core.String("wx324345hgfhfghfg"),
				HallId:          core.String("233455656"),
				CodeDisplayMode: merchantexclusivecoupon.CODEDISPLAYMODE_NOT_SHOW.Ptr(),
			},
			StockName:    core.String("8月1日活动券"),
			Comment:      core.String("活动使用"),
			GoodsName:    core.String("xxx商品使用"),
			OutRequestNo: core.String("6122352020010133287985742"),
			DisplayPatternInfo: &merchantexclusivecoupon.DisplayPatternInfo{
				Description:     core.String("xxx门店可用"),
				MerchantLogoUrl: core.String("https://xxx"),
				MerchantName:    core.String("微信支付"),
				BackgroundColor: core.String("xxxxx"),
				CouponImageUrl:  core.String("图片cdn地址"),
				FinderInfo: &merchantexclusivecoupon.FinderInfo{
					FinderId:                 core.String("sph6Rngt2T4RlUf"),
					FinderVideoId:            core.String("export/UzFfAgtgekIEAQAAAAAAb4MgnPInmAAAAAstQy6ubaLX4KHWvLEZgBPEwIEgVnk9HIP-zNPgMJofG6tpdGPJNg_ojtEjoT94"),
					FinderVideoCoverImageUrl: core.String("https://wxpaylogo.qpic.cn/xxx"),
				},
			},
			CouponUseRule: &merchantexclusivecoupon.ModifyCouponUseRule{
				UseMethod:         merchantexclusivecoupon.COUPONUSEMETHOD_OFF_LINE.Ptr(),
				MiniProgramsAppid: core.String("wx23232232323"),
				MiniProgramsPath:  core.String("/path/index/index"),
			},
			StockSendRule: &merchantexclusivecoupon.ModifyStockSendRule{
				NaturalPersonLimit: core.Bool(false),
				PreventApiAbuse:    core.Bool(false),
			},
			NotifyConfig: &merchantexclusivecoupon.NotifyConfig{
				NotifyAppid: core.String("wx23232232323"),
			},
			Subsidy: core.Bool(true),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ModifyStockInfo err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d", result.Response.StatusCode)
	}
}
```

### 参数列表
参数名 | 参数类型 | 参数描述
------------- | ------------- | -------------
**ctx** | **context.Context** | Golang 上下文，可用于日志、请求取消、请求跟踪等功能|
**req** | [**ModifyStockInfoRequest**](ModifyStockInfoRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## QueryCouponCodeList

> CouponCodeListResponse QueryCouponCodeList(QueryCouponCodeListRequest)

查询预存code列表



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	resp, result, err := svc.QueryCouponCodeList(ctx,
		merchantexclusivecoupon.QueryCouponCodeListRequest{
			StockId: core.String("98065001"),
			Limit:   core.Int64(100),
			Offset:  core.Int64(10),
			Appid:   core.String("wx1234567889999"),
			Status:  merchantexclusivecoupon.COUPONCODESTATUS_AVAILABLE.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryCouponCodeList err:%s", err)
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
**req** | [**QueryCouponCodeListRequest**](QueryCouponCodeListRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**CouponCodeListResponse**](CouponCodeListResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## QueryStock

> StockGetResponse QueryStock(QueryStockRequest)

查询商家券批次详情



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	resp, result, err := svc.QueryStock(ctx,
		merchantexclusivecoupon.QueryStockRequest{
			StockId: core.String("100088"),
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
**req** | [**QueryStockRequest**](QueryStockRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**StockGetResponse**](StockGetResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## UploadCouponCode

> UploadCouponCodeResponse UploadCouponCode(UploadCouponCodeRequest)

上传预存code



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/merchantexclusivecoupon"
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

	svc := merchantexclusivecoupon.BusiFavorApiService{Client: client}
	resp, result, err := svc.UploadCouponCode(ctx,
		merchantexclusivecoupon.UploadCouponCodeRequest{
			StockId:         core.String("98065001"),
			CouponCodeList:  []string{"ABC9588200"},
			UploadRequestNo: core.String("100002322019090134234sfdf"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call UploadCouponCode err:%s", err)
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
**req** | [**UploadCouponCodeRequest**](UploadCouponCodeRequest.md) | API `merchantexclusivecoupon` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**UploadCouponCodeResponse**](UploadCouponCodeResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#merchantexclusivecouponbusifavorapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

