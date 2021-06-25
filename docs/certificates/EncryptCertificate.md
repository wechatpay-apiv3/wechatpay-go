# EncryptCertificate

为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密 

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | 加密所使用的算法，目前可能取值仅为 AEAD_AES_256_GCM  | 
**Nonce** | **string** | 加密所使用的随机字符串  | 
**AssociatedData** | **string** | 附加数据包（可能为空）  | 
**Ciphertext** | **string** | 证书内容密文，解密后会获得证书完整内容  | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


