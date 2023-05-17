# Transaction

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | appid是商户在微信申请公众号或移动应用成功后分配的帐号ID，登录平台为mp.weixin.qq.com或open.weixin.qq.com | 
**SubAppid** | **string** | 子商户申请的公众号或移动应用appid，需要在服务商的商户平台为子商户绑定 | [可选] 
**SpMchid** | **string** | 微信支付分配的商户号 | 
**SubMchid** | **string** | 微信支付分配的子商户号 | [可选] 
**Description** | **string** | 商户自定义字段，用于交易账单中对扣费服务的描述。 | 
**CreateTime** | **time.Time** | 订单成功创建时返回，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**OutTradeNo** | **string** | 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一 | 
**TransactionId** | **string** | 微信支付订单号 | [可选] 
**TradeState** | **string** | SUCCESS—支付成功 ACCEPTED—已接收，等待扣款 PAY_FAIL–支付失败(其他原因，如银行返回失败) REFUND—转入退款 | 
**TradeStateDescription** | **string** | 对当前订单状态的描述和下一步操作的指引 | [可选] 
**SuccessTime** | **time.Time** | 订单支付完成时间，遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | [可选] 
**BankType** | **string** | 银行类型，采用字符串类型的银行标识。BPA：该笔订单由微信进行垫付 | [可选] 
**UserRepaid** | **string** | 枚举值： Y：用户已还款 N：用户未还款 注意：使用此字段前需先确认bank_type字段值为BPA以及 trade_state字段值为SUCCESS。 | [可选] 
**Attach** | **string** | 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用 | [可选] 
**TradeScene** | **string** | 交易场景值，目前支持 ：PARKING：车场停车场景  | 
**ParkingInfo** | [**ParkingTradeScene**](ParkingTradeScene.md) | 返回信息中的trade_scene为PARKING，返回该场景信息 | [可选] 
**Payer** | [**Payer**](Payer.md) | 支付者信息 | 
**Amount** | [**QueryOrderAmount**](QueryOrderAmount.md) | 订单金额信息 | 
**PromotionDetail** | [**[]PromotionDetail**](PromotionDetail.md) | 优惠信息 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


