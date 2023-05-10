# SetAdvertisingIndustryFilterRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**SubMchid** | **string** | 需要设置“同业过滤标签”的特约商户号，由微信支付生成并下发。 | 
**AdvertisingIndustryFilters** | [**[]IndustryType**](IndustryType.md) | 特约商户同业过滤的同业过滤标签值，同业过滤标签最少传一个，最多三个。如已设置同业过滤标签，再次请求传入，视为新增，将覆盖原有同业标签配置。 注：若配置完成后需清空标签的，可登陆商户平台手动清空（路径：商户平台-&gt;服务商功能-&gt;点金计划-&gt;特约商户管理-&gt;同业过滤标签） | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


