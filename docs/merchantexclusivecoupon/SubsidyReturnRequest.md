# SubsidyReturnRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 由微信支付生成，调用创建商家券API成功时返回的唯一批次ID 仅支持“满减券”和“折扣券”的批次，“换购券”批次不支持 | 
**CouponCode** | **string** | 券的唯一标识。 在WECHATPAY\\_MODE的券code模式下，商家券code是由微信支付生成的唯一ID； 在MERCHANT\\_UPLOAD、MERCHANT\\_API的券code模式下，商家券code是由商户上传或指定，在批次下保证唯一； | 
**TransactionId** | **string** | 微信支付下单支付成功返回的订单号 | 
**RefundId** | **string** | 微信支付退款单号 | 
**PayerMerchant** | **string** | 原营销补差扣款商户号，即回退资金收款商户号 | 
**PayeeMerchant** | **string** | 原营销补差入账商户号，即回退资金扣款商户号 | 
**Amount** | **int64** | 本次补差回退金额，单位为分。单个券Code回退总金额不能超过补差金额 | 
**Description** | **string** | 回退备注描述，查询的时候原样带回 | 
**OutSubsidyReturnNo** | **string** | 商户侧需保证唯一性，可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


