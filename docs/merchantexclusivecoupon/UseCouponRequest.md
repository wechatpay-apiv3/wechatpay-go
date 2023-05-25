# UseCouponRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**CouponCode** | **string** | 券的唯一标识 | 
**StockId** | **string** | 微信为每个商家券批次分配的唯一ID，当你在创建商家券接口中的coupon\\_code\\_mode参数传值为MERCHANT\\_API或者MERCHANT\\_UPLOAD时，则核销接口中该字段必传，否则该字段可不传 | [可选] 
**Appid** | **string** | 支持传入与当前调用接口商户号有绑定关系的AppID。支持小程序AppID与公众号AppID。核销接口返回的OpenID会在该传入AppID下进行计算获得。 校验规则：传入的AppID得是与调用方商户号（即请求头里面的商户号）有绑定关系的AppID或传入的AppID得是归属商户号有绑定关系的AppID | 
**UseTime** | **string** | 商户请求核销用户券的时间。 遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**UseRequestNo** | **string** | 每次核销请求的唯一标识，商户需保证唯一 | 
**Openid** | **string** | 用户的唯一标识，做安全校验使用，非必填。 校验规则：传入的OpenID得是调用方商户号（即请求头里面的商户号）有绑定关系的AppID获取的OpenID或传入的OpenID得是归属商户号有绑定关系的AppID获取的OpenID。[获取OpenID文档](https://pay.weixin.qq.com/wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml#part-3) | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


