# PlateService

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**PlateNumber** | **string** | 车牌号，仅包括省份+车牌，不包括特殊字符。 | 
**PlateColor** | [**PlateColor**](PlateColor.md) | 车牌颜色 | 
**ServiceOpenTime** | **time.Time** | 车牌服务开通时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | [可选] 
**Openid** | **string** | 用户在商户对应appid下的唯一标识，此处返回商户请求中的openid | 
**ServiceState** | **string** | 车牌服务开通状态， NORMAL 正常服务 PAUSE 暂停服务 OUT_SERVICE 未开通 商户根据状态带用户跳转至对应的微信支付分停车服务小程序页面。 其中NORMAL 和 PAUSE状态，可跳转至车牌管理页，进行车牌服务状态管理。OUT_SERVICE状态，可跳转至服务开通页面。 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


