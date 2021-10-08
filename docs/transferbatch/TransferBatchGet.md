# TransferBatchGet

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Mchid** | **string** | 微信支付分配的商户号  | 
**OutBatchNo** | **string** | 商户系统内部的商家批次单号，在商户系统内部唯一  | 
**BatchId** | **string** | 微信批次单号，微信商家转账系统返回的唯一标识  | 
**Appid** | **string** | 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）  | 
**BatchStatus** | **string** | ACCEPTED:已受理。批次已受理成功，若发起批量转账的30分钟后，转账批次单仍处于该状态，可能原因是商户账户余额不足等。商户可查询账户资金流水，若该笔转账批次单的扣款已经发生，则表示批次已经进入转账中，请再次查单确认   PROCESSING:转账中。已开始处理批次内的转账明细单   FINISHED:已完成。批次内的所有转账明细单都已处理完成   CLOSED:已关闭。可查询具体的批次关闭原因确认  | 
**BatchType** | **string** | API:API方式发起   WEB:页面方式发起  | 
**BatchName** | **string** | 该笔批量转账的名称  | 
**BatchRemark** | **string** | 转账说明，UTF8编码，最多允许32个字符  | 
**CloseReason** | [**CloseReasonType**](CloseReasonType.md) | 如果批次单状态为“CLOSED”（已关闭），则有关闭原因 * &#x60;MERCHANT_REVOCATION&#x60; - 商户主动撤销 * &#x60;OVERDUE_CLOSE&#x60; - 系统超时关闭  | [可选] 
**TotalAmount** | **int64** | 转账金额单位为“分”  | 
**TotalNum** | **int64** | 一个转账批次单最多发起三千笔转账  | 
**CreateTime** | **time.Time** | 批次受理成功时返回，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE  | [可选] 
**UpdateTime** | **time.Time** | 批次最近一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE  | [可选] 
**SuccessAmount** | **int64** | 转账成功的金额，单位为“分”。当批次状态为“PROCESSING”（转账中）时，转账成功金额随时可能变化  | [可选] 
**SuccessNum** | **int64** | 转账成功的笔数。当批次状态为“PROCESSING”（转账中）时，转账成功笔数随时可能变化  | [可选] 
**FailAmount** | **int64** | 转账失败的金额，单位为“分”  | [可选] 
**FailNum** | **int64** | 转账失败的笔数  | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


