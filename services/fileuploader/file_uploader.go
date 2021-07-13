// Copyright 2021 Tencent Inc. All rights reserved.

package fileuploader

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

// baseFileUploader 基础文件上传
type baseFileUploader services.Service

// upload 将指定文件内容上传到指定地址
//
// 注意：urlpath 不要包含微信支付API服务地址，只包含路径即可，例如 `/v3/merchant/media/upload`
func (s *baseFileUploader) upload(
	ctx context.Context, urlpath string, fileReader io.Reader, filename, contentType string,
	extra map[string]interface{},
) (*core.APIResult, error) {
	urlpath = consts.WechatPayAPIServer + urlpath

	content, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return nil, err
	}

	meta := make(map[string]interface{})
	meta["filename"] = core.String(filename)
	meta["sha256"] = core.String(fmt.Sprintf("%x", sha256.Sum256(content)))

	// Override with extra info
	for key, value := range extra {
		meta[key] = value
	}

	metaStr, err := core.ParameterToJSON(meta)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err = core.CreateFormField(writer, "meta", "application/json", []byte(metaStr)); err != nil {
		return nil, err
	}

	if err = core.CreateFormFile(writer, filename, contentType, content); err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	result, err := s.Client.Upload(ctx, urlpath, metaStr, body.String(), writer.FormDataContentType())
	return result, err
}
