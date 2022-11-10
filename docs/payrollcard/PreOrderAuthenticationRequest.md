# PreOrderAuthenticationRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Openid** | **string** | 用户OpenID | 
**Appid** | **string** | 当输入服务商AppID时，会校验其与服务商商户号的绑定关系。服务商AppID和与特约商户AppID至少输入一个，且必须要有拉起微工卡小程序时使用的AppID。 | [可选] 
**SubMchid** | **string** | 特约商户号 | 
**SubAppid** | **string** | 当输入特约商户AppID时，会校验其与特约商户号的绑定关系。服务商AppID和与特约商户AppID至少输入一个，且必须要有拉起微工卡小程序时使用的AppID。 | [可选] 
**AuthenticateNumber** | **string** | 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一 | 
**ProjectName** | **string** | 该劳务活动的项目名称 | 
**EmployerName** | **string** | 该工人所属的用工企业 | 
**AuthenticateType** | [**AuthenticationType**](AuthenticationType.md) | 核身类型 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


