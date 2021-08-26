# CreateOrderReceiver

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Account** | **string** | 1、类型是MERCHANT_ID时，是商户号 2、类型是PERSONAL_OPENID时，是个人openid  3、类型是PERSONAL_SUB_OPENID时，是个人sub_openid | 
**Amount** | **int64** | 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额 | 
**Description** | **string** | 分账的原因描述，分账账单中需要体现 | 
**Name** | **string** | 可选项，在接收方类型为个人的时可选填，若有值，会检查与 name 是否实名匹配，不匹配会拒绝分账请求 1、分账接收方类型是PERSONAL_OPENID或PERSONAL_SUB_OPENID时，是个人姓名的密文（选传，传则校验） 此字段的加密的方式为：敏感信息加密说明 2、使用微信支付平台证书中的公钥 3、使用RSAES-OAEP算法进行加密 4、将请求中HTTP头部的Wechatpay-Serial设置为证书序列号 | [可选] 
**Type** | **string** | 1、MERCHANT_ID：商户号 2、PERSONAL_OPENID：个人openid（由父商户APPID转换得到） 3、PERSONAL_SUB_OPENID: 个人sub_openid（由子商户APPID转换得到） | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


