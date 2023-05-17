# DeleteReceiverRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要） | [可选] 
**Appid** | **string** | 微信分配的公众账号ID | 
**SubAppid** | **string** | 微信分配的子商户公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要） | [可选] 
**Type** | [**ReceiverType**](ReceiverType.md) | 枚举值： MERCHANT_ID：商户ID PERSONAL_OPENID：个人openid（由父商户APPID转换得到） PERSONAL_SUB_OPENID：个人sub_openid（由子商户APPID转换得到） | 
**Account** | **string** | 类型是MERCHANT\\_ID时，是商户号 类型是PERSONAL\\_OPENID时，是个人openid [openid获取方法](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/open/pay/chapter3_3_1.shtml#part-13) | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


