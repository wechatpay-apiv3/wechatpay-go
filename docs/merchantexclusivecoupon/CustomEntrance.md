# CustomEntrance

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**MiniProgramsInfo** | [**MiniAppInfo**](MiniAppInfo.md) | 需要小程序APPID、path、入口文案、引导文案。如果需要跳转小程序，APPID、path、入口文案为必填，引导文案非必填。 AppID要与归属商户号有M-A or M-m-suba关系。 注：请查看[绑定关系说明文档](https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/chapter5_2_2.shtml#part-7) | [可选] 
**Appid** | **string** | 可配置商户公众号，从券详情可跳转至公众号，用户自定义字段。 校验规则：传入的AppID得是与调用方商户号（即请求头里面的商户号）有绑定关系的AppID 或 传入的AppID得是归属商户号有绑定关系的AppID | [可选] 
**HallId** | **string** | 填写微信支付营销馆的馆ID，用户自定义字段。[营销馆](https://pay.weixin.qq.com/index.php/xphp/cfav_market/hall#/pages/list/list)需在商户平台创建。 | [可选] 
**StoreId** | **string** | 填写代金券可用门店ID | [可选] 
**CodeDisplayMode** | [**CodeDisplayMode**](CodeDisplayMode.md) | code展示模式 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


