# ParkingTradeScene

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**ParkingId** | **string** | 微信支付分停车服务为商户分配的入场id，商户通过入场通知接口获取入场id | 
**PlateNumber** | **string** | 车牌号，仅包括省份+车牌，不包括特殊字符。 | 
**PlateColor** | [**PlateColor**](PlateColor.md) | 车牌颜色 | 
**StartTime** | **time.Time** | 用户入场时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**EndTime** | **time.Time** | 用户出场时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**ParkingName** | **string** | 所在停车位车场的名称 | 
**ChargingDuration** | **int64** | 计费的时间长，单位为秒 | 
**DeviceId** | **string** | 停车场设备id | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


