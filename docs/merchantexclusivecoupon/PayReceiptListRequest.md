# PayReceiptListRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 由微信支付生成，调用创建商家券API成功时返回的唯一批次ID | 
**CouponCode** | **string** | 券的唯一标识。 在WECHATPAY\\_MODE的券code模式下，商家券code是由微信支付生成的唯一ID； 在MERCHANT\\_UPLOAD、MERCHANT\\_API的券code模式下，商家券code是由商户上传或指定，在批次下保证唯一； | 
**OutSubsidyNo** | **string** | 商户调用补差付款API时填写的“业务请求唯一单号” | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


