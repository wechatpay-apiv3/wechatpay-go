# AuthenticationEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Mchid** | **string** | 服务商商户号 | 
**SubMchid** | **string** | 特约商户商户号 | 
**Openid** | **string** | 用户OpenID | 
**AuthenticateScene** | [**AuthenticationScene**](AuthenticationScene.md) | 核身渠道，发起核身时的来源渠道，如通过小程序，硬件设备等 | 
**AuthenticateSource** | **string** | 核身渠道标识，用于定位渠道具体来源，如果是扫码打卡渠道标识就是具体的小程序AppID，若是硬件设备，则是设备的序列号等 | 
**ProjectName** | **string** | 项目名称 | 
**EmployerName** | **string** | 该工人所属的用工企业，由商户核身下单时传入 | 
**AuthenticateState** | [**AuthenticationState**](AuthenticationState.md) | 核身状态 | 
**AuthenticateTime** | **string** | 核身时间，遵循RFC3339标准格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。 | [可选] 
**AuthenticateNumber** | **string** | 商户系统内部的商家核身单号，要求此参数只能由数字、大小写字母组成，在服务商内部唯一 | 
**AuthenticateFailedReason** | **string** | 结果为核身失败时的原因描述，仅在失败记录返回 | [可选] 
**AuthenticateType** | [**AuthenticationType**](AuthenticationType.md) | 核身类型 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


