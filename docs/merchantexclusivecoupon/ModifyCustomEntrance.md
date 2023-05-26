# ModifyCustomEntrance

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**MiniProgramsInfo** | [**ModifyMiniAppInfo**](ModifyMiniAppInfo.md) | 需要小程序APPID、path、入口文案、引导文案。如果需要跳转小程序，APPID、path、入口文案为必填，引导文案非必填。AppID要与归属商户号有绑定关系 | [可选] 
**Appid** | **string** | 可配置商户公众号，从券详情可跳转至公众号 | [可选] 
**HallId** | **string** | 填写微信支付营销馆的馆id，用户自定义字段。[营销馆](https://pay.weixin.qq.com/index.php/xphp/cfav_market/hall#/pages/list/list)需在商户平台创建。 | [可选] 
**CodeDisplayMode** | [**CodeDisplayMode**](CodeDisplayMode.md) | code展示模式 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


