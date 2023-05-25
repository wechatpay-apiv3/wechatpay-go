# DisplayPatternInfo

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Description** | **string** | 用于说明详细的活动规则，会展示在代金券详情页。 | [可选] 
**MerchantLogoUrl** | **string** | 若券归属商户号有认证品牌，则系统将自动拉取对应品牌logo；若券归属商户号不在认证品牌下，需自定义上传logo，未上传时将展示兜底灰色logo样式，影响券详情页用户体验，请及时上传。 商户logo的URL地址，仅支持通过《[图片上传API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_0_1.shtml)》接口获取的图片URL地址。 1、商户logo大小需为120像素\\*120像素。 2、支持JPG/JPEG/PNG格式，且图片小于1M。 注：该字段暂不支持修改 | [可选] 
**MerchantName** | **string** | 不支持商户自定义。若券归属商户号有认证品牌，系统将自动拉取认证品牌号下的品牌名称；若券归属商户号不在认证品牌下，则拉取本商户号的商户简称。展示上限12个字符。 注：该字段暂不支持修改 | [可选] 
**BackgroundColor** | **string** | 券的背景颜色，可设置10种颜色，色值请参考[卡券背景颜色图](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_2_1.shtml#part-01)。颜色取值为颜色图中的颜色名称。 | [可选] 
**CouponImageUrl** | **string** | 券详情图片，1074像素（宽）\\*603像素（高），图片大小不超过2M，支持JPG/PNG格式。仅支持通过《[图片上传API](https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_0_1.shtml)》接口获取的图片URL地址。 | [可选] 
**FinderInfo** | [**FinderInfo**](FinderInfo.md) | 视频号相关信息 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


