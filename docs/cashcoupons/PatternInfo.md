# PatternInfo

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Description** | **string** | 用于说明详细的活动规则，会展示在代金券详情页 | 
**MerchantLogo** | **string** | 商户logo | [可选] 
**MerchantName** | **string** | 商户名称 | [可选] 
**BackgroundColor** | [**BackgroundColor**](BackgroundColor.md) | 背景颜色 | [可选] 
**CouponImage** | **string** | 券详情图片 | [可选] 
**JumpTarget** | [**JumpTarget**](JumpTarget.md) | 枚举值：PAYMENT_CODE：跳转至微信支付付款码，点击“立即使用”跳转至微信支付付款码 MINI_PROGRAM：跳转至小程序，点击“立即使用”跳转至配置的商家小程序（需要指定小程序appid和path） DEFAULT_PAGE：跳转至默认页， 点击“立即使用”跳转至默认页面  如未传该参数，则默认跳转至默认页。 | [可选] 
**MiniProgramAppid** | **string** | 跳转的小程序appid，跳转至小程序时必填。跳转的小程序appid需至少和一个可核销商户有绑定关系。 | [可选] 
**MiniProgramPath** | **string** | 跳转的小程序path，跳转至小程序时必填。 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


