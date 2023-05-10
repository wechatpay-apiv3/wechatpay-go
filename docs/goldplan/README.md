# 微信支付 API v3 Go SDK - goldplan

特约商户点金计划管理API

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 0.2.3

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*MerchantsApi* | [**CloseAdvertisingShow**](MerchantsApi.md#closeadvertisingshow) | **Post** /v3/goldplan/merchants/close-advertising-show | 关闭广告展示
*MerchantsApi* | [**OpenAdvertisingShow**](MerchantsApi.md#openadvertisingshow) | **Patch** /v3/goldplan/merchants/open-advertising-show | 开通广告展示
*MerchantsApi* | [**SetAdvertisingIndustryFilter**](MerchantsApi.md#setadvertisingindustryfilter) | **Post** /v3/goldplan/merchants/set-advertising-industry-filter | 同业过滤标签管理
*StatusApi* | [**ChangeCustomPageStatus**](StatusApi.md#changecustompagestatus) | **Post** /v3/goldplan/merchants/changecustompagestatus | 商家小票管理
*StatusApi* | [**ChangeGoldPlanStatus**](StatusApi.md#changegoldplanstatus) | **Post** /v3/goldplan/merchants/changegoldplanstatus | 点金计划管理


## 类型列表

 - [ChangeCustomPageStatusRequest](ChangeCustomPageStatusRequest.md)
 - [ChangeCustomPageStatusResponse](ChangeCustomPageStatusResponse.md)
 - [ChangeGoldPlanStatusRequest](ChangeGoldPlanStatusRequest.md)
 - [ChangeGoldPlanStatusResponse](ChangeGoldPlanStatusResponse.md)
 - [CloseAdvertisingShowRequest](CloseAdvertisingShowRequest.md)
 - [IndustryType](IndustryType.md)
 - [OpenAdvertisingShowRequest](OpenAdvertisingShowRequest.md)
 - [OperationType](OperationType.md)
 - [SetAdvertisingIndustryFilterRequest](SetAdvertisingIndustryFilterRequest.md)

