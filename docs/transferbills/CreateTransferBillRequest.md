# CreateTransferBillRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Appid** | **string** | 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid） | 
**OutBillNo** | **string** | 商户系统内部的商家订单号，要求32个字符内，只能是数字、大小写字母，在商户系统内部唯一 | 
**Openid** | **string** | 用户在直连商户appid下的唯一标识 | 
**TransferAmount** | **int64** | 转账金额，单位为"分" | 
**TransferRemark** | **string** | 转账备注，会显示在微信用户的零钱记录中 | 
**TransferSceneReportInfos** | [**[]TransferSceneReportInfo**](TransferSceneReportInfo.md) | 转账场景信息，用于风控和合规 | 
**UserName** | **string** | 收款方真实姓名。支持标准RSA算法和国密算法，公钥由微信侧提供 | [可选] 
**NotifyUrl** | **string** | 异步接收微信支付结果通知的回调地址 | [可选] 
**UserRecvPerception** | **string** | 用户收款时感知到的收款原因 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)

