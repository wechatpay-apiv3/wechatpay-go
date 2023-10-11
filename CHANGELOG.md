# Changelog

## [0.2.18] - 2023-10-11

### Changed

+ 更新小店 `retailstore`
+ 更新点金计划 `goldplan`

## [0.2.17] - 2023-06-29

### Added

+ 增加点金计划 `goldplan`
+ 增加微信支付分停车 `wexinpayscoreparking`
+ 增加商家券 `merchantexclusivecoupon`
+ 增加支付有礼 `giftactivity`
+ 增加代金券 `cachcoupons`

### Changed

+ 更新商家转账 `transferbatch`

## [0.2.15] - 2022-12-12

### Added

+ 增加微工卡 `payrollcard`
+ 增加微信支付刷码乘车 `weixinpayscanandride`
+ 增加爱心餐 `lovefeast`

### Changed

+ 批量转账 `transferbatch` 更名为 商家转账

## [0.2.9] - 2021-10-15

### Fixed

+ 修复批量转账接口部分字段误设置为 `os.File` 的问题，包括服务商（[代码](services/partnertransferbatch)，[文档](docs/partnertransferbatch)）与直连商户（[代码](services/transferbatch)，[文档](docs/transferbatch)）

## [0.2.8] - 2021-10-08

### Added

+ 批量转账接口，支持服务商（[代码](services/partnertransferbatch)，[文档](docs/partnertransferbatch)）和直连商户（[代码](services/transferbatch)，[文档](docs/transferbatch)）。

## [0.2.7] - 2021-09-14

### Fixed

+ 修复服务商支付接口`partnerpayments`与实际契约不匹配的问题

## [0.2.6] - 2021-08-26

### Added

+ 增加分账接口`profitsharing`

### Changed

+ 移除 Enum 类型的 UnmarshalJSON 对数据的检查，避免因为增加枚举值导致的不兼容问题

## [0.2.5] - 2021-07-22

### Fixed

+ 修复获取证书序列号时会丢弃其头部0的问题

## [0.2.4] - 2021-07-21

### Added

+ 增加服务商支付接口`partnerpayments`

## [0.2.3] - 2021-07-13

### Fixed

+ 修复`notify.Handler`读取服务端`Request`的错误

## [0.2.2] - 2021-07-09

### Added

+ 微信支付境内退款（refunddomestic）接口SDK

### Changed

+ `BREAKING CHANGE` 将现有接口SDK中的整型参数统一为`int64`。受影响接口包括:
  <details>
  <summary>Click to expand!</summary>
  
    + payments/app
    + payments/h5
    + payments/jsapi
    + payments/native
  </details>

## [0.2.1] - 2021-06-25

### Added

+ 平台证书下载器与自动下载管理器
+ 回调通知验签、解密器
+ 敏感字段自动加解密库
+ 提供拉起支付签名计算接口
+ 一键构建Client的签名器、验签器的复合Option
+ 平台证书下载命令行工具

### Changed

+ `core.WithXXX` 等方法移动至 `option.WithXXX`

### Removed

+ 移除 `core.Client` 自定义 Header 选项，如需设置 Header 可使用 `client.Request` 接口发起请求
+ 移除 `core.Client` 自定义 Timeout 选项，如需设置 Timeout 可设置自定义 HTTPClient

## [0.2.0] - 2021-05-31

### Added

+ 增加微信支付支付API的SDK
+ 增加文件上传API的SDK

## [0.1.0] - 2020-03-03

+ Initial version

