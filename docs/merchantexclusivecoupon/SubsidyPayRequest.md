# SubsidyPayRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 由微信支付生成，调用创建商家券API成功时返回的唯一批次ID 仅支持“满减券”和“折扣券”的批次，“换购券”批次不支持 | 
**CouponCode** | **string** | 券的唯一标识。 在WECHATPAY\\_MODE的券code模式下，商家券code是由微信支付生成的唯一ID； 在MERCHANT\\_UPLOAD、MERCHANT\\_API的券code模式下，商家券code是由商户上传或指定，在批次下保证唯一； | 
**TransactionId** | **string** | 微信支付下单支付成功返回的订单号 | 
**PayerMerchant** | **string** | 营销补差扣款商户号 注：补差扣款商户号 &#x3D; 制券商户号 或 补差扣款商户号 &#x3D; 归属商户号 | 
**PayeeMerchant** | **string** | 营销补差入账商户号 注：补差入账商户号 &#x3D; 券归属商户号 或者和 券归属商户号有连锁品牌关系 | 
**Amount** | **int64** | 单位为分，单笔订单补差金额不得超过券的优惠金额，最高补差金额为5000元 &gt; 券的优惠金额定义：   满减券：满减金额即为优惠金额   折扣券：优惠金额 &#x3D; 微信支付订单金额 ÷ 折扣比例 × (1 - 折扣比例)   换购券：不支持 | 
**Description** | **string** | 付款备注描述，查询的时候原样带回 | 
**OutSubsidyNo** | **string** | 商户侧需保证唯一性，可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号 | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


