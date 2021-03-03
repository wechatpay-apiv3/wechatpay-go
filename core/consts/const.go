//Package consts 微信支付api v3 go http-client 常量
package consts

// http request header 相关常量
const (
	Version             = "1.0"                                // client版本
	Authorization       = "Authorization"                      // header中的Authorization字段
	Accept              = "Accept"                             // header中的Accept字段
	ContentType         = "Content-Type"                       // header中的ContentType字段
	ApplicationJSON     = "application/json"                   // ContentType为application/json
	ImageJPG            = "image/jpg"                          // ContentType为image/jpg
	VideoMP4            = "video/mp4"                          // ContentType为video/mp4
	UserAgent           = "User-Agent"                         // header中的UserAgent字段
	UserAgentContent    = "WechatPay-Go-HttpClient/" + Version // UserAgent中的信息
	HeaderAuthorization = "WECHATPAY2-SHA256-RSA2048 mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\"," +
		"signature=\"%s\"" //Authorization信息

)

// http response header 相关常量
const (
	WechatPayTimestamp = "Wechatpay-Timestamp" // 微信支付回包时间戳
	WechatPayNonce     = "Wechatpay-Nonce"     // 微信支付回包随机字符串
	WechatPaySignature = "Wechatpay-Signature" // 微信支付回包签名信息
	WechatPaySerial    = "Wechatpay-Serial"    // 微信支付回包平台序列号
	RequestID          = "Request-Id"          // 微信支付回包请求ID
)

const (
	FormatMessage = "%s\n%s\n%d\n%s\n%s\n"                                           //格式化信息
	Symbols       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 生成随机字符串时依赖所有内容
	NonceLength   = 32                                                               // 随机字符串的长度
	FiveMinute    = 5 * 60                                                           // 回包校验最长时间
	CERTIFICATE   = "CERTIFICATE"                                                    // 证书标识
)
