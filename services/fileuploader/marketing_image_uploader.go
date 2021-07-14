// Copyright 2021 Tencent Inc. All rights reserved.

package fileuploader

import (
	"context"
	"io"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

// MarketingImageUploadResponse 图片上传API（营销专用）返回结果
type MarketingImageUploadResponse struct {
	MediaUrl *string `json:"media_url"` // revive:disable-line:var-naming
}

// MarketingImageUploader 图片上传API（营销专用）
//
// 通过本接口上传图片后可获得图片url地址。图片url可在微信支付营销相关的API使用，
// 包括商家券、代金券、支付有礼等。
// 接口文档地址：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_0_1.shtml
type MarketingImageUploader services.Service

// Upload 上传图片至微信支付营销系统
func (u *MarketingImageUploader) Upload(
	ctx context.Context, fileReader io.Reader, filename string, contentType string,
) (*MarketingImageUploadResponse, *core.APIResult, error) {
	result, err := (*baseFileUploader)(u).upload(
		ctx, "/v3/marketing/favor/media/image-upload", fileReader, filename, contentType, map[string]interface{}{},
	)
	if err != nil {
		return nil, result, err
	}

	var resp = new(MarketingImageUploadResponse)
	if err = core.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
