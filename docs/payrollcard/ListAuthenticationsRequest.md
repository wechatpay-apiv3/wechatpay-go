# ListAuthenticationsRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Openid** | **string** | 微信用户OpenID | 
**Appid** | **string** | 当输入服务商AppID时会校验其与服务商商户号绑定关系。 服务商AppID和特约商户AppID至少输入一个。 | [可选] 
**SubAppid** | **string** | 当输入特约商户AppID时会校验其与特约商户号关系。 特约商户AppID和服务商AppID至少输入一个。 | [可选] 
**SubMchid** | **string** | 特约商户号 | 
**AuthenticateDate** | **string** | 核身日期，一次只能查询一天，最久可查询90天内的记录，格式为“yyyy-MM-DD” | 
**AuthenticateState** | **string** | 核身状态，列表查询仅提供成功状态的核身记录查询，故此字段固定默认值即可 | [可选] 
**Offset** | **int64** | 非负整数，表示该次请求资源的起始位置，从0开始计数。调用方选填，默认为0。offset为10，limit为10时，查询第10-19条数据 | [可选] 
**Limit** | **int64** | 非0非负的整数，该次请求可返回的最大资源条数，默认值为10，最大支持10条。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


