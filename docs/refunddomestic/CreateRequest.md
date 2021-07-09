# CreateRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 子商户的商户号，由微信支付生成并下发。服务商模式下必须传递此参数  | [可选] 
**TransactionId** | **string** | 原支付交易对应的微信订单号  | [可选] 
**OutTradeNo** | **string** | 原支付交易对应的商户订单号  | [可选] 
**OutRefundNo** | **string** | 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。  | 
**Reason** | **string** | 若商户传入，会在下发给用户的退款消息中体现退款原因  | [可选] 
**NotifyUrl** | **string** | 异步接收微信支付退款结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。 如果参数中传了notify_url，则商户平台上配置的回调地址将不会生效，优先回调当前传的这个地址。  | [可选] 
**FundsAccount** | [**ReqFundsAccount**](ReqFundsAccount.md) | 若传递此参数则使用对应的资金账户退款，否则默认使用未结算资金退款（仅对老资金流商户适用）  枚举值： - AVAILABLE：可用余额账户    * &#x60;AVAILABLE&#x60; - 可用余额  | [可选] 
**Amount** | [**AmountReq**](AmountReq.md) | 订单金额信息  | 
**GoodsDetail** | [**[]GoodsDetail**](GoodsDetail.md) | 指定商品退款需要传此参数，其他场景无需传递  | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


