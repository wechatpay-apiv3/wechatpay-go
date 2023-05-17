# 微信支付 API v3 Go SDK - marketingbankpackages

# 前言

银行营销API是专为银行商户设置的一套API。目前仅包含“创建快捷交易协议号上传任务”接口，供银行上传快捷交易协议号文件。微信侧收到文件中，将快捷交易协议号转换为微信用户id，存储到指定的号码包文件中。

## 总览
本 SDK 由 WechatPay APIv3 SDK 生成器生成。生成器基于 [OpenAPI Generator](https://openapi-generator.tech) 构建。

- API 版本: 0.1.0

## 接口列表

所有URI均基于微信支付 API 地址：*https://api.mch.weixin.qq.com*

服务名 | 方法名 | HTTP 请求 | 描述
------------ | ------------- | ------------- | -------------
*TasksApi* | [**ListTask**](TasksApi.md#listtask) | **Get** /v3/marketing/bank/packages/{package_id}/tasks | 查询上传任务列表


## 类型列表

 - [CreateTaskRequest](CreateTaskRequest.md)
 - [FileMeta](FileMeta.md)
 - [FormFile](FormFile.md)
 - [ListTaskRequest](ListTaskRequest.md)
 - [ListTaskResponse](ListTaskResponse.md)
 - [Task](Task.md)
 - [TaskStatus](TaskStatus.md)

