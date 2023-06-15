# giftactivity/ActivityApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**AddActivityMerchant**](#addactivitymerchant) | **Post** /v3/marketing/paygiftactivity/activities/{activity_id}/merchants/add | 新增活动发券商户号
[**CreateFullSendAct**](#createfullsendact) | **Post** /v3/marketing/paygiftactivity/unique-threshold-activity | 创建全场满额送活动
[**DeleteActivityMerchant**](#deleteactivitymerchant) | **Post** /v3/marketing/paygiftactivity/activities/{activity_id}/merchants/delete | 删除活动发券商户号
[**GetActDetail**](#getactdetail) | **Get** /v3/marketing/paygiftactivity/activities/{activity_id} | 获取活动详情接口
[**ListActivities**](#listactivities) | **Get** /v3/marketing/paygiftactivity/activities | 获取支付有礼活动列表
[**ListActivityMerchant**](#listactivitymerchant) | **Get** /v3/marketing/paygiftactivity/activities/{activity_id}/merchants | 获取活动发券商户号
[**ListActivitySku**](#listactivitysku) | **Get** /v3/marketing/paygiftactivity/activities/{activity_id}/goods | 获取活动指定商品列表
[**TerminateActivity**](#terminateactivity) | **Post** /v3/marketing/paygiftactivity/activities/{activity_id}/terminate | 终止活动



## AddActivityMerchant

> AddActivityMerchantResponse AddActivityMerchant(AddActivityMerchantRequest)

新增活动发券商户号



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.AddActivityMerchant(ctx,
		giftactivity.AddActivityMerchantRequest{
			ActivityId:     core.String("126002309"),
			MerchantIdList: []string{"100123456"},
			AddRequestNo:   core.String("100002322019090134234sfdf"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call AddActivityMerchant err:%s", err)
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
**req** | [**AddActivityMerchantRequest**](AddActivityMerchantRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**AddActivityMerchantResponse**](AddActivityMerchantResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## CreateFullSendAct

> CreateFullSendActResponse CreateFullSendAct(CreateFullSendActRequest)

创建全场满额送活动



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.CreateFullSendAct(ctx,
		giftactivity.CreateFullSendActRequest{
			ActivityBaseInfo: &giftactivity.ActBaseInfo{
				ActivityName:        core.String("良品铺子回馈活动"),
				ActivitySecondTitle: core.String("海飞丝的券"),
				MerchantLogoUrl:     core.String("https://tool.oschina.net/regex.jpg"),
				BackgroundColor:     core.String("Color010"),
				BeginTime:           core.String("2015-05-20T13:29:35.120+08:00"),
				EndTime:             core.String("2015-05-20T13:29:35.120+08:00"),
				AvailablePeriods: &giftactivity.AvailablePeriod{
					AvailableTime: []giftactivity.AvailableTime{giftactivity.AvailableTime{
						BeginTime: core.String("2015-05-20T00:00:00.000+08:00"),
						EndTime:   core.String("2015-05-20T23:59:59.000+08:00"),
					}},
					AvailableDayTime: []giftactivity.AvailableDayTime{giftactivity.AvailableDayTime{
						BeginDayTime: core.String("110000"),
						EndDayTime:   core.String("135959"),
					}},
				},
				OutRequestNo:      core.String("100002322019090134234sfdf"),
				DeliveryPurpose:   giftactivity.DELIVERYPURPOSECATEGORY_OFF_LINE_PAY.Ptr(),
				MiniProgramsAppid: core.String("wx23232232323"),
				MiniProgramsPath:  core.String("/path/index/index"),
			},
			AwardSendRule: &giftactivity.FullSendRule{
				TransactionAmountMinimum: core.Int64(100),
				SendContent:              giftactivity.SENDCONTENTCATEGORY_SINGLE_COUPON.Ptr(),
				AwardType:                giftactivity.AWARDTYPE_BUSIFAVOR.Ptr(),
				AwardList: []giftactivity.AwardBaseInfo{giftactivity.AwardBaseInfo{
					StockId:          core.String("98065001"),
					OriginalImageUrl: core.String("https://tool.oschina.net/regex.jpg"),
					ThumbnailUrl:     core.String("https://tool.oschina.net/regex.jpg"),
				}},
				MerchantOption: giftactivity.SENDMERCHANTOPTION_IN_SEVICE_COUPON_MERCHANT.Ptr(),
				MerchantIdList: []string{"10000022"},
			},
			AdvancedSetting: &giftactivity.ActAdvancedSetting{
				DeliveryUserCategory: giftactivity.DELIVERYUSERCATEGORY_DELIVERY_ALL_PERSON.Ptr(),
				MerchantMemberAppid:  core.String("34567890"),
				PaymentMode: &giftactivity.PaymentMode{
					PaymentSceneList: []giftactivity.PaymentScene{giftactivity.PAYMENTSCENE_APP_SCENE},
				},
				PaymentMethodInformation: &giftactivity.PaymentMethodInfo{
					PaymentMethod:    giftactivity.PAYMENTMETHODCATEGORY_CFT.Ptr(),
					BankAbbreviation: core.String("AHRCUB_CREDIT"),
				},
				GoodsTags: []string{"xxx"},
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateFullSendAct err:%s", err)
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
**req** | [**CreateFullSendActRequest**](CreateFullSendActRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**CreateFullSendActResponse**](CreateFullSendActResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## DeleteActivityMerchant

> DeleteActivityMerchantResponse DeleteActivityMerchant(DeleteActivityMerchantRequest)

删除活动发券商户号



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.DeleteActivityMerchant(ctx,
		giftactivity.DeleteActivityMerchantRequest{
			ActivityId:      core.String("126002309"),
			MerchantIdList:  []string{"100123456"},
			DeleteRequestNo: core.String("100002322019090134234sfdf"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call DeleteActivityMerchant err:%s", err)
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
**req** | [**DeleteActivityMerchantRequest**](DeleteActivityMerchantRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**DeleteActivityMerchantResponse**](DeleteActivityMerchantResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## GetActDetail

> GetActDetailResponse GetActDetail(GetActDetailRequest)

获取活动详情接口



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.GetActDetail(ctx,
		giftactivity.GetActDetailRequest{
			ActivityId: core.String("10028001"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetActDetail err:%s", err)
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
**req** | [**GetActDetailRequest**](GetActDetailRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**GetActDetailResponse**](GetActDetailResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ListActivities

> ListActivitiesResponse ListActivities(ListActivitiesRequest)

获取支付有礼活动列表



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.ListActivities(ctx,
		giftactivity.ListActivitiesRequest{
			Offset:         core.Int64(1),
			Limit:          core.Int64(20),
			ActivityName:   core.String("良品铺子回馈活动"),
			ActivityStatus: giftactivity.ACTSTATUS_ACT_STATUS_UNKNOWN.Ptr(),
			AwardType:      giftactivity.AWARDTYPE_BUSIFAVOR.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListActivities err:%s", err)
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
**req** | [**ListActivitiesRequest**](ListActivitiesRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ListActivitiesResponse**](ListActivitiesResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ListActivityMerchant

> ListActMchResponse ListActivityMerchant(ListActivityMerchantRequest)

获取活动发券商户号



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.ListActivityMerchant(ctx,
		giftactivity.ListActivityMerchantRequest{
			ActivityId: core.String("126002309"),
			Offset:     core.Int64(1),
			Limit:      core.Int64(20),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListActivityMerchant err:%s", err)
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
**req** | [**ListActivityMerchantRequest**](ListActivityMerchantRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ListActMchResponse**](ListActMchResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ListActivitySku

> ListActSkuResponse ListActivitySku(ListActivitySkuRequest)

获取活动指定商品列表



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.ListActivitySku(ctx,
		giftactivity.ListActivitySkuRequest{
			ActivityId: core.String("10028001"),
			Offset:     core.Int64(1),
			Limit:      core.Int64(20),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListActivitySku err:%s", err)
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
**req** | [**ListActivitySkuRequest**](ListActivitySkuRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ListActSkuResponse**](ListActSkuResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## TerminateActivity

> TerminateActResponse TerminateActivity(TerminateActivityRequest)

终止活动



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/giftactivity"
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

	svc := giftactivity.ActivityApiService{Client: client}
	resp, result, err := svc.TerminateActivity(ctx,
		giftactivity.TerminateActivityRequest{
			ActivityId: core.String("10028001"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call TerminateActivity err:%s", err)
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
**req** | [**TerminateActivityRequest**](TerminateActivityRequest.md) | API `giftactivity` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**TerminateActResponse**](TerminateActResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#giftactivityactivityapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

