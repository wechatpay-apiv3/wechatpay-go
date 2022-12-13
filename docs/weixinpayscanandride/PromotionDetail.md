# PromotionDetail

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**CouponId** | **string** | 券或者立减优惠ID | 
**Name** | **string** | 优惠名称 | [可选] 
**Scope** | **string** | GLOBAL-全场代金券，SINGLE-单品优惠 | [可选] 
**Type** | [**PromotionType**](PromotionType.md) | 优惠类型，枚举值有： | [可选] 
**StockId** | **string** | 在微信商户后台配置的批次ID | [可选] 
**Amount** | **int64** | 用户享受优惠的金额 | 
**WechatpayContribute** | **int64** | 特指由微信支付商户平台创建的优惠，出资金额等于本项优惠总金额，单位为分 | [可选] 
**MerchantContribute** | **int64** | 特指商户自己创建的优惠，出资金额等于本项优惠总金额，单位为分 | [可选] 
**OtherContribute** | **int64** | 其他出资方出资金额，单位为分 | [可选] 
**Currency** | **string** | CNY：人民币，境内商户号仅支持人民币。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


