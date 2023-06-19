# ActBaseInfo

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**ActivityName** | **string** | 活动名称 | 
**ActivitySecondTitle** | **string** | 活动副标题 | 
**MerchantLogoUrl** | **string** | 商户logo，送出优惠券时展示 | 
**BackgroundColor** | **string** | 代金券的背景颜色，可设置10种颜色，颜色取值可参考[开发指引](../../../products/gift-activity/development.html) 3.2.1  background_color取值，默认为微信支付绿色， 颜色取值为颜色图中的颜色名称。 | [可选] 
**BeginTime** | **string** | 活动开始时间，最长可以配置1年内的活动，活动有效期最长90天 | 
**EndTime** | **string** | 活动结束时间，最长可以配置1年内的活动，活动有效期最长90天 | 
**AvailablePeriods** | [**AvailablePeriod**](AvailablePeriod.md) | 可自定义活动有效时间内可用的多个时间段以及天内时间点 | [可选] 
**OutRequestNo** | **string** | 商户创建批次凭据号（格式：商户id+日期+流水号），商户侧需保持唯一性，可包含英文字母，数字，｜，_，*，-等内容，不允许出现其他不合法符号 | 
**DeliveryPurpose** | [**DeliveryPurposeCategory**](DeliveryPurposeCategory.md) | OFF_LINE_PAY：拉用户回店消费；JUMP_MINI_APP：引导用户前往小程序消费 | 
**MiniProgramsAppid** | **string** | 投放目的为跳转小程序时必填 | [可选] 
**MiniProgramsPath** | **string** | 投放目的为跳转小程序必填 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


