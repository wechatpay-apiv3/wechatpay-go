# ListCouponsByFilterRequest

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**Openid** | **string** | 用户在商户appid 下的唯一标识 | 
**Appid** | **string** | 微信为发券方商户分配的公众账号ID，接口传入的所有appid应该为公众号的appid（在mp.weixin.qq.com申请的），不能为APP的appid（在open.weixin.qq.com申请的）。 | 
**StockId** | **string** | 批次号，是否指定批次号查询，填写available_mchid，该字段不生效 | [可选] 
**Status** | **string** | 代金券状态：SENDED-可用，USED-已实扣，填写available_mchid，该字段不生效 | [可选] 
**CreatorMchid** | **string** | 批次创建方商户号。创建批次的商户号，批次发放商户号，可用商户号三个参数，任意选填一个。 | [可选] 
**SenderMchid** | **string** | 批次发放商户号。创建批次的商户号，批次发放商户号，可用商户号三个参数，任意选填一个。 | [可选] 
**AvailableMchid** | **string** | 可用商户号。 创建批次的商户号，批次发放商户号，可用商户号三个参数，任意选填一个。 | [可选] 
**Offset** | **int64** | 分页页码，默认0，填写available_mchid，该字段不生效 | [可选] 
**Limit** | **int64** | 分页大小，默认20，填写available_mchid，该字段不生效 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


