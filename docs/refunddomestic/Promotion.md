# Promotion

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**PromotionId** | **string** | 券或者立减优惠id  | 
**Scope** | [**Scope**](Scope.md) | 枚举值： - GLOBAL- 全场代金券 - SINGLE- 单品优惠 * &#x60;GLOBAL&#x60; - 全场代金券 * &#x60;SINGLE&#x60; - 单品优惠  | 
**Type** | [**Type**](Type.md) | 枚举值： - COUPON- 代金券，需要走结算资金的充值型代金券 - DISCOUNT- 优惠券，不走结算资金的免充值型优惠券 * &#x60;COUPON&#x60; - 代金券 * &#x60;DISCOUNT&#x60; - 优惠券  | 
**Amount** | **int64** | 用户享受优惠的金额（优惠券面额&#x3D;微信出资金额+商家出资金额+其他出资方金额 ），单位为分  | 
**RefundAmount** | **int64** | 优惠退款金额&lt;&#x3D;退款金额，退款金额-代金券或立减优惠退款金额为用户支付的现金，说明详见代金券或立减优惠，单位为分  | 
**GoodsDetail** | [**[]GoodsDetail**](GoodsDetail.md) | 优惠商品发生退款时返回商品信息  | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


