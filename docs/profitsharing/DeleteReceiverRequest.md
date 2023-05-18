# DeleteReceiverRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Account** | **string** | 类型是MERCHANT_ID时，是商户号 类型是PERSONAL_OPENID时，是个人openid 类型是PERSONAL_SUB_OPENID时，是个人sub_openid | 
**Appid** | **string** | 微信分配的公众账号ID | 
**SubAppid** | **string** | 微信分配的子商户公众账号ID，分账接收方类型包含PERSONAL_SUB_OPENID时必填。（直连商户不需要，服务商需要） | [可选] 
**SubMchid** | **string** | 微信支付分配的子商户号，即分账的出资商户号。（直连商户不需要，服务商需要） | [可选] 
**Type** | [**ReceiverType**](ReceiverType.md) | 枚举值： MERCHANT_ID：商户ID PERSONAL_OPENID：个人openid（由父商户APPID转换得到） PERSONAL_SUB_OPENID：个人sub_openid（由子商户APPID转换得到）  * &#x60;MERCHANT_ID&#x60; - 商户号，  * &#x60;PERSONAL_OPENID&#x60; - 个人openid（由父商户APPID转换得到），  * &#x60;PERSONAL_SUB_OPENID&#x60; - 个人sub_openid（由子商户APPID转换得到）（直连商户不需要，服务商需要）， | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


