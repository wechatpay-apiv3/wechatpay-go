# CreateParkingRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 微信支付分配的子商户号，服务商模式下必传 | [可选] 
**OutParkingNo** | **string** | 商户侧入场标识id，在同一个商户号下唯一 | 
**PlateNumber** | **string** | 车牌号，仅包括省份+车牌，不包括特殊字符。 | 
**PlateColor** | [**PlateColor**](PlateColor.md) | 车牌颜色 | 
**NotifyUrl** | **string** | 接受入场状态变更回调通知的url，注意回调url只接受https | 
**StartTime** | **time.Time** | 入场时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**ParkingName** | **string** | 所在停车位车场的名称 | 
**FreeDuration** | **int64** | 停车场的免费停车时长，单位为秒 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


