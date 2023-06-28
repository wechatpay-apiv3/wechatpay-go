# StockUseRule

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**MaxCoupons** | **int64** | 最大发券数 | 
**MaxAmount** | **int64** | 总消耗金额，单位分 | 
**MaxAmountByDay** | **int64** | 单天最高消耗金额，单位分 | 
**FixedNormalCoupon** | [**FixedValueStockMsg**](FixedValueStockMsg.md) | 固定面额发券批次特定信息 | [可选] 
**MaxCouponsPerUser** | **int64** | 单个用户可领个数 | 
**CouponType** | **string** | NORMAL-满减券；CUT_TO-减至券 | [可选] 
**GoodsTag** | **[]string** | 订单优惠标记 | [可选] 
**TradeType** | [**[]TradeType**](TradeType.md) | 默认不限制，可设置以下各种组合方式 1、付款码支付 2、小程序支付 3、其他支付 4、付款码支付&amp;小程序支付 5、付款码支付&amp;其他支付 | [可选] 
**CombineUse** | **bool** | true-是；false-否 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


