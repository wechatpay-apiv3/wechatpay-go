// Copyright 2021 Tencent Inc. All rights reserved.

package fileuploader

import (
	"context"
	"io"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

// ImageUploadResponse 图片上传API返回结果
type ImageUploadResponse struct {
	MediaId *string `json:"media_id"` // revive:disable-line:var-naming
}

// ImageUploader 图片上传API
//
// 部分微信支付业务指定商户需要使用图片上传 API来上报图片信息，从而获得必传参数的值：图片MediaID 。
// 接口文档地址：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_1.shtml
type ImageUploader services.Service

// Upload 上传图片至微信支付
func (u *ImageUploader) Upload(
	ctx context.Context, fileReader io.Reader, filename, contentType string,
) (*ImageUploadResponse, *core.APIResult, error) {
	result, err := (*baseFileUploader)(u).upload(
		ctx, "/v3/merchant/media/upload", fileReader, filename, contentType, map[string]interface{}{},
	)
	if err != nil {
		return nil, result, err
	}

	var resp = new(ImageUploadResponse)
	if err = core.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, err
}
