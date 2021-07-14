// Copyright 2021 Tencent Inc. All rights reserved.

package fileuploader

import (
	"context"
	"io"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

// VideoUploadResponse 视频上传API返回结果
type VideoUploadResponse struct {
	MediaId *string `json:"media_id"` // revive:disable-line:var-naming
}

// VideoUploader 视频上传API
//
// 部分微信支付业务指定商户需要使用视频上传 API来上报视频信息，从而获得必传参数的值：视频MediaID 。
// 接口文档地址：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter2_1_1.shtml
type VideoUploader services.Service

// Upload 上传视频至微信支付
func (u *VideoUploader) Upload(
	ctx context.Context, fileReader io.Reader, filename, contentType string,
) (*VideoUploadResponse, *core.APIResult, error) {
	result, err := (*baseFileUploader)(u).upload(
		ctx, "/v3/merchant/media/video_upload", fileReader, filename, contentType, map[string]interface{}{},
	)
	if err != nil {
		return nil, result, err
	}

	var resp = new(VideoUploadResponse)
	if err = core.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
