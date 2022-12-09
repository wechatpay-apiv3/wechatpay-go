# 微信支付 API v3 Go SDK - lovefeast

微信支付爱心餐公益计划旨在面向深圳市的市政一线环卫工人提供每周一餐的1分钱用餐公益服务。在受助端，微信支付联动上千家餐饮门店关爱特殊人群，通过微信支付数字化能力将人群身份认证与公益福利领用全流程线上化，实现公益福利精准到人。在捐赠端，微信支付发挥连接优势与平台能力，结合用户就餐场景通过爱心餐一块捐插件让用户可在点餐时顺手捐1元，带动更多社会力量致谢城市美容师。

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 0.0.4

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*BrandsApi* | [**GetBrand**](BrandsApi.md#getbrand) | **Get** /v3/lovefeast/brands/{brand_id} | 查询爱心餐品牌信息
*OrdersApi* | [**GetByUser**](OrdersApi.md#getbyuser) | **Get** /v3/lovefeast/users/{openid}/orders/out-trade-no/{out_trade_no} | 查询用户捐赠单详情
*OrdersApi* | [**ListByUser**](OrdersApi.md#listbyuser) | **Get** /v3/lovefeast/users/{openid}/orders/brand-id/{brand_id} | 查询用户捐赠单列表


## 类型列表

 - [Amount](Amount.md)
 - [BrandEntity](BrandEntity.md)
 - [GetBrandRequest](GetBrandRequest.md)
 - [GetByUserRequest](GetByUserRequest.md)
 - [ListByUserRequest](ListByUserRequest.md)
 - [MerchantOrder](MerchantOrder.md)
 - [OrdersEntity](OrdersEntity.md)
 - [OrdersListByUserResponse](OrdersListByUserResponse.md)
 - [Payer](Payer.md)

