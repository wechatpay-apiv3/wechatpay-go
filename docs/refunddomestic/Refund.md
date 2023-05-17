# Refund

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**RefundId** | **string** | 微信支付退款号 | 
**OutRefundNo** | **string** | 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。 | 
**TransactionId** | **string** | 微信支付交易订单号 | 
**OutTradeNo** | **string** | 原支付交易对应的商户订单号 | 
**Channel** | [**Channel**](Channel.md) | 退款渠道 | 
**UserReceivedAccount** | **string** | 取当前退款单的退款入账方，有以下几种情况： 1）退回银行卡：{银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱:支付用户零钱 3）退还商户:商户基本账户商户结算银行账户 4）退回支付用户零钱通:支付用户零钱通 | 
**SuccessTime** | **time.Time** | 退款成功时间，退款状态status为SUCCESS（退款成功）时，返回该字段。遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | [可选] 
**CreateTime** | **time.Time** | 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | 
**Status** | [**Status**](Status.md) | 退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。 | 
**FundsAccount** | [**FundsAccount**](FundsAccount.md) | 退款所使用资金对应的资金账户类型 | [可选] 
**Amount** | [**Amount**](Amount.md) | 金额详细信息 | 
**PromotionDetail** | [**[]Promotion**](Promotion.md) | 优惠退款信息 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


