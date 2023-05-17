# CreateRefundRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 微信支付分配二级商户的商户号 | 
**SpAppid** | **string** | 电商平台在微信公众平台申请服务号对应的APPID，申请商户功能的时候微信支付会配置绑定关系 | 
**SubAppid** | **string** | 二级商户在微信申请公众号成功后分配的帐号ID，需要电商平台侧有配置绑定关系才能传 | [可选] 
**TransactionId** | **string** | 原支付交易对应的微信订单号 | [可选] 
**OutTradeNo** | **string** | 原支付交易对应的商户订单号 | [可选] 
**OutRefundNo** | **string** | 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。 | 
**Reason** | **string** | 若商户传入，会在下发给用户的退款消息中体现退款原因 | [可选] 
**Amount** | [**RefundReqAmount**](RefundReqAmount.md) | 订单金额信息 | 
**NotifyUrl** | **string** | 异步接收微信支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。 | [可选] 
**RefundAccount** | **string** | REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付，需要向微信支付申请开通 REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值 注意：若传入REFUND_SOURCE_PARTNER_ADVANCE，仅代表可以使用垫付退款，实际出款账户需以退款申请受理结果或查单结果为准。 | [可选] 
**FundsAccount** | **string** | 若订单处于待分账状态，可以传入此参数，指定退款资金来源账户。当该字段不存在时，默认使用订单交易资金所在账户出款，即待分账时使用不可用余额的资金进行退款，已分账或无分账时使用可用余额的资金进行退款。 AVAILABLE：可用余额  | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


