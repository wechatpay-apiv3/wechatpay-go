# CouponEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**BelongMerchant** | **string** | 代金券的所属商户号 | 
**StockName** | **string** | 批次名称，字数上限为21个，一个中文汉字/英文字母/数字均占用一个字数。 | 
**Comment** | **string** | 仅配置商户可见，用于自定义信息。字数上限为20个，一个中文汉字/英文字母/数字均占用一个字数。 | [可选] 
**GoodsName** | **string** | 适用商品范围，字数上限为15个，一个中文汉字/英文字母/数字均占用一个字数。 | 
**StockType** | [**BusiFavorStockType**](BusiFavorStockType.md) | 批次类型 | 
**Transferable** | **bool** | 不填默认否，枚举值： true：是 false：否 该字段暂未开放 | [可选] 
**Shareable** | **bool** | 不填默认否，枚举值： true：是 false：否 该字段暂未开放 | [可选] 
**CouponState** | [**CouponStatus**](CouponStatus.md) | 商家券状态 | [可选] 
**DisplayPatternInfo** | [**DisplayPatternInfo**](DisplayPatternInfo.md) |  | [可选] 
**CouponUseRule** | [**CouponUseRule**](CouponUseRule.md) |  | 
**CustomEntrance** | [**CustomEntrance**](CustomEntrance.md) |  | [可选] 
**CouponCode** | **string** | 券的唯一标识 | [可选] 
**StockId** | **string** | 批次号 | [可选] 
**AvailableStartTime** | **time.Time** | 用户领取到的这张券实际可使用的开始时间：如批次设置的领取后可用，则开始时间即为券的领取时间； 如批次设置的领取后第X天可用，则为领取时间后第X天00:00:00 | [可选] 
**ExpireTime** | **time.Time** | 用户领取到这张券的过期时间 | [可选] 
**ReceiveTime** | **time.Time** | 用户领取到这张券的时间 | [可选] 
**SendRequestNo** | **string** | 发券时传入的唯一凭证 | [可选] 
**UseRequestNo** | **string** | 核销时传入的唯一凭证（如券已被核销，将返回此字段） | [可选] 
**UseTime** | **time.Time** | 券被核销的时间（如券已被核销，将返回此字段） | [可选] 
**AssociateOutTradeNo** | **time.Time** | 若商家券操作过关联商户订单信息，则该字段返回商家券已关联的商户订单号。 | [可选] 
**ReturnRequestNo** | **string** | 回退时传入的唯一凭证（如券发生了退回，将返回此字段） | [可选] 
**ReturnTime** | **time.Time** | 券被回退的时间（如券发生了退回，将返回此字段） | [可选] 
**DeactivateRequestNo** | **string** | 失效时传入的唯一凭证（如果一张券已失效，将返回此字段） | [可选] 
**DeactivateTime** | **time.Time** | 券被失效的时间（如果一张券已失效，将返回此字段） | [可选] 
**DeactivateReason** | **string** | 失效一张券的原因（如果一张券已失效，可能返回此字段） | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


