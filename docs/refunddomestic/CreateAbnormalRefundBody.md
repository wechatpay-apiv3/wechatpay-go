# CreateAbnormalRefundBody

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 子商户的商户号，由微信支付生成并下发。服务商模式下必须传递此参数 | [可选] 
**OutRefundNo** | **string** | 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。 | 
**Type** | [**Type**](Type.md) | 可选：退款至用户、退款至交易商户银行账户 | 
**BankType** | **string** | 银行类型，采用字符串类型的银行标识，值列表详见[银行类型](https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter&#x3D;4_2)。仅支持招行、交通银行、农行、建行、工商、中行、平安、浦发、中信、光大、民生、兴业、广发、邮储、宁波银行的借记卡。 若退款至用户此字段必填。 | [可选] 
**BankAccount** | **string** | 用户的银行卡账号，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/min-gan-xin-xi-jia-mi)。 若退款至用户此字段必填。 | [可选] 
**RealName** | **string** | 收款用户姓名，该字段需进行加密处理，加密方法详见[敏感信息加密说明](https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/min-gan-xin-xi-jia-mi)。 若退款至用户此字段必填。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


