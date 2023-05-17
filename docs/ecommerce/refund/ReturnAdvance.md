# ReturnAdvance

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**RefundId** | **string** | 必须是垫付退款的微信退款单 | 
**AdvanceReturnId** | **string** | 微信支付生成的垫付回补操作单号 | 
**ReturnAmount** | **int64** | 退款单对应的垫付退款的金额 | 
**PayerMchid** | **string** | 微信支付分配给出款方的商户号 | 
**PayerAccount** | **string** | BASIC  基本账户 OPERATION 运营账户 | 
**PayeeMchid** | **string** | 微信支付分配给入账方的商户号 | 
**PayeeAccount** | **string** | BASIC  基本账户 OPERATION 运营账户 | 
**Result** | **string** | SUCCESS 回补成功 FAILED 回补失败，出款方账户余额不足时发生 PROCESSING 处理中 | 
**SuccessTime** | **time.Time** | 垫付回补完成的时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


