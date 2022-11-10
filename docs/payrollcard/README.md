# 微信支付 API v3 Go SDK - payrollcard

服务商通过本API文档提供的接口，查询商户和微工卡的授权关系、生成预授权的token口令、核身预下单、核身结果的查询等。

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.5.2

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*AuthenticationsApi* | [**GetAuthentication**](AuthenticationsApi.md#getauthentication) | **Get** /v3/payroll-card/authentications/{authenticate_number} | 获取核身结果
*AuthenticationsApi* | [**ListAuthentications**](AuthenticationsApi.md#listauthentications) | **Get** /v3/payroll-card/authentications | 查询核身记录
*AuthenticationsApi* | [**PreOrderAuthentication**](AuthenticationsApi.md#preorderauthentication) | **Post** /v3/payroll-card/authentications/pre-order | 微工卡核身预下单
*AuthenticationsApi* | [**PreOrderAuthenticationWithAuth**](AuthenticationsApi.md#preorderauthenticationwithauth) | **Post** /v3/payroll-card/authentications/pre-order-with-auth | 微工卡核身预下单（流程中完成授权）
*RelationsApi* | [**GetRelation**](RelationsApi.md#getrelation) | **Get** /v3/payroll-card/relations/{openid} | 查询微工卡授权关系
*TokensApi* | [**CreateToken**](TokensApi.md#createtoken) | **Post** /v3/payroll-card/tokens | 生成授权token
*TransferBatchApi* | [**CreateTransferBatch**](TransferBatchApi.md#createtransferbatch) | **Post** /v3/payroll-card/transfer-batches | 发起批量转账


## 类型列表

 - [AuthType](AuthType.md)
 - [AuthenticationEntity](AuthenticationEntity.md)
 - [AuthenticationScene](AuthenticationScene.md)
 - [AuthenticationState](AuthenticationState.md)
 - [AuthenticationType](AuthenticationType.md)
 - [BusinessType](BusinessType.md)
 - [CreateTokenRequest](CreateTokenRequest.md)
 - [CreateTransferBatchRequest](CreateTransferBatchRequest.md)
 - [EmploymentScene](EmploymentScene.md)
 - [EmploymentType](EmploymentType.md)
 - [GetAuthenticationRequest](GetAuthenticationRequest.md)
 - [GetRelationRequest](GetRelationRequest.md)
 - [ListAuthenticationsRequest](ListAuthenticationsRequest.md)
 - [ListAuthenticationsResponse](ListAuthenticationsResponse.md)
 - [PayrollCardRegisterState](PayrollCardRegisterState.md)
 - [PreOrderAuthenticationRequest](PreOrderAuthenticationRequest.md)
 - [PreOrderAuthenticationResponse](PreOrderAuthenticationResponse.md)
 - [PreOrderAuthenticationWithAuthRequest](PreOrderAuthenticationWithAuthRequest.md)
 - [PreOrderAuthenticationWithAuthResponse](PreOrderAuthenticationWithAuthResponse.md)
 - [RelationEntity](RelationEntity.md)
 - [TokenEntity](TokenEntity.md)
 - [TransferBatchEntity](TransferBatchEntity.md)
 - [TransferDetailInput](TransferDetailInput.md)

