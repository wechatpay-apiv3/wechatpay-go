# 微信支付 API v3 Go SDK - retailstore

指定服务商可通过该接口报名加价购活动、查询某个区域内的加价购活动列表、锁定加价活动购资格以及解锁加价购活动资格。

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.4.0

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*ActivityApi* | [**ApplyActivity**](ActivityApi.md#applyactivity) | **Post** /v3/marketing/goods-subsidy-activity/activity/{activity_id}/apply | 门店报名品牌加价购活动
*ActivityApi* | [**ListActsByArea**](ActivityApi.md#listactsbyarea) | **Get** /v3/marketing/goods-subsidy-activity/activities | 按区域查询品牌加价购活动
*QualificationApi* | [**LockQualification**](QualificationApi.md#lockqualification) | **Post** /v3/marketing/goods-subsidy-activity/qualification/lock | 锁定品牌加价购活动资格
*QualificationApi* | [**UnlockQualification**](QualificationApi.md#unlockqualification) | **Post** /v3/marketing/goods-subsidy-activity/qualification/unlock | 解锁品牌加价购活动资格
*RetailStoreActApi* | [**AddRepresentative**](RetailStoreActApi.md#addrepresentative) | **Put** /v3/marketing/goods-subsidy-activity/retail-store-act/{activity_id}/representative | 添加零售小店活动业务代理
*RetailStoreActApi* | [**AddStores**](RetailStoreActApi.md#addstores) | **Post** /v3/marketing/goods-subsidy-activity/retail-store-act/{brand_id}/stores | 添加小店活动门店
*RetailStoreActApi* | [**CreateMaterials**](RetailStoreActApi.md#creatematerials) | **Post** /v3/marketing/goods-subsidy-activity/retail-store-act/{brand_id}/materials | 生成小店活动物料码
*RetailStoreActApi* | [**DeleteRepresentative**](RetailStoreActApi.md#deleterepresentative) | **Delete** /v3/marketing/goods-subsidy-activity/retail-store-act/{activity_id}/representative | 删除零售小店活动业务代理
*RetailStoreActApi* | [**DeleteStores**](RetailStoreActApi.md#deletestores) | **Delete** /v3/marketing/goods-subsidy-activity/retail-store-act/{brand_id}/stores | 删除小店活动门店
*RetailStoreActApi* | [**GetStore**](RetailStoreActApi.md#getstore) | **Get** /v3/marketing/goods-subsidy-activity/retail-store-act/{brand_id}/stores/{store_code} | 查询小店活动门店详情
*RetailStoreActApi* | [**ListRepresentative**](RetailStoreActApi.md#listrepresentative) | **Get** /v3/marketing/goods-subsidy-activity/retail-store-act/{activity_id}/representatives | 查询零售小店活动业务代理
*RetailStoreActApi* | [**ListStore**](RetailStoreActApi.md#liststore) | **Get** /v3/marketing/goods-subsidy-activity/retail-store-act/{brand_id}/stores | 查询小店活动门店列表


## 类型列表

 - [ActApplyInfo](ActApplyInfo.md)
 - [ActArea](ActArea.md)
 - [ActInfo](ActInfo.md)
 - [ActRule](ActRule.md)
 - [AddRepresentativeRequest](AddRepresentativeRequest.md)
 - [AddRepresentativesRequest](AddRepresentativesRequest.md)
 - [AddRepresentativesResponse](AddRepresentativesResponse.md)
 - [AddStoresBody](AddStoresBody.md)
 - [AddStoresRequest](AddStoresRequest.md)
 - [AddStoresResponse](AddStoresResponse.md)
 - [ApplyActivityBody](ApplyActivityBody.md)
 - [ApplyActivityRequest](ApplyActivityRequest.md)
 - [ApplyActivityResponse](ApplyActivityResponse.md)
 - [CreateMaterialsBody](CreateMaterialsBody.md)
 - [CreateMaterialsRequest](CreateMaterialsRequest.md)
 - [DeleteRepresentativeBody](DeleteRepresentativeBody.md)
 - [DeleteRepresentativeRequest](DeleteRepresentativeRequest.md)
 - [DeleteRepresentativeResponse](DeleteRepresentativeResponse.md)
 - [DeleteStoresBody](DeleteStoresBody.md)
 - [DeleteStoresRequest](DeleteStoresRequest.md)
 - [DeleteStoresResponse](DeleteStoresResponse.md)
 - [FmcgAct](FmcgAct.md)
 - [GetStoreRequest](GetStoreRequest.md)
 - [GoodsInfo](GoodsInfo.md)
 - [ListActsByAreaRequest](ListActsByAreaRequest.md)
 - [ListActsByAreaResponse](ListActsByAreaResponse.md)
 - [ListRepresentativeRequest](ListRepresentativeRequest.md)
 - [ListRepresentativeResponse](ListRepresentativeResponse.md)
 - [ListStoreRequest](ListStoreRequest.md)
 - [ListStoreResponse](ListStoreResponse.md)
 - [LockQualificationRequest](LockQualificationRequest.md)
 - [LockQualificationResponse](LockQualificationResponse.md)
 - [MaterialInfo](MaterialInfo.md)
 - [Materials](Materials.md)
 - [OrderInfo](OrderInfo.md)
 - [RepresentativeInfo](RepresentativeInfo.md)
 - [RetailStoreInfo](RetailStoreInfo.md)
 - [StoreInfo](StoreInfo.md)
 - [UnlockQualificationRequest](UnlockQualificationRequest.md)
 - [UnlockQualificationResponse](UnlockQualificationResponse.md)

