# QueryCouponRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**CouponCode** | **string** | 券的唯一标识 | 
**Appid** | **string** | 支持传入与当前调用接口商户号有绑定关系的AppID。支持小程序AppID与公众号AppID。 校验规则：传入的AppID得是与调用方商户号（即请求头里面的商户号）有绑定关系的AppID或传入的AppID得是归属商户号有绑定关系的AppID | 
**Openid** | **string** | OpenID信息，用户在AppID下的唯一标识。 校验规则：传入的OpenID得是调用方商户号（即请求头里面的商户号）有绑定关系的AppID获取的OpenID或传入的OpenID得是归属商户号有绑定关系的AppID获取的OpenID。[获取OpenID文档](https://pay.weixin.qq.com/wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml#part-3) | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


