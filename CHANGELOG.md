# Changelog

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

