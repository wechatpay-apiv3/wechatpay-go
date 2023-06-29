# CouponRule

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**CouponAvailableTime** | [**FavorAvailableTime**](FavorAvailableTime.md) | 需要指定领取后延时生效可选填 | [可选] 
**FixedNormalCoupon** | [**FixedValueStockMsg**](FixedValueStockMsg.md) | stock_type为NORMAL时必填 | [可选] 
**GoodsTag** | **[]string** | 订单优惠标记 | [可选] 
**TradeType** | [**[]TradeType**](TradeType.md) | 支付方式 | [可选] 
**CombineUse** | **bool** | true-是；false-否 | [可选] 
**AvailableItems** | **[]string** | 可核销商品编码 | [可选] 
**UnavailableItems** | **[]string** | 不参与优惠商品编码 | [可选] 
**AvailableMerchants** | **[]string** | 可核销商户号 | [可选] 
**LimitCard** | [**CardLimitation**](CardLimitation.md) | 当批次指定支付方式为某张银行卡时才生效，可选的 | [可选] 
**LimitPay** | **[]string** | 限定该批次的指定支付方式，如零钱、指定银行卡等，需填入支付方式编码， 条目个数限制为[1，1] 。当前支持的支付方式，及其编码枚举值，请参考该文档： https://docs.qq.com/sheet/DWGpMbWx3b1JCbldy?c&#x3D;E3A0A0  | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


