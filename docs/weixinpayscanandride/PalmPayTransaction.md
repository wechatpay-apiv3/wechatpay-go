# PalmPayTransaction

刷掌直连商户的订单结构

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | 商户在微信申请公众号或移动应用成功后分配的账号ID，登录平台为mp.weixin.qq.com或open.weixin.qq.com | 
**Mchid** | **string** | 微信支付分配的商户号 | 
**Description** | **string** | 商户自定义字段，用于交易账单中对扣费服务的描述。 | 
**CreateTime** | **string** | 订单成功创建时返回，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE | 
**OutTradeNo** | **string** | 商户系统内部订单号，只能是数字、大小写字母，且在同一个商户号下唯一 | 
**TransactionId** | **string** | 微信支付订单号 | [可选] 
**TradeState** | [**TradeState**](TradeState.md) | 交易状态，枚举值：SUCCESS—支付成功 ACCEPT—已接收，等待扣款 PAY_FAIL--支付失败(其他原因，如银行返回失败) REFUND—转入退款 | 
**TradeStateDescription** | **string** | 对当前订单状态的描述和下一步操作的指引 | [可选] 
**SuccessTime** | **string** | 订单支付完成时间，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE | [可选] 
**BankType** | **string** | 银行类型，采用字符串类型的银行标识，详细可参考 微信支付银行类型标识（https://pay.weixin.qq.com/wiki/doc/apiv3/terms_definition/chapter1_1_3.shtml）。 特殊标识 BPA：该笔订单由微信进行垫付 | [可选] 
**UserRepayState** | [**UserRepayState**](UserRepayState.md) | 该字段标识对应的垫资单是否已还款或者已撤销。注意，使用此字段前需先确认bank_type字段值为BPA 以及 trade_state字段值为SUCCESS或者REFUND。枚举值如下：NOT_REPAY—用户未还款，REPAID—用户已还款，ALL_REFUND—由于全额退款，无需还款 | [可选] 
**RepayTransactionId** | **string** | 用户还款成功的微信支付订单号。用户已还款会返回该字段。 | [可选] 
**RepayTime** | **string** | 垫资还款时间，该笔订单发生过垫资，并且用户还款成功后，会返回该字段信息，按照使用rfc3339所定义的格式，格式为yyyy-MM-DDThh:mm:ss+TIMEZONE | [可选] 
**Attach** | **string** | 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用 | [可选] 
**TradeScene** | [**PalmPayTradeScene**](PalmPayTradeScene.md) | 交易场景值，目前支持 ：PALM_METRO—刷掌地铁； | [可选] 
**MetroInfo** | [**PalmMetroSceneInfo**](PalmMetroSceneInfo.md) | 返回信息中的trade_scene为PALM_METRO，返回该场景信息 | [可选] 
**Amount** | [**QueryOrderAmount**](QueryOrderAmount.md) | 订单金额信息 | 
**PromotionDetail** | [**[]PromotionDetail**](PromotionDetail.md) | 优惠信息 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


