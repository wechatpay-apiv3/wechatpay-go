# CouponCodeEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Code** | **string** | 商户上传code | 
**CodeStatus** | [**CouponCodeStatus**](CouponCodeStatus.md) | 商户上传code分配状态 | 
**UploadTime** | **time.Time** | code上传成功的时间 | [可选] 
**DispatchedTime** | **time.Time** | 当code状态为被分配时，code被分配的时间 | [可选] 
**Openid** | **string** | 领券用户的OpenID | [可选] 
**Unionid** | **string** | 领券用户的UnionID，具体参见https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/union-id.html | [可选] 
**CouponCode** | **string** | 券的唯一标识 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


