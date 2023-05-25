# CouponSendGovCardRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | 消费卡cardi\\_d创建方的AppID，需与API调用方商户号有绑定关系，入参中的用户OpenID也需用此AppID生成 | 
**Openid** | **string** | 待发券用户的OpenID，需为消费卡创建方AppID生成的OpenID | 
**OutRequestNo** | **string** | 商户此次发放凭据号。推荐使用大小写字母和数字，不同添加请求发放凭据号不同，商户侧需保证同一发券请求的out_request_no+send_time唯一性 | 
**SendTime** | **string** | 单次请求发卡时间，消费卡在商户系统的实际发放时间，为东八区标准时间（UTC+8）。商户需保证同一次请求的out_request_no+send_time唯一。由于系统限制，暂不支持传入早于当前时间24小时以上的时间进行发券请求。 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


