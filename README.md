# 微信支付 API v3 Go SDK
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/wechatpay-apiv3/wechatpay-go)
[![huntr](https://cdn.huntr.dev/huntr_security_badge_mono.svg)](https://huntr.dev)
[![licence](https://badgen.net/github/license/wechatpay-apiv3/wechatpay-go)](https://github.com/wechatpay-apiv3/wechatpay-go/blob/main/LICENSE)

[微信支付 APIv3](https://wechatpay-api.gitbook.io/wechatpay-api-v3/) 官方Go语言客户端代码库。

## 功能介绍

1. 接口 SDK。详见 [接口介绍](services)。
1. HTTP 客户端 `core.Client`，支持请求签名和应答验签。如果 SDK 未支持你需要的接口，请用此客户端发起请求。
3. 回调通知处理库 `core/notify`，支持微信支付回调通知的验签和解密。详见 [回调通知验签与解密](#回调通知的验签与解密)。
4. 证书下载、[敏感信息加解密](#敏感信息加解密) 等辅助能力。

### 兼容性

当前版本为测试版本，微信支付会尽量保持向后兼容。但可能因为可用性或易用性，同历史版本存在不兼容。如果你使用版本 `<= v0.2.2`，升级前请参考 [升级指南](UPGRADING.md)。

## 快速开始

### 安装

#### 1、使用 Go Modules 管理你的项目

如果你的项目还不是使用 Go Modules 做依赖管理，在项目根目录下执行：

```shell
go mod init
```

#### 2、无需 clone 仓库中的代码，直接在项目目录中执行： 
```shell
go get -u github.com/wechatpay-apiv3/wechatpay-go
```
来添加依赖，完成 `go.mod` 修改与 SDK 下载。

### 发送请求

先初始化一个 `core.Client` 实例，再向微信支付发送请求。

```go
package main

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func main() {
	var (
		mchID                      string = "190000****"                                // 商户号
		mchCertificateSerialNumber string = "3775B6A45ACD588826D15E583A95F5DD********"  // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"          // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}
	
	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
}
```

`resp` 是反序列化（UnmarshalJSON）后的应答。上例中是 `services/certificates` 包中的 `*certificates.Certificate`。

`result` 是 `*core.APIResult` 实例，包含了完整的请求报文 `*http.Request` 和应答报文 `*http.Response`。

#### 名词解释

+ **商户 API 证书**，是用来证实商户身份的。证书中包含商户号、证书序列号、证书有效期等信息，由证书授权机构（Certificate Authority ，简称 CA）签发，以防证书被伪造或篡改。如何获取请见 [商户 API 证书](https://wechatpay-api.gitbook.io/wechatpay-api-v3/ren-zheng/zheng-shu#shang-hu-api-zheng-shu) 。
+ **商户 API 私钥**。商户申请商户 API 证书时，会生成商户私钥，并保存在本地证书文件夹的文件 apiclient_key.pem 中。
> :warning: 不要把私钥文件暴露在公共场合，如上传到 Github，写在客户端代码等。
+ **微信支付平台证书**。微信支付平台证书是指由微信支付负责申请的，包含微信支付平台标识、公钥信息的证书。商户使用微信支付平台证书中的公钥验证应答签名。获取微信支付平台证书需通过 [获取平台证书列表](https://wechatpay-api.gitbook.io/wechatpay-api-v3/ren-zheng/zheng-shu#ping-tai-zheng-shu) 接口下载。
+ **证书序列号**。每个证书都有一个由 CA 颁发的唯一编号，即证书序列号。扩展阅读 [如何查看证书序列号](https://wechatpay-api.gitbook.io/wechatpay-api-v3/chang-jian-wen-ti/zheng-shu-xiang-guan#ru-he-cha-kan-zheng-shu-xu-lie-hao) 。
+ **微信支付 APIv3 密钥**，是在回调通知和微信支付平台证书下载接口中，为加强数据安全，对关键信息 `AES-256-GCM` 加密时使用的对称加密密钥。

## 更多示例

### 以 [JSAPI下单](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_1.shtml) 为例

```go
import (
	"log"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

svc := jsapi.JsapiApiService{Client: client}
// 得到prepay_id，以及调起支付所需的参数和签名
resp, result, err := svc.PrepayWithRequestPayment(ctx,
	jsapi.PrepayRequest{
		Appid:       core.String("wxd678efh567hg6787"),
		Mchid:       core.String("1900009191"),
		Description: core.String("Image形象店-深圳腾大-QQ公仔"),
		OutTradeNo:  core.String("1217752501201407033233368018"),
		Attach:      core.String("自定义数据说明"),
		NotifyUrl:   core.String("https://www.weixin.qq.com/wxpay/pay.php"),
		Amount: &jsapi.Amount{
			Total: core.Int64(100),
		},
		Payer: &jsapi.Payer{
			Openid: core.String("oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"),
		},
	},
)

if err == nil {
	log.Println(resp)
} else {
	log.Println(err)
}
```

### 以 [查询订单](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_2.shtml) 为例

```go
import (
	"log"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

svc := jsapi.JsapiApiService{Client: client}

resp, result, err := svc.QueryOrderById(ctx,
	jsapi.QueryOrderByIdRequest{
		TransactionId: core.String("4200000985202103031441826014"),
		Mchid:         core.String("1900009191"),
	},
)

if err == nil {
	log.Println(resp)
} else {
	log.Println(err)
}

```

### 以 [图片上传API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_1.shtml) 为例：
```go
import (
	"os"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services/fileuploader"
)

file, err := os.Open("resource/demo.jpg")
defer file.Close()
if err != nil {
	return err
}

svc := fileuploader.ImageUploader{Client: client}
resp, result, err := svc.Upload(ctx, file, "demo.jpg", consts.ImageJPG)

```

### 示例程序

为了方便开发者快速上手，微信支付给每个服务生成了示例代码 `api_xx_example_test.go`。请按需查阅。例如：

+ [api_jsapi_example_test.go](services/payments/jsapi/api_jsapi_example_test.go)
+ [api_refunds_example_test.go](services/refunddomestic/api_refunds_example_test.go)

## 发送 HTTP 请求

如果 SDK 还未支持你需要的接口，使用 `core.Client` 的 `GET`、`POST` 等方法发送 HTTP 请求，而不用关注签名、验签等逻辑。

以 [下载微信支付平台证书](https://wechatpay-api.gitbook.io/wechatpay-api-v3/jie-kou-wen-dang/ping-tai-zheng-shu) 为例：

```go
result, err := client.Get(ctx, "https://api.mch.weixin.qq.com/v3/certificates")
```

使用 `core.Client` 发送 HTTP 请求后会得到 `*core.APIResult` 实例。

## 错误处理

以下情况，SDK 发送请求会返回 `error`：

- HTTP 网络错误，如应答接收超时或网络连接失败
- 客户端失败，如生成签名失败
- 服务器端返回了**非** `2xx` HTTP 状态码
- 应答签名验证失败

为了方便使用，SDK 将服务器返回的 `4xx` 和 `5xx` 错误，转换成了 `APIError`。

```go
// 错误处理示例
result, err := client.Get(ctx, "https://api.mch.weixin.qq.com/v3/certificates")
if err != nil {
	if core.IsAPIError(err, "INVALID_REQUEST") { 
		// 处理无效请求 
	}
	// 处理的其他错误
}
```

## 敏感信息加解密

为了保证通信过程中敏感信息字段（如用户的住址、银行卡号、手机号码等）的机密性，

+ 微信支付要求加密上送的敏感信息
+ 微信支付会加密下行的敏感信息

详见 [接口规则 - 敏感信息加解密](https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/min-gan-xin-xi-jia-mi)。

### 使用加解密算法工具包

使用工具包 [utils](utils) 中的函数，手动对敏感信息加解密。

```go
package utils

// EncryptOAEPWithPublicKey 使用公钥加密
func EncryptOAEPWithPublicKey(message string, publicKey *rsa.PublicKey) (ciphertext string, err error)
// EncryptOAEPWithCertificate 使用证书中的公钥加密
func EncryptOAEPWithCertificate(message string, certificate *x509.Certificate) (ciphertext string, err error)

// DecryptOAEP 使用私钥解密
func DecryptOAEP(ciphertext string, privateKey *rsa.PrivateKey) (message string, err error)
```

[rsa_crypto_test.go](utils/rsa_crypto_test.go) 中演示了如何使用以上函数做敏感信息加解密。

### 获取微信支付平台证书

请求的敏感信息，使用微信支付平台证书中的公钥加密。推荐 [使用平台证书下载管理器](FAQ.md#如何在更多地方使用平台证书下载管理器) 获取微信支付平台证书，或者 [下载平台证书](FAQ.md#如何下载微信支付平台证书)。

### 设置 `Wechatpay-Serial` 请求头

请求的敏感信息加密后，在 HTTP 请求头中添加微信支付平台证书序列号 `Wechatpay-Serial`。该序列号用于告知微信支付加密使用的证书。

使用 `core.Client` 的 `Request` 方法来传输自定义 HTTPHeader。

```go
// Request 向微信支付发送请求
//
// 相比于 Get / Post / Put / Patch / Delete 方法，本方法支持设置更多内容
// 特别地，如果需要为当前请求设置 Header，应使用本方法
func (client *Client) Request(
	ctx context.Context,
	method, requestPath string,
	headerParams http.Header,
	queryParams url.Values,
	postBody interface{},
	contentType string,
) (result *APIResult, err error)

// 示例代码
// 微信支付平台证书序列号，对应加密使用的私钥
header.Add("Wechatpay-Serial", "5157F09EFDC096DE15EBE81A47057A72*******")
result, err := client.Request(
	ctx,
	"POST",
	"https://api.mch.weixin.qq.com/v3/profitsharing/receivers/add",
	header,
	nil,
	body,
	"application/json")

```

## 回调通知的验签与解密

1. 使用微信支付平台证书（验签）和商户 APIv3 密钥（解密）初始化 `notify.Handler`
2. 调用 `handler.ParseNotifyRequest` 验签，并解密报文。

### 初始化
+ 方法一（大多数场景）：先手动注册下载器，再获取微信平台证书访问器。

适用场景： 仅需要对回调通知验证签名并解密的场景。例如，基础支付的回调通知。

```go
ctx := context.Background()
// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
err := downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, mchCertificateSerialNumber, mchID, mchAPIV3Key)
// 2. 获取商户号对应的微信支付平台证书访问器
certVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
// 3. 使用证书访问器初始化 `notify.Handler`
handler := notify.NewNotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))
```

+ 方法二：像 [发送请求](#发送请求) 那样使用 `WithWechatPayAutoAuthCipher` 初始化 `core.Client`，然后再用client进行接口调用。

适用场景：需要对回调通知验证签名并解密，并且后续需要使用 Client 的场景。例如，电子发票的回调通知，验签与解密后还需要通过 Client 调用用户填写抬头接口。

```go
ctx := context.Background()
// 1. 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
opts := []core.ClientOption{
	option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
}
client, err := core.NewClient(ctx, opts...)	
// 2. 获取商户号对应的微信支付平台证书访问器
certVisitor := downloader.MgrInstance().GetCertificateVisitor(mchID)
// 3. 使用证书访问器初始化 `notify.Handler`
handler := notify.NewNotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))
// 4. 使用client进行接口调用
// ...
```

+ 方法三：使用本地的微信支付平台证书和商户 APIv3 密钥初始化 `Handler`。

适用场景：首次通过工具下载平台证书到本地，后续使用本地管理的平台证书进行验签与解密。

```go
// 1. 初始化商户API v3 Key及微信支付平台证书
mchAPIv3Key := "<your apiv3 key>"
wechatPayCert, err := utils.LoadCertificate("<your wechat pay certificate>")
// 2. 使用本地管理的微信支付平台证书获取微信支付平台证书访问器
certVisitor := core.NewCertificateMapWithList([]*x509.Certificate{wechatPayCert})
// 3. 使用apiv3 key、证书访问器初始化 `notify.Handler`
handler := notify.NewNotifyHandler(mchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))
```

建议：为了正确使用平台证书下载管理器，你应阅读并理解 [如何使用平台证书下载管理器](FAQ.md#如何使用平台证书下载管理器)。

### 验签与解密

将支付回调通知中的内容，解析为 `payments.Transaction`。

```go
transaction := new(payments.Transaction)
notifyReq, err := handler.ParseNotifyRequest(context.Background(), request, transaction)
// 如果验签未通过，或者解密失败
if err != nil {
	fmt.Println(err)
	return
}
// 处理通知内容
fmt.Println(notifyReq.Summary)
fmt.Println(transaction.TransactionId)
```

将 SDK 未支持的回调消息体，解析至 `map[string]interface{}`。


```go
content := make(map[string]interface{})
notifyReq, err := handler.ParseNotifyRequest(context.Background(), request, content)
// 如果验签未通过，或者解密失败
if err != nil {
	fmt.Println(err)
	return
}
// 处理通知内容
fmt.Println(notifyReq.Summary)
fmt.Println(content)

```

## 自定义签名生成器与验证器

当默认的本地签名和验签方式不适合你的系统时，实现 `Signer` 或者 `Verifier` 来定制签名和验签。

比如，你把商户私钥集中存储，业务系统通过远程调用获得请求签名。

```golang
// 签名器
type CustomSigner struct {
}

func (s *CustomSigner) Sign(ctx context.Context, message string) (*auth.SignatureResult, error) {
    // TODO: 远程调用获取签名信息
    return &auth.SignatureResult{MchID: "xxx", MchCertificateSerialNo: "xxx", Signature: "xxx"}, nil
}

// 校验器
type CustomVerifier struct {
}

func (v *CustomVerifier) Verify(ctx context.Context, serial, message, signature string) error {
    // TODO: 远程调用验签
    return nil
}
```

当你需要使用自定的签名器和校验器时，这样创建客户端

```golang
package core_test

import (
	"context"

	"path/to/your/custom_signer"
	"path/to/your/custom_verifier"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
)

func NewCustomClient(ctx context.Context, mchID string) (*core.Client, error) {
	signer := &custom_signer.CustomSigner{
		// ... 
	}
	verifier := &custom_verifier.CustomVerifier{
		// ...
	}

	opts := []core.ClientOption{
		option.WithSigner(signer),
		option.WithVerifier(verifier),
	}

	return core.NewClient(ctx, opts...)
}
```

## 常见问题

常见问题请见 [FAQ.md](FAQ.md)。

## 如何参与开发

微信支付欢迎来自社区的开发者贡献你们的想法和代码。请你在提交 PR 之前，先提一个对应的 issue 说明以下内容：

- 背景（如，遇到的问题）和目的
- **着重**说明你的想法
- 通过代码或者其他方式，简要的说明是如何实现的，或者它会是如何使用
- 是否影响现有的接口

[#35](https://github.com/wechatpay-apiv3/wechatpay-go/issues/35) 是一个很好的参考。

### 测试

开发者提交的代码，应能通过本 SDK 所有的测试用例。

SDK 在单元测试中使用了 [agiledragon/gomonkey](https://github.com/agiledragon/gomonkey) 和 [stretchr/testify](https://github.com/stretchr/testify)，测试前请确认相关的依赖。使用以下命令获取所有的依赖。

```bash
go get -t -v
```

由于 `gomonkey` 的原因，在执行测试用例时需要携带参数 `-gcflags=all=-l`。使用以下命令发起测试。

```bash
go test -gcflags=all=-l ./...
```

## 联系微信支付
如果你发现了 BUG，或者需要的功能还未支持，或者有任何疑问、建议，欢迎通过 [issue](https://github.com/wechatpay-apiv3/wechatpay-go/issues) 反馈。

也欢迎访问微信支付的 [开发者社区](https://developers.weixin.qq.com/community/pay)。

### 帮助微信支付改进 SDK

为了向广大开发者提供更好的使用体验，微信支付诚挚邀请您反馈使用微信支付 APIv3 SDK中的感受。
您的反馈将对改进 SDK 大有帮助，[点击参与问卷调查](https://wj.qq.com/s2/8774719/ef10/)。
