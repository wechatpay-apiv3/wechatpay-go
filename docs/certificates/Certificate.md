# Certificate

微信支付平台证书信息

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SerialNo** | **string** | 证书序列号  | 
**EffectiveTime** | **time.Time** | 证书有效期开始时间  | 
**ExpireTime** | **time.Time** | 证书过期时间  | 
**EncryptCertificate** | [**EncryptCertificate**](EncryptCertificate.md) | 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密  | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


