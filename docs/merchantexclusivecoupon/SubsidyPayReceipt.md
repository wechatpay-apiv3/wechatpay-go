# SubsidyPayReceipt

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubsidyReceiptId** | **string** | 补差付款唯一单号，由微信支付生成，仅在补差付款成功后有返回 | 
**StockId** | **string** | 由微信支付生成，调用创建商家券API成功时返回的唯一批次ID | 
**CouponCode** | **string** | 券的唯一标识 | 
**TransactionId** | **string** | 微信支付下单支付成功返回的订单号 | 
**PayerMerchant** | **string** | 营销补差扣款商户号 | 
**PayeeMerchant** | **string** | 营销补差入账商户号 | 
**Amount** | **int64** | 单位为分，单笔订单补差金额不得超过券的优惠金额，最高补差金额为5000元 &gt; 券的优惠金额定义：   满减券：满减金额即为优惠金额   折扣券：优惠金额 &#x3D; 微信支付订单金额 ÷ 折扣比例 × (1 - 折扣比例)    换购券：不支持 | 
**Description** | **string** | 付款备注描述，查询的时候原样带回 | 
**Status** | [**SubsidyPayReceiptStatus**](SubsidyPayReceiptStatus.md) | 补差付款单据状态 | 
**FailReason** | [**SubsidyPayReceiptFailReason**](SubsidyPayReceiptFailReason.md) | 仅在补差付款失败时，返回告知对应失败的原因 | [可选] 
**SuccessTime** | **time.Time** | 仅在补差付款成功时，返回完成时间。遵循[rfc3339](https://datatracker.ietf.org/doc/html/rfc3339)标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | [可选] 
**OutSubsidyNo** | **string** | 商户侧需保证唯一性。可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号 | 
**CreateTime** | **time.Time** | 补差付款单据创建时间。遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


