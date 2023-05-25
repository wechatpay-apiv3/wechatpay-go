# SubsidyReturnReceipt

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubsidyReturnReceiptId** | **string** | 补差回退唯一单号，由微信支付生成，仅在补差回退成功后有返回 | [可选] 
**StockId** | **string** | 由微信支付生成，调用创建商家券API成功时返回的唯一批次ID | 
**CouponCode** | **string** | 券的唯一标识 | 
**TransactionId** | **string** | 微信支付下单支付成功返回的订单号 | 
**RefundId** | **string** | 微信支付退款单号 | 
**PayerMerchant** | **string** | 原营销补差扣款商户号，即回退资金收款商户号 | 
**PayeeMerchant** | **string** | 原营销补差入账商户号，即回退资金扣款商户号 | 
**Amount** | **int64** | 本次补差回退金额，单位为分。单个券code回退总金额不能超过补差金额 | 
**Description** | **string** | 回退备注描述，查询的时候原样带回 | 
**Status** | [**SubsidyReturnReceiptStatus**](SubsidyReturnReceiptStatus.md) | 补差付款单据状态 | 
**FailReason** | [**SubsidyReturnReceiptFailReason**](SubsidyReturnReceiptFailReason.md) | 仅在补差回退失败时，返回告知对应失败的原因 | [可选] 
**ReturnDoneTime** | **time.Time** | 仅在补差回退成功时，返回完成时间 | [可选] 
**SubsidyReceiptId** | **string** | 此次补差回退操作对应的补差付款单号 | 
**OutSubsidyReturnNo** | **string** | 商户侧需保证唯一性。可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号 | 
**ReturnCreateTime** | **time.Time** | 补差回退单据创建时间 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


