# TransferDetailEntity

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Mchid** | **string** | 微信支付分配的商户号 | 
**OutBatchNo** | **string** | 商户系统内部的商家批次单号，在商户系统内部唯一 | 
**BatchId** | **string** | 微信批次单号，微信商家转账系统返回的唯一标识 | 
**Appid** | **string** | 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid） | 
**OutDetailNo** | **string** | 商户系统内部区分转账批次单下不同转账明细单的唯一标识 | 
**DetailId** | **string** | 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识 | 
**DetailStatus** | **string** | INIT: 初始态。 系统转账校验中 WAIT_PAY: 待确认。待商户确认, 符合免密条件时, 系统会自动扭转为转账中 PROCESSING:转账中。正在处理中，转账结果尚未明确 SUCCESS:转账成功 FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单） | 
**TransferAmount** | **int64** | 转账金额单位为“分” | 
**TransferRemark** | **string** | 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符 | 
**FailReason** | [**FailReasonType**](FailReasonType.md) | 如果转账失败则有失败原因 | [可选] 
**Openid** | **string** | 商户appid下，某用户的openid | 
**UserName** | **string** | 收款方姓名。采用标准RSA算法，公钥由微信侧提供 商户转账时传入了收款用户姓名、查询时会返回收款用户姓名 | [可选] 
**InitiateTime** | **time.Time** | 转账发起的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE | 
**UpdateTime** | **time.Time** | 明细最后一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE | 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


