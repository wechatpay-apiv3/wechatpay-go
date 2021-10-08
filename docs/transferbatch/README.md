# 微信支付 API v3 Go SDK - transferbatch

- 场景及业务流程：
商户可通过该产品实现同时向多个用户微信零钱进行转账的操作，可用于发放奖金补贴、佣金货款结算、员工报销等场景。
[https://pay.weixin.qq.com/index.php/public/product/detail?pid=108&productType=0](https://pay.weixin.qq.com/index.php/public/product/detail?pid=108&productType=0)
- 接入步骤：
    - 商户在微信支付商户平台开通“批量转账到零钱”产品权限，并勾选“使用API方式发起转账”。
    - 调用批量转账接口，对多个用户微信零钱发起转账。
    - 调用查询批次接口，可获取到转账批次详情及当前状态。
    - 调用查询明细接口，可获取到单条转账明细详情及当前状态。


## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.0.0

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*TransferBatchApi* | [**GetTransferBatchByNo**](TransferBatchApi.md#gettransferbatchbyno) | **Get** /v3/transfer/batches/batch-id/{batch_id} | 通过微信批次单号查询批次单
*TransferBatchApi* | [**GetTransferBatchByOutNo**](TransferBatchApi.md#gettransferbatchbyoutno) | **Get** /v3/transfer/batches/out-batch-no/{out_batch_no} | 通过商家批次单号查询批次单
*TransferBatchApi* | [**InitiateBatchTransfer**](TransferBatchApi.md#initiatebatchtransfer) | **Post** /v3/transfer/batches | 发起批量转账
*TransferDetailApi* | [**GetTransferDetailByNo**](TransferDetailApi.md#gettransferdetailbyno) | **Get** /v3/transfer/batches/batch-id/{batch_id}/details/detail-id/{detail_id} | 通过微信明细单号查询明细单
*TransferDetailApi* | [**GetTransferDetailByOutNo**](TransferDetailApi.md#gettransferdetailbyoutno) | **Get** /v3/transfer/batches/out-batch-no/{out_batch_no}/details/out-detail-no/{out_detail_no} | 通过商家明细单号查询明细单


## 类型列表

 - [CloseReasonType](CloseReasonType.md)
 - [FailReasonType](FailReasonType.md)
 - [GetTransferBatchByNoRequest](GetTransferBatchByNoRequest.md)
 - [GetTransferBatchByOutNoRequest](GetTransferBatchByOutNoRequest.md)
 - [GetTransferDetailByNoRequest](GetTransferDetailByNoRequest.md)
 - [GetTransferDetailByOutNoRequest](GetTransferDetailByOutNoRequest.md)
 - [InitiateBatchTransferRequest](InitiateBatchTransferRequest.md)
 - [InitiateBatchTransferResponse](InitiateBatchTransferResponse.md)
 - [TransferBatchEntity](TransferBatchEntity.md)
 - [TransferBatchGet](TransferBatchGet.md)
 - [TransferDetailCompact](TransferDetailCompact.md)
 - [TransferDetailEntity](TransferDetailEntity.md)
 - [TransferDetailInput](TransferDetailInput.md)

