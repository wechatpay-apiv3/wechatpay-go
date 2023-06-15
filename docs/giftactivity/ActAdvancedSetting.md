# ActAdvancedSetting

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**DeliveryUserCategory** | [**DeliveryUserCategory**](DeliveryUserCategory.md) | DELIVERY_ALL_PERSON:所有用户；DELIVERY_MEMBER_PERSON：会员用户 | [可选] 
**MerchantMemberAppid** | **string** | 当投放用户类别为会员用户时必填 | [可选] 
**PaymentMode** | [**PaymentMode**](PaymentMode.md) | 主要用于指定可用的支付场景，不限制支付模式不填，限制选填（商家券暂不开放） | [可选] 
**PaymentMethodInformation** | [**PaymentMethodInfo**](PaymentMethodInfo.md) | 支付方式信息，不填为不限制（商家券暂不开放） | [可选] 
**GoodsTags** | **[]string** | 商户下单时需要传入相同的标记(goods_tag)，用户同时符合其他规则才能享受优惠 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


