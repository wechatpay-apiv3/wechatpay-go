
# wechatpay-go-httpclient

## 概览
[微信支付API v3](https://wechatpay-api.gitbook.io/wechatpay-api-v3/) 的Golang HttpClient扩展，实现了请求签名的生成和应答签名的验证。

如果你是使用Golang的商户开发者，可以使用它构造HttpClient, 该HttpClient在执行请求时将自动携带身份认证信息，并检查应答的微信支付签名。

## 项目状态

当前版本为0.1.0为测试版本。请商户的专业技术人员在使用时注意系统和软件的正确性和兼容性，依旧带来的风险。

## 环境要求

- Golang 1.12+

## 安装


运行`go get -u github.com/wechatpay-apiv3/wechatpay-go`

或者在`go.mod`文件中添加
```golang
require (
    github.com/wechatpay-apiv3/wechatpay-go  v0.1.0
)
```
然后运行`go mod tidy`进行下载
## 开始

首先通过创建一个微信支付`apiv3`的`client`，创建的同时需要将各个属性进行赋值。其中，你可以使用自己已经创建好的`http.client`,
如果你没有传入一个`http.client`, 我们会你创建一个默认的`http.client`。
### 初始化客户端
```golang
import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/wechatpay-apiv3/wechatpay-go/core"
    "github.com/wechatpay-apiv3/wechatpay-go/core/option"
    "github.com/wechatpay-apiv3/wechatpay-go/utils"
)
func main() {
    // 增加客户端配置
    ctx := context.Background()
    opts := []option.ClientOption{
        option.WithMerchant(mchID, mchCertificateSerialNumber, privateKey), // 设置商户信息，用于生成签名信息
        option.WithWechatPay(wechatPayCertificateList),  // 设置微信支付平台证书信息，对回包进行校验
        option.WithHTTPClient(&http.Client{}),  // 可以不设置
        option.WithTimeout(2 * time.Second),    // 自行进行超时时间配置
    }
    client, err := core.NewClient(ctx, opts...)
    if err != nil {
        log.Printf("new wechat pay client err:%s", err.Error())
        return
    }
    // 后面可以开始写你的逻辑
}
```

### 发送请求
```golang
getURL := "https://api.mch.weixin.qq.com/v3/certificates"
response, err := client.Get(ctx, getURL)
if err != nil {
    log.Printf("weChatPayHttpClient Get:%s err:%s", getURL, err.Error())
    return
}
// 你还可以校验回包是否正常，如果有错可能是这样
// wechat pay sdk: error HTTP response
//        StatusCode: 400 
//        Code: INVALID_REQUEST 
//        Message: 商户号无效，请核实
if err := core.CheckResponse(response); err != nil{
    log.Printf("weChatPayHttpClient Get:%s Response Err:%s", getURL, err.Error())
    return
}
if response.Body != nil {
    defer response.Body.Close()
}
body, err = ioutil.ReadAll(response.Body)
if err != nil {
    log.Printf("read response body err:%s", err.Error())
    return
}
log.Printf("response body:%s", string(body))
```

## 定制

当默认当本地签名和验签方式不适合你当系统时，你可以通过实现`Signer`或者`Verifier`来定制签名和验签。
比如，你可以把商户私钥集中存储，业务系统通过远程调用进行签名，你可以这样做。
```golang
// 签名器
type CustomSigner struct{
}

func (customSigner *CustomSigner) Sign(ctx context.Context, message string) (string, error) {
    // 调用签名RPC服务，然后返回签名信息
    return "", nil
}

// 校验器
type CustomVerifier struct {
}

func (customSigner *CustomVerifier) Verify(ctx context.Context, serial, message, signature string) error {
    // 你的校验逻辑
    return nil
}
```

当你需要使用自定的签名器和校验器时，你可以这样来创建客户端
```golang
import (
    "xxxx/xxxx/youPackageName"

    "github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
    "github.com/wechatpay-apiv3/wechatpay-go/core/auth/validators"
    "github.com/wechatpay-apiv3/wechatpay-go/core/option"
    "github.com/wechatpay-apiv3/wechatpay-go/core"
)

credential := &credentials.WechatPayCredentials{
    Signer: &yourPackageName.CustomSigner{}
    MchID:               mchID,                   // 你的商户号
    CertificateSerialNo: mchCertificateSerialNumber, // 你的商户证书序列号
}
validator := &validators.WechatPayValidator{
    Verifier: &yourPackageName.CustomVerifier{}
}

opts := []option.ClientOption{
    option.WithCredential(credential),
    option.WithValidator(validator),
}

client, err := core.NewClient(ctx, opts...)
```


## 敏感信息加密

开发者应当使用微信支付平台证书中的公钥，对上送的敏感信息进行加密。这样只有拥有私钥的微信支付才能对密文进行解密，从而保证了信息的机密性。

另一方面，微信支付使用 商户证书中的公钥对下行的敏感信息进行加密。开发者应使用商户私钥对下行的敏感信息的密文进行解密。
### 加密

使用`utils.EncryptOAEPWithPublicKey(message string, publicKey *rsa.PublicKey)`直接使用公钥进行加密

或者

使用`utils.EncryptOAEPWithCertificate(message string, certificate *x509.Certificate)`使用证书进行加密（这里方法将会先解析出证书中的公钥）


### 解密

使用`utils.DecryptOAEP(ciphertext string, privateKey *rsa.PrivateKey)`进行私钥解密


## 图片/视频上传
当前进行图片上传时需要自己拼装body部分，我们我们将会这拼装过程封装成api供大家使用，尽情期待。

示例代码如下
```golang
uploadURL := "https://api.mch.weixin.qq.com/v3/merchant/media/upload"
filePath := ""
fileName := ""
type Meta struct {
    FileName string `json:"filename" binding:"required"` // 商户上传的媒体图片的名称，商户自定义，必须以JPG、BMP、PNG为后缀。
    Sha256   string `json:"sha256" binding:"required"`   // 图片文件的文件摘要，即对图片文件的二进制内容进行sha256计算得到的值。
}
// 读取文件
pictureByes, err := ioutil.ReadFile(filePath)
if err != nil {
    log.Printf("read file err:%s", err.Error())
    return
}
// 计算文件序列化后的sha256
h := sha256.New()
if _, err = h.Write(pictureByes); err != nil {
    return
}

meta := &Meta{}
pictureSha256 := h.Sum(nil)
meta.FileName = fileName
meta.Sha256 = fmt.Sprintf("%x", string(pictureSha256))
metaByte, err := json.Marshal(meta)
if err != nil {
    log.Printf("json marshal meta info err:%s", err.Error())
    return
}
reqBody := &bytes.Buffer{}
writer := multipart.NewWriter(reqBody)
// 设置reqbody中的meta部分
if err = core.CreateFormField(writer, "meta", "application/json", metaByte); err != nil {
    log.Printf("read form field err:%s", err.Error())
    return
}
// 设置reqbody中的file部分
if err = core.CreateFormFile(writer, fileName, "image/jpg", pictureByes); err != nil {
    log.Printf("read form file err:%s", err.Error())
    return
}
if err = writer.Close(); err != nil {
    log.Printf("close io.writer err:%s", err.Error())
    return
}
response, err = client.Upload(ctx, uploadURL, string(metaByte), reqBody.String(), writer.FormDataContentType())
if err != nil {
    log.Printf("weChatPayHttpClient Upload:%s uploadBody:%+v Response Err:%s",
        uploadURL, reqBody.String(), err.Error())
    return
}
// 校验回包
err = core.CheckResponse(response)
if err != nil {
    log.Printf("weChatPayHttpClient Upload uploadBody:%+v CheckResponse Err:%s", reqBody.String(),
        err.Error())
    return
}
if response.Body != nil {
    defer response.Body.Close()
}
// 获取回包中的信息
body, err = ioutil.ReadAll(response.Body)
if err != nil {
    log.Printf("read response body err:%s", err.Error())
    return
}
log.Printf("upload success response body:%s", string(body))
```

## 常见问题
### 如何下载平台证书
使用`core.NewClient(ctx, opts...)`需要通过设置`Validator
`属性添加微信支付平台证书来对回包进行验证，而平台证书又只能通过调用获取平台证书接口下载。为了解开"死循环"，你可以在第一次下载平台证书时，按照下述方法临时"跳过”应答签名的验证。
```golang
//初始化的时候使用不校验回包的verifier
validator := &validators.WechatPayValidator{
    Verifier: &verifiers.WechatPayDefaultVerifier{},
}

opts := []option.ClientOption{
    option.WithMerchant(mchID, mchCertificateSerialNumber, privateKey),
    option.WithValidator(validator),
}

client, err := core.NewClient(ctx, opts...)

```
**注意**：业务请求请使用标准的初始化流程，务必验证应答签名。


### 如何查看http请求的request信息
```golang
// 你可以直接调用来获取request信息，便于调试
request := client.RequestInfo() 
```

### 证书和回调解密需要的AesGcm解密在哪里？
请参考`utils/aes.go`

### 示例demo
可以查看`github.com/wechatpay-apiv3/wechatpay-go/core/client_test.go`里面的demo。

### 工具包
我们提供了一些常用的方法来加载你的私钥和证书,这些都在`utils`工具包中,具体包括以下常用方法
```golang
// utils/aes.go
    // 解密回包中的证书
    DecryptToString(apiv3Key, associatedData, nonce, ciphertext string) 

// utils/pem.go
    // 通过读取的证书字符串内容加载证书
    LoadCertificate(certificateStr string) (certificate *x509.Certificate, err error)
    // 通过读取的私钥字符串内容加载私钥
    LoadPrivateKey(privateKeyStr string) (privateKey *rsa.PrivateKey, err error)
    // 通过读取的公钥字符串内容加载公钥
    LoadPublicKey(publicKeyStr string) (publicKey *rsa.PublicKey, err error)
    // 通过证书的路径加载证书    
    LoadCertificateWithPath(path string) (certificate *x509.Certificate, err error)
    // 通过私钥的路径加载私钥
    LoadPrivateKeyWithPath(path string) (privateKey *rsa.PrivateKey, err error)
    // 通过公钥的路径加载公钥
    LoadPublicKeyWithPath(path string) (publicKey *rsa.PublicKey, err error)

// utils/rsa_crypto.go
    // 直接使用公钥进行加密
    EncryptOAEPWithPublicKey(message string, publicKey *rsa.PublicKey) (ciphertext string, err error)
    // 先解析出证书中的公钥，然后使用公钥进行加密
    EncryptOAEPWithCertificate(message string, certificate *x509.Certificate) (ciphertext string, err error)
    // 使用私钥进行解密
    DecryptOAEP(ciphertext string, privateKey *rsa.PrivateKey) (message string, err error)
```

## 联系我们
如果你发现了BUG或者有任何疑问、建议，请通过issue进行反馈。

也欢迎访问我们的[开发者社区](https://developers.weixin.qq.com/community/pay)
