# SendCouponRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**StockId** | **string** | 微信为每个批次分配的唯一id | 
**Openid** | **string** | Openid信息，用户在appid下的唯一标识 | 
**OutRequestNo** | **string** | 商户此次发放凭据号（格式：商户id+日期+流水号），商户侧需保持唯一性 | 
**Appid** | **string** | 微信为发券方商户分配的公众账号ID，接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。 | 
**StockCreatorMchid** | **string** | 批次创建方商户号 | 
**CouponValue** | **int64** | 指定面额发券场景，券面额，其他场景不需要填，单位分 | [可选] 
**CouponMinimum** | **int64** | 指定面额发券批次门槛，其他场景不需要，单位分 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


