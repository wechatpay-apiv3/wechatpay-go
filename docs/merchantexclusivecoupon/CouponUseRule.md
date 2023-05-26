# CouponUseRule

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**CouponAvailableTime** | [**FavorAvailableTime**](FavorAvailableTime.md) | 日期区间内可以使用优惠 | 
**FixedNormalCoupon** | [**FixedValueStockMsg**](FixedValueStockMsg.md) | 固定面额满减，折扣券，换购券使用规则三选一，stock_type为NORMAL时必填。 | [可选] 
**DiscountCoupon** | [**DiscountMsg**](DiscountMsg.md) | 固定面额满减，折扣券，换购券使用规则三选一，stock_type为DISCOUNT时必填。 | [可选] 
**ExchangeCoupon** | [**ExchangeMsg**](ExchangeMsg.md) | 固定面额满减，折扣券，换购券使用规则三选一，stock_type为EXCHANGE时必填。 | [可选] 
**UseMethod** | [**CouponUseMethod**](CouponUseMethod.md) | 核销方式 | 
**MiniProgramsAppid** | **string** | 核销方式为线上小程序核销才有效 | [可选] 
**MiniProgramsPath** | **string** | 核销方式为线上小程序核销才有效 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


