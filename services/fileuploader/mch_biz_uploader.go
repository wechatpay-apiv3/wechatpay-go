// Copyright 2021 Tencent Inc. All rights reserved.

package fileuploader

import (
	"context"
	"io"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

// MchBizUploadResponse 商户上传反馈图片API返回结果
type MchBizUploadResponse struct {
	MediaId *string `json:"media_id"` // revive:disable-line:var-naming
}

// MchBizUploader 商户上传反馈图片API
//
// 将媒体图片进行二进制转换，得到的媒体图片二进制内容，在请求body中上传此二进制内容。 媒体图片只支持jpg、png、bmp格式，文件大小不能超过2M。
// 接口文档地址：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_10.shtml
type MchBizUploader services.Service

// Upload 上传反馈图片至微信支付
func (u *MchBizUploader) Upload(
	ctx context.Context, fileReader io.Reader, filename, contentType string,
) (*MchBizUploadResponse, *core.APIResult, error) {
	result, err := (*baseFileUploader)(u).upload(
		ctx, "/v3/merchant-service/images/upload", fileReader, filename, contentType, map[string]interface{}{},
	)
	if err != nil {
		return nil, result, err
	}

	var resp = new(MchBizUploadResponse)
	if err = core.UnMarshalResponse(result.Response, resp); err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
