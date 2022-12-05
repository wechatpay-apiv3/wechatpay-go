# OrdersEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**WelfareTradeId** | **string** | 捐赠订单号 | 
**Appid** | **string** | 直连商户或服务商申请的公众号或移动应用AppID | 
**SubAppid** | **string** | 子商户申请的公众号或移动应用AppID | [可选] 
**BrandId** | **int64** | 品牌的品牌ID，商家进驻微信支付品牌商家后获得的品牌ID | 
**DonateSource** | **string** | 捐赠来源 枚举值 MINIPROGRAM_PAY：一起捐小程序 ENTRUST_PAY：一块捐插件 | 
**MerchantOrder** | [**MerchantOrder**](MerchantOrder.md) | 商户订单，仅在一块捐来源回调中出现 | [可选] 
**InstitutionName** | **string** | 捐赠公益机构名称。 | 
**BusinessId** | **string** | 本次捐赠项目编号。 | 
**BusinessName** | **string** | 本次捐赠活动名称。 | 
**SuccessTime** | **string** | 支付完成时间，遵循rfc3339标准格式，格式为yyyy-MM-DDTHH:mm:ss+TIMEZONE，yyyy-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日 13点29分35秒。 | 
**Payer** | [**Payer**](Payer.md) | 支付者信息 | 
**Amount** | [**Amount**](Amount.md) | 订单金额信息 | 
**DeviceId** | **string** | 设备编号，通过物料码进入爱心餐一起捐小程序捐赠时携带的设备编号参数 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


