# ListCouponsByFilterRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Openid** | **string** | Openid信息，用户在AppID下的唯一标识。 校验规则：传入的OpenID得是调用方商户号（即请求头里面的商户号）有绑定关系的AppID获取的OpenID或传入的OpenID得是归属商户号有绑定关系的AppID获取的OpenID。[获取OpenID文档](https://pay.weixin.qq.com/wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml#part-3) | 
**Appid** | **string** | 支持传入与当前调用接口商户号有绑定关系的AppID。支持小程序AppID与公众号AppID。 校验规则：传入的AppID得是与调用方商户号（即请求头里面的商户号）有绑定关系的AppID 或 传入的AppID得是归属商户号有绑定关系的AppID | 
**StockId** | **string** | 批次号，是否指定批次号查询 | [可选] 
**CreatorMerchant** | **string** | 批次创建方商户号 校验规则：当调用方商户号（即请求头中的商户号）为创建批次方商户号时，该参数必传 | [可选] 
**BelongMerchant** | **string** | 批次归属商户号 校验规则：当调用方商户号（即请求头中的商户号）为批次归属商户号时，该参数必传 | [可选] 
**SenderMerchant** | **string** | 批次发放商户号 校验规则：当调用方商户号（即请求头中的商户号）为批次发放商户号时，该参数必传；委托营销关系下，请填写委托发券的商户号 | [可选] 
**Offset** | **int64** | 分页页码 | [可选] 
**Limit** | **int64** | 分页大小 | [可选] 
**CouponState** | [**CouponStatus**](CouponStatus.md) | 商家券状态 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


