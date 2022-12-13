# BusSceneInfo

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | 用户乘车时间（上车），按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE | 
**LineName** | **string** | 乘车路线 路线名称的展示规则：\\^[A-Za-z0-9]{1,10}路$ 如line_name符合上述的规则,扣款凭证将增加展示乘车路线的字段 | [可选] 
**PlateNumber** | **string** | 车牌号。仅包括省份+车牌，不包括特殊字符。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


