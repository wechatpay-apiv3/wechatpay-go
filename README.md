# 微信支付 API v3 Go SDK
[微信支付 API v3](https://wechatpay-api.gitbook.io/wechatpay-api-v3/) 官方Go语言客户端代码库。

## 版本信息

版本号：`v0.2.0`

版本功能：
1. 支持微信支付 API v3 请求签名与应答验签能力的 HTTP Client: `core.Client`，该 HTTP Client 在执行请求时将自动携带身份认证信息，并检查应答的微信支付签名。
2. 微信支付各服务API对应的SDK，目前仅包含： 
    1) 微信核心支付4种常用支付接口（JSAPI支付, APP支付，H5支付，Native支付）的SDK 
	2) 微信支付4种文件上传接口的SDK
	3) 更多API跟进中

兼容性：
本版本为测试版本，因为接口重命名/代码结构调整等原因，与之前版本存在不兼容的情况。

## 快速开始
### 安装
#### 1、使用 Go Modules 管理你的项目（如果项目目录下存在`go.mod`文件的话说明你的项目已经使用 Go Modules 进行管理）
如果你的项目还不是使用 Go Modules 进行管理，在项目根目录下执行：
```shell
go mod init
```

#### 2、添加依赖
在 `go.mod` 文件中加入对本SDK的依赖：
```
require (
    github.com/wechatpay-apiv3/wechatpay-go v0.2.0
)
```
并使用`go mod tidy`进行下载。

也可以直接在项目目录中执行： 
```shell
go get -u github.com/wechatpay-apiv3/wechatpay-go@v0.2.0
```
来自动完成`go.mod`的修改与SDK的下载。

### 初始化 Client
为向微信支付 API 发送请求，需要初始化一个 `core.Client` 实例。
```go
package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"log"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

func main() {
	// 示例参数，实际使用时请自行初始化
	var (
		mchID                      string              // 商户号
		mchCertificateSerialNumber string              // 商户API证书序列号
		mchPrivateKey              *rsa.PrivateKey     // 商户API私钥
		wechatPayCertList          []*x509.Certificate // 微信支付平台证书
		customHTTPClient           *http.Client
		customHTTPHeader           http.Header
	)

	ctx := context.Background()
	opts := []core.ClientOption{
		core.WithMerchantCredential(mchID, mchCertificateSerialNumber, mchPrivateKey), // 必要，使用商户信息生成默认 WechatPayCredential
		core.WithWechatPayValidator(wechatPayCertList),                                // 必要，使用微信支付平台证书列表生成默认 WechatPayValidator
		core.WithHTTPClient(customHTTPClient),                                         // 可选，设置自定义 HTTPClient 实例，不设置时使用默认 http.Client{}
		core.WithTimeout(2 * time.Second),                                             // 可选，设置自定义超时时间，不设置时使用 http.Client{} 默认超时
		core.WithHeader(customHTTPHeader),                                             // 可选，设置自定义请求头
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}
	// 接下来使用 client 进行请求发送
	_ = client
}
```
`core.Client`初始化完成后，可以在多个goroutine中并发使用。 

#### 名词解释

