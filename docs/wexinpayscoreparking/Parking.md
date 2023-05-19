# Parking

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Id** | **string** | 车主服务为商户分配的入场id | 
**OutParkingNo** | **string** | 商户侧入场标识id，在同一个商户号下唯一 | 
**PlateNumber** | **string** | 车牌号，仅包括省份+车牌，不包括特殊字符。 | 
**PlateColor** | [**PlateColor**](PlateColor.md) | 车牌颜色 | 
**StartTime** | **string** | 入场时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE | 
**ParkingName** | **string** | 所在停车位车场的名称 | 
**FreeDuration** | **int64** | 停车场的免费停车时长，单位为秒 | 
**State** | **string** | 本次入场车牌的服务状态， NORMAL:正常状态，可以使用微信支付分停车服务， BLOCKED 不可用状态，暂时不可以使用微信支付分停车服务 | 
**BlockReason** | **string** | block服务状态描述，返回车牌状态为BLOCKED，会返回该字段，描述具体BLOCKED的原因， PAUSE：已暂停微信支付分停车服务； OVERDUE：已授权签约但欠费，不能提供服务, OUT_SERVICE : 车牌未开通微信支付分停车服务， EVALUATION_FAILED 综合评估未通过，用户支付分不可用的情况，会返回该状态。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


