# 微信支付 API v3 Go SDK - transferbills

* 场景及业务流程：
    商户可通过该产品实现向单个用户微信零钱进行转账的操作，可用于红包奖励、佣金结算、退款补偿等场景。
    [https://pay.weixin.qq.com/doc/v3/merchant/4012716434](https://pay.weixin.qq.com/doc/v3/merchant/4012716434)
* 接入步骤：
    * 商户在微信支付商户平台开通"转账到零钱"产品权限，并配置相关参数。
    * 调用转账接口，向指定用户微信零钱发起转账。
    * 调用查询接口，可获取到转账详情及当前状态。
    * 可调用撤销接口，在用户确认前撤销转账。

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 1.0.0

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*TransferBillsApi* | [**CreateTransferBill**](TransferBillsApi.md#createtransferbill) | **Post** /v3/fund-app/mch-transfer/transfer-bills | 发起商家转账到用户零钱
*TransferBillsApi* | [**GetTransferBillByOutBillNo**](TransferBillsApi.md#gettransferbillbyoutbillno) | **Get** /v3/fund-app/mch-transfer/transfer-bills/out-bill-no/{out_bill_no} | 通过商户订单号查询转账单
*TransferBillsApi* | [**GetTransferBillByTransferBillNo**](TransferBillsApi.md#gettransferbillbytransferbillno) | **Get** /v3/fund-app/mch-transfer/transfer-bills/transfer-bill-no/{transfer_bill_no} | 通过微信转账单号查询转账单
*TransferBillsApi* | [**CancelTransferBill**](TransferBillsApi.md#canceltransferbill) | **Post** /v3/fund-app/mch-transfer/transfer-bills/out-bill-no/{out_bill_no}/cancel | 撤销转账

## 类型列表

 - [CreateTransferBillRequest](CreateTransferBillRequest.md)
 - [CreateTransferBillResponse](CreateTransferBillResponse.md)
 - [GetTransferBillByOutBillNoRequest](GetTransferBillByOutBillNoRequest.md)
 - [GetTransferBillByTransferBillNoRequest](GetTransferBillByTransferBillNoRequest.md)
 - [GetTransferBillResponse](GetTransferBillResponse.md)
 - [CancelTransferBillRequest](CancelTransferBillRequest.md)
 - [CancelTransferBillResponse](CancelTransferBillResponse.md)
 - [TransferSceneReportInfo](TransferSceneReportInfo.md)
 - [TransferState](TransferState.md)
 - [TransferNotifyContent](TransferNotifyContent.md)