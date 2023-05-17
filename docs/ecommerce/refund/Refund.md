# Refund

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**RefundId** | **string** | 微信支付退款订单订单号 | 
**OutRefundNo** | **string** | 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。 | 
**TransactionId** | **string** | 微信支付交易订单号 | 
**OutTradeNo** | **string** | 返回的原交易订单号 | 
**Channel** | **string** | ORIGINAL—原路退款 BALANCE—退回到余额 OTHER_BALANCE—原账户异常退到其他余额账户 OTHER_BANKCARD—原银行卡异常退到其他银行卡 | [可选] 
**UserReceivedAccount** | **string** | 取当前退款单的退款入账方 1）退回银行卡： {银行名称}{卡类型}{卡尾号} 2）退回支付用户零钱: 支付用户零钱 3）退还商户: 商户基本账户 商户结算银行账户 4）退回支付用户零钱通: 支付用户零钱通 | [可选] 
**SuccessTime** | **time.Time** | 退款成功时间，退款状态status为SUCCESS（退款成功）时，返回该字段。遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | [可选] 
**CreateTime** | **string** | 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。 | 
**Status** | **string** | 退款状态： SUCCESS—退款成功 CLOSED—退款关闭。 PROCESSING—退款处理中 ABNORMAL—退款异常，退款到银行发现用户的卡作废或者冻结了，导致原路退款银行卡失败，可前往商户平台（pay.weixin.qq.com）-交易中心，手动处理此笔退款。 | 
**Amount** | [**RefundAmount**](RefundAmount.md) | 退款金额信息 | 
**PromotionDetail** | [**[]PromotionDetail**](PromotionDetail.md) | 优惠退款信息 | [可选] 
**RefundAccount** | **string** | REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付 REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值 注意：申请退款时可能会出现暂未扣除退款资金的情况，此时不返回本字段，需继续查单确认。 | [可选] 
**FundsAccount** | **string** | UNSETTLED : 未结算资金 AVAILABLE : 可用余额 UNAVAILABLE : 不可用余额 OPERATION : 运营户 BASIC : 基本账户（含可用余额和不可用余额） ECNY_BASIC : 数字人民币基本账户 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


