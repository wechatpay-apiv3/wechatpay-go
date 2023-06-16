# FullSendRule

## 属性列表

名称 | 类型 | 描述 | 补充说明
------------ | ------------- | ------------- | -------------
**TransactionAmountMinimum** | **int64** | 消费金额门槛 单位分 | 
**SendContent** | [**SendContentCategory**](SendContentCategory.md) | 发放内容，可选单张券或礼包，选礼包时奖品限定3-5个 | 
**AwardType** | [**AwardType**](AwardType.md) | 奖品类型，暂时只支持商家券 | 
**AwardList** | [**[]AwardBaseInfo**](AwardBaseInfo.md) | 奖品基本信息列表 | 
**MerchantOption** | [**SendMerchantOption**](SendMerchantOption.md) | 发券商户号选取规则，支持选择在用券商户或手动输入商户号两种，选择手动时，发券商户号必填（商家券只支持手动输入） | 
**MerchantIdList** | **[]string** | 发券商户号，列表 | [可选] 

[\[返回类型列表\]](README.md#类型列表)
[\[返回接口列表\]](README.md#接口列表)
[\[返回服务README\]](README.md)


