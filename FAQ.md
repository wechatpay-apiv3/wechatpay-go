# 常见问题

**目录**

- [证书](#证书)
- [回调](#回调)
- [其他](#其他)

## 证书

### 如何下载微信支付平台证书

现在本 SDK 已经提供了命令行工具供开发者使用。 

首先使用 `go` 指令下载命令行工具
```shell
go get -u github.com/wechatpay-apiv3/wechatpay-go/cmd/wechatpay_download_certs
```
然后执行 `wechatpay_download_certs` 即可下载微信支付平台证书到当前目录
```shell
wechatpay_download_certs -m <mchID> -p <mchPrivateKeyPath> -s <mchSerialNo> -k <mchAPIv3Key>
```
完整参数列表可运行 `wechatpay_download_certs -h` 查看。

### 如何使用平台证书下载管理器

平台证书下载管理器提供已注册商户的微信支付平台证书下载和自动更新。

```go
// GetCertificate 获取商户的某个平台证书
func (mgr *CertificateDownloaderMgr) GetCertificate(ctx context.Context, mchID, serialNo string) (*x509.Certificate, bool)

// GetCertificateVisitor 获取某个商户的平台证书访问器
func (mgr *CertificateDownloaderMgr) GetCertificateVisitor(mchID string) core.CertificateVisitor
```

使用前，先注册商户信息至平台证书下载管理器。

+ 使用 `option.WithWechatPayAutoAuthCipher` 创建 `core.Client`，自动注册。

+ 如果你不需要 `core.Client`，则使用 `downloader.MgrInstance().RegisterDownloaderWithPrivateKey` 手动注册。

以上两种方法，在进程中调用其中之一即可，调用多次无实际意义。如果以上二者都没有调用过，查询该商户的平台证书的结果为空。

如果你希望了解更多，或自行管理微信支付平台证书下载管理器的生命周期，请参阅 [`core/downloader`](core/downloader) 的代码。

### 为什么收到应答中的证书序列号和发起请求的证书序列号不一致

请求和应答使用[数字签名](https://zh.wikipedia.org/wiki/%E6%95%B8%E4%BD%8D%E7%B0%BD%E7%AB%A0)，保证数据传递的真实、完整和不可否认。为了验签方能识别数字签名使用的密钥（特别是密钥和证书更换期间），微信支付 API v3 要求签名和相应的证书序列号一起传输。

+ 商户请求使用**商户API私钥**签名。商户应上送商户证书序列号。
+ 微信支付应答使用**微信支付平台私钥**签名。微信支付应答返回微信支付平台证书序列号。

综上所述，请求和应答的证书序列号是不一致的。

## 回调

### 证书和回调解密需要的AesGcm解密在哪里？

请参考 [aes.go](https://github.com/wechatpay-apiv3/wechatpay-go/blob/main/utils/aes.go) 和 [aes_test.go](https://github.com/wechatpay-apiv3/wechatpay-go/blob/main/utils/aes_test.go)。

由于 SDK 已经提供了微信支付平台证书下载器`downloader.CertificateDownloader`以及回调处理器`notify.Handler`，这两者会完成所有的解析与解密工作。因此除非你想要自定义实现，否则你应该不需要用到`aes.go`中提供的方法。

### 回调验签失败，返回 `crypto/rsa: verification error`

该错误表示回调的证书序列号正确，但验证签名未通过，说明数据不正确。

如果是真实的微信支付回调，请检查是否在 `ParseNotifyRequest` 前消费过 `Request.Body`。`Request.Body` 定义为 `io.Reader`，不支持重复读取。

## 其他

### 如何下载账单
[账单下载API](https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/bill/chapter3_3.shtml) 分成了两个步骤：
1. `/v3/bill/tradebill` 获取账单下载链接和账单摘要
2. `/v3/billdownload/file` 账单文件下载，请求需签名但应答不签名

其中第二步的应答中不包含应答数字签名，无法验签，应使用`WithoutValidator()`**跳过**应答签名的校验。
```go
opts := []core.ClientOption{
	option.WithMerchantCredential(mchID, mchCertificateSerialNumber, privateKey),
	option.WithoutValidator(),
}

client, err := core.NewClient(ctx, opts...)
```

> **注意**：第一步中应正常对应答验证签名
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
