# payrollcard/AuthenticationsApi

所有URI均基于微信支付 API 地址： *https://api.mch.weixin.qq.com*

方法名 | HTTP 请求 | 描述
------------- | ------------- | -------------
[**GetAuthentication**](#getauthentication) | **Get** /v3/payroll-card/authentications/{authenticate_number} | 获取核身结果
[**ListAuthentications**](#listauthentications) | **Get** /v3/payroll-card/authentications | 查询核身记录
[**PreOrderAuthentication**](#preorderauthentication) | **Post** /v3/payroll-card/authentications/pre-order | 微工卡核身预下单
[**PreOrderAuthenticationWithAuth**](#preorderauthenticationwithauth) | **Post** /v3/payroll-card/authentications/pre-order-with-auth | 微工卡核身预下单（流程中完成授权）



## GetAuthentication

> AuthenticationEntity GetAuthentication(GetAuthenticationRequest)

获取核身结果



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payrollcard"
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.GetAuthentication(ctx,
		payrollcard.GetAuthenticationRequest{
			SubMchid:           core.String("1111111"),
			AuthenticateNumber: core.String("mcdhehfgisdhfjghed39384564i83"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call GetAuthentication err:%s", err)
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
**req** | [**GetAuthenticationRequest**](GetAuthenticationRequest.md) | API `payrollcard` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**AuthenticationEntity**](AuthenticationEntity.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#payrollcardauthenticationsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## ListAuthentications

> ListAuthenticationsResponse ListAuthentications(ListAuthenticationsRequest)

查询核身记录



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payrollcard"
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.ListAuthentications(ctx,
		payrollcard.ListAuthenticationsRequest{
			Openid:            core.String("onqOjjmo8wmTOOtSKwXtGjg9Gb58"),
			Appid:             core.String("wxa1111111"),
			SubAppid:          core.String("wxa1111111"),
			SubMchid:          core.String("1111111"),
			AuthenticateDate:  core.String("2020-12-25"),
			AuthenticateState: core.String("AUTHENTICATE_SUCCESS"),
			Offset:            core.Int64(0),
			Limit:             core.Int64(10),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call ListAuthentications err:%s", err)
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
**req** | [**ListAuthenticationsRequest**](ListAuthenticationsRequest.md) | API `payrollcard` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**ListAuthenticationsResponse**](ListAuthenticationsResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#payrollcardauthenticationsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## PreOrderAuthentication

> PreOrderAuthenticationResponse PreOrderAuthentication(PreOrderAuthenticationRequest)

微工卡核身预下单



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payrollcard"
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.PreOrderAuthentication(ctx,
		payrollcard.PreOrderAuthenticationRequest{
			Openid:             core.String("onqOjjmo8wmTOOtSKwXtGjg9Gb58"),
			Appid:              core.String("wxa1111111"),
			SubMchid:           core.String("1111111"),
			SubAppid:           core.String("wxa1111111"),
			AuthenticateNumber: core.String("mcdhehfgisdhfjghed39384564i83"),
			ProjectName:        core.String("某项目"),
			EmployerName:       core.String("某用工企业"),
			AuthenticateType:   payrollcard.AUTHENTICATIONTYPE_NORMAL.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call PreOrderAuthentication err:%s", err)
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
**req** | [**PreOrderAuthenticationRequest**](PreOrderAuthenticationRequest.md) | API `payrollcard` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**PreOrderAuthenticationResponse**](PreOrderAuthenticationResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#payrollcardauthenticationsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)


## PreOrderAuthenticationWithAuth

> PreOrderAuthenticationWithAuthResponse PreOrderAuthenticationWithAuth(PreOrderAuthenticationWithAuthRequest)

微工卡核身预下单（流程中完成授权）



### 调用示例

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payrollcard"
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

	svc := payrollcard.AuthenticationsApiService{Client: client}
	resp, result, err := svc.PreOrderAuthenticationWithAuth(ctx,
		payrollcard.PreOrderAuthenticationWithAuthRequest{
			Openid:             core.String("onqOjjmo8wmTOOtSKwXtGjg9Gb58"),
			Appid:              core.String("wxa1111111"),
			SubMchid:           core.String("1111111"),
			SubAppid:           core.String("wxa1111111"),
			AuthenticateNumber: core.String("mcdhehfgisdhfjghed39384564i83"),
			ProjectName:        core.String("某项目"),
			EmployerName:       core.String("某用工企业"),
			UserName:           core.String("LP7bT4hQXUsOZCEvK2YrSiqFsnP0oRMfeoLN0vBg"),
			IdCardNumber:       core.String("7FzH5XksJG3a8HLLsaaUV6K54y1OnPMY5"),
			EmploymentType:     payrollcard.EMPLOYMENTTYPE_LONG_TERM_EMPLOYMENT.Ptr(),
			AuthenticateType:   payrollcard.AUTHENTICATIONTYPE_NORMAL.Ptr(),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call PreOrderAuthenticationWithAuth err:%s", err)
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
**req** | [**PreOrderAuthenticationWithAuthRequest**](PreOrderAuthenticationWithAuthRequest.md) | API `payrollcard` 所定义的本接口需要的所有参数，包括`Path`/`Query`/`Body` 3类参数|

### 返回结果
Name | Type | Description
------------- | ------------- | -------------
**resp** | \*[**PreOrderAuthenticationWithAuthResponse**](PreOrderAuthenticationWithAuthResponse.md) | 结构化的接口返回结果
**result** | **\*core.APIResult** | 本次 API 访问的请求与应答信息
**err** | **error** | 本次 API 访问中发生的错误，当且仅当 API 失败时存在

[\[返回顶部\]](#payrollcardauthenticationsapi)
[\[返回接口列表\]](README.md#接口列表)
[\[返回类型列表\]](README.md#类型列表)
[\[返回服务README\]](README.md)

