# TransferDetailInput

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**OutDetailNo** | **string** | 商户系统内部区分转账批次单下不同转账明细单的唯一标识，要求此参数只能由数字、大小写字母组成 | 
**TransferAmount** | **int64** | 转账金额单位为“分” | 
**TransferRemark** | **string** | 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符 | 
**Openid** | **string** | 商户appid下，某用户的openid | 
**UserName** | **string** | 收款方真实姓名。支持标准RSA算法和国密算法，公钥由微信侧提供 明细转账金额&lt;0.3元时，不允许填写收款用户姓名 明细转账金额 &gt;&#x3D; 2,000元时，该笔明细必须填写收款用户姓名 同一批次转账明细中的姓名字段传入规则需保持一致，也即全部填写、或全部不填写 若商户传入收款用户姓名，微信支付会校验用户openID与姓名是否一致，并提供电子回单 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


