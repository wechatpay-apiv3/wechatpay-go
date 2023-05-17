# Task

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | 上传任务的主键，唯一定义此资源的标识 | 
**PackageId** | **string** | 号码包唯一标识符。可在微信支付商户平台创建号码包后获得。 | 
**Filename** | **string** | 任务上传的文件名 | 
**CreateTime** | **string** | 时间格式采用遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**UpdateTime** | **string** | 时间格式采用遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35.+08:00表示，北京时间2015年5月20日 13点29分35秒。 | [可选] 
**Status** | [**TaskStatus**](TaskStatus.md) | 上传任务的状态：PROCESSING - 处理中，FINISHED - 已完成 | [可选] 
**SuccessCount** | **int64** | 匹配成功的协议号数。仅当任务状态是FINISHED时，该数据才有效 | [可选] 
**FailCount** | **int64** | 匹配失败的协议号数。仅当任务状态是FINISHED时，该数据才有效 | [可选] 
**SuccessUserCount** | **int64** | 匹配成功的微信用户数，不小于匹配成功的协议号数。当一张银行卡被用户绑定到不同微信号时，才可能出现一个协议号对应多个微信账户的情况。仅当任务状态是FINISHED时，该数据才有效 | [可选] 
**BankType** | **string** | 银行类型，用于标识协议号所属的银行以及卡类型(信用卡或借记卡)。采用字符串类型的银行标识，值列表详见[银行类型](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/terms_definition/chapter1_1_3.shtml#part-7) | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


