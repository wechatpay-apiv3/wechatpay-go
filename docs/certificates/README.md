# 微信支付 API v3 Go SDK - certificates

为了确保在定期更换平台证书时，不影响商户使用微信支付的各种功能，微信支付提供API接口供商户下载最新的平台证书。
商户可使用该接口实现平台证书的平滑切换。


## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.0.0

想获取更多信息，请访问 [https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml](https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml)

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*CertificatesApi* | [**DownloadCertificates**](CertificatesApi.md#downloadcertificates) | **Get** /v3/certificates | 获取平台证书列表


## 类型列表

 - [Certificate](Certificate.md)
 - [DownloadCertificatesResponse](DownloadCertificatesResponse.md)
 - [EncryptCertificate](EncryptCertificate.md)

