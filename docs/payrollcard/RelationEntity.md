# RelationEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Openid** | **string** | 用户OpenID | 
**Mchid** | **string** | 商户号 | 
**SubMchid** | **string** | 特约商户号 | 
**AuthorizeState** | **string** | 授权状态：UNAUTHORIZED：未授权 AUTHORIZED：已授权 DEAUTHORIZED：已取消授权 | 
**AuthorizeTime** | **string** | 授权时间 | [可选] 
**DeauthorizeTime** | **string** | 取消授权时间，遵循RFC3339标准格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。 | [可选] 
**RegisterState** | [**PayrollCardRegisterState**](PayrollCardRegisterState.md) | 微工卡的开通状态：UNREGISTERED：未开通；REGISTERED：已开通；CLOSED：已关闭 | 
**RegisterTime** | **string** | 开通时间，未开通时不返回。遵循RFC3339标准格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。 | [可选] 
**CloseTime** | **string** | 关闭时间，仅当开通状态是已关闭时返回。遵循RFC3339标准格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


