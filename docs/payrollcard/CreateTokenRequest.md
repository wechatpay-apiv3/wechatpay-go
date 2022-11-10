# CreateTokenRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Openid** | **string** | 用户OpenID | 
**Appid** | **string** | 当输入服务商AppID时，会校验其与服务商商户号的绑定关系。 服务商AppID和与特约商户AppID至少输入一个，且必须要有拉起微工卡小程序时使用的AppID。 | [可选] 
**SubMchid** | **string** | 特约商户号 | 
**SubAppid** | **string** | 当输入特约商户AppID时，会校验其与特约商户号的绑定关系。 服务商AppID和与特约商户AppID至少输入一个，且必须要有拉起微工卡小程序时使用的AppID。 | [可选] 
**UserName** | **string** | 用户实名信息，按照APIV3标准加密该字段 | 
**IdCardNumber** | **string** | 用户证件号，按照APIV3标准加密该字段 | 
**EmploymentType** | [**EmploymentType**](EmploymentType.md) | 微工卡服务仅支持用于与商户有用工关系的用户，需明确用工类型；参考值：长期用工：LONG_TERM_EMPLOYMENT，短期用工：SHORT_TERM_EMPLOYMENT，合作关系：COOPERATION_EMPLOYMENT | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