+ 商户API证书，是用来证实商户身份的。证书中包含商户号、证书序列号、证书有效期等信息，由证书授权机构(Certificate Authority ，简称CA)签发，以防证书被伪造或篡改。如何获取请见 [商户API证书](https://wechatpay-api.gitbook.io/wechatpay-api-v3/ren-zheng/zheng-shu#shang-hu-api-zheng-shu) 。
+ 商户API私钥。商户申请商户API证书时，会生成商户私钥，并保存在本地证书文件夹的文件apiclient_key.pem中。注：不要把私钥文件暴露在公共场合，如上传到Github，写在客户端代码等。
+ 微信支付平台证书。平台证书是指由微信支付负责申请的，包含微信支付平台标识、公钥信息的证书。商户可以使用平台证书中的公钥进行应答签名的验证。获取平台证书需通过 [获取平台证书列表](https://wechatpay-api.gitbook.io/wechatpay-api-v3/ren-zheng/zheng-shu#ping-tai-zheng-shu) 接口下载。
+ 证书序列号。每个证书都有一个由CA颁发的唯一编号，即证书序列号。如何查看证书序列号请看 [这里](https://wechatpay-api.gitbook.io/wechatpay-api-v3/chang-jian-wen-ti/zheng-shu-xiang-guan#ru-he-cha-kan-zheng-shu-xu-lie-hao) 。

### 发送 HTTP 请求
以 [下载微信支付平台证书](https://wechatpay-api.gitbook.io/wechatpay-api-v3/jie-kou-wen-dang/ping-tai-zheng-shu) 为例：
```go
var (
	result *core.APIResult
	err error
)
result, err = client.Get(ctx, "https://api.mch.weixin.qq.com/v3/certificates")
// 处理 result.Response
```
使用 `core.Client` 发送 HTTP 请求后会得到一个 `core.APIResult` 实例，其中包含了当次请求的请求报文 `*http.Request` 以及应答报文 `*http.Response`。

### 使用服务接口 SDK

#### 以 [JSAPI下单](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_1.shtml) 为例

```go
import (
	"log"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

func JsapiPrepay() error {
	// 请自行初始化 client
	var client *core.Client
	
	svc := jsapi.JsapiApiService{Client: client}
	resp, result, err := svc.Prepay(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String("wxd678efh567hg6787"),
			Mchid:       core.String("1900009191"),
			Description: core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:  core.String("1217752501201407033233368018"),
			Attach:      core.String("自定义数据说明"),
			NotifyUrl:   core.String("https://www.weixin.qq.com/wxpay/pay.php"),
			Amount: &jsapi.Amount{
				Total: core.Int32(100),
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
		return err
	}
}
```

#### 以 [查询订单](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_2.shtml) 为例

```go
import (
	"log"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

func QueryOrder() error {
	// 请自行初始化 client
	var client *core.Client
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
		return err
	}
}
```

#### 以 [图片上传API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_1.shtml) 为例：
```go
import (
	"os"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services/fileuploader"
)

func UploadImage() error {
	// 请自行初始化 client
	var client *core.Client
	
	file, err := os.Open("resource/demo.jpg")
	defer file.Close()
	if err != nil {
		return err
	}
	
	svc := fileuploader.ImageUploader{Client: client}
	resp, result, err := svc.Upload(ctx, file, "demo.jpg", consts.ImageJPG)
	// 处理 resp 
}
```

## 目录介绍
```
github.com/wechatpay-apiv3/wechatpay-go
├── core       // 本代码库核心功能
├── services   // 微信支付各服务API对应的SDK，目前只包含部分API，更多服务跟进中
├── docs       // 以上微信支付各服务API对应的SDK的说明文档
├── utils      // 各种常用工具函数
└── README.md  // 本文档
```

## 敏感信息加解密
为了保证通信过程中敏感信息字段（如用户的住址、银行卡号、手机号码等）的机密性，微信支付 API v3 要求商户对上送的敏感信息字段进行加密。
与之相对应，微信支付会对下行的敏感信息字段进行加密，商户需解密后方能得到原文。
详见[这里](https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/min-gan-xin-xi-jia-mi)。

### 加解密算法工具
工具包 `github.com/wechatpay-apiv3/wechatpay-go/utils` 中提供了加解密所需要的函数，可以使用这些方法手动对敏感信息字段进行加解密
```go
package utils

// EncryptOAEPWithPublicKey 使用公钥进行加密
func EncryptOAEPWithPublicKey(message string, publicKey *rsa.PublicKey) (ciphertext string, err error)
// EncryptOAEPWithCertificate 先解析出证书中的公钥，然后使用公钥进行加密
func EncryptOAEPWithCertificate(message string, certificate *x509.Certificate) (ciphertext string, err error)

// DecryptOAEP 使用私钥进行解密
func DecryptOAEP(ciphertext string, privateKey *rsa.PrivateKey) (message string, err error)
```

### [建设中] 服务 SDK 自动加解密
未来会在服务 SDK 中集成请求/应答隐私字段自动加解密功能，省去手动对敏感信息字段进行处理的繁琐工作。
尽情期待。

## 自定义签名生成器与验证器
当默认的本地签名和验签方式不适合你的系统时，你可以通过实现`Signer`或者`Verifier`来定制签名和验签。
比如，你可以把商户私钥集中存储，业务系统通过远程调用进行签名，你可以这样做。
```golang
// 签名器
type CustomSigner struct {
}

func (customSigner *CustomSigner) Sign(ctx context.Context, message string) (*auth.SignatureResult, error) {
    // TODO: 远程调用获取签名信息
    return &auth.SignatureResult{MchCertificateSerialNo: "xxx", Signature: "xxxx"}, nil
}

// 校验器
type CustomVerifier struct {
}

func (customSigner *CustomVerifier) Verify(ctx context.Context, serial, message, signature string) error {
    // TODO: 远程调用验签
    return nil
}
```

当你需要使用自定的签名器和校验器时，你可以这样来创建客户端
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
	credential := &credentials.WechatPayCredentials{
		Signer: &custom_signer.CustomSigner{},
		MchID:  mchID,
	}
	validator := &validators.WechatPayValidator{
		Verifier: &custom_verifier.CustomVerifier{},
	}

	opts := []core.ClientOption{
		core.WithCredential(credential),
		core.WithValidator(validator),
	}

	return core.NewClient(ctx, opts...)
}
```

## 常见问题
### 如何下载平台证书
首次下载平台证书时，用户还没有平台证书，此时无法完成验签行为。
可以使用如下方法**临时**跳过应答签名验证：
```go
opts := []core.ClientOption{
	core.WithMerchantCredential(mchID, mchCertificateSerialNumber, privateKey),
	core.WithoutValidator(),
}

client, err := core.NewClient(ctx, opts...)
```
使用 `core.WithoutValidator()` 可以构建一个空验证器 `NullValidator`，该验证器不会对应答进行检查。
> **注意**：为了保证请求安全，请务必在一般业务请求中使用正确的签名验证器

### 如何下载账单
[账单下载API](https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/bill/chapter3_3.shtml) 分成了两个步骤：
1. `/v3/bill/tradebill` 获取账单下载链接和账单摘要
2. `/v3/billdownload/file` 账单文件下载，请求需签名但应答不签名

其中第二步的应答中不包含应答数字签名，我们可以参考上一个问题下载平台证书的方法，使用`WithoutValidator()`**跳过**应答签名的校验。

> **注意**：第一步中应正常对应答签名进行验证
> 
> **注意**：开发者在下载文件之后，应使用第一步获取的账单摘要校验文件的完整性

### 如何查看 HTTP 请求的 Request 信息
不论是使用 `Client` 的 HTTP 方法（`Get/Post/Put/Delete`等）直接发送 HTTP 请求，还是使用服务API对应的SDK发起请求，均会返回 `*core.APIResult` 结构。
该结构中包含了本次发起 HTTP 请求的 `http.Request` 对象和微信支付应答的 `http.Response` 对象。
```golang
var (
	request *http.Request
	response *http.Response
)
result, err = client.Get(ctx, "")

request = result.Request
response = result.Response 
```

## 联系我们
如果你发现了BUG或者有任何疑问、建议，欢迎通过 [issue](https://github.com/wechatpay-apiv3/wechatpay-go/issues) 进行反馈

也欢迎访问我们的[开发者社区](https://developers.weixin.qq.com/community/pay)
