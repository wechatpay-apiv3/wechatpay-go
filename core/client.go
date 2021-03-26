// Copyright 2020 Tencent Inc.  All rights reserved.

// Package core 微信支付api v3 go http-client 基础库，你可以使用它来创建一个client，并向微信支付发送http请求
//
// 只需要你在初始化客户端的时候，传递credential以及validator
//
// credential用来生成http header中的authorization信息
//
// validator则用来校验回包是否被篡改
//
// 如果http请求返回的err为nil，一般response.Body 都不为空，你可以尝试对其进行序列化
//
// 请注意及时关闭response.Body
package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/core/setting"
)

// Client 微信支付apiv3 基础client
type Client struct {
	hc         *http.Client
	request    *http.Request
	header     *http.Header
	credential auth.Credential
	validator  auth.Validator
}

//NewClient 初始化一个微信支付apiv3 http client
//
// 初始化的时候你可以传递多个配置信息, 比如
//	ctx := context.Background()
//	opts := []option.ClientOption{
//		option.WithMerchant(mchID, mchCertificateSerialNumber, privateKey), // 设置商户信息，用于生成签名信息
//		option.WithWechatPay(wechatPayCertificateList),  // 设置微信支付平台证书信息，对回包进行校验
//		option.WithHTTPClient(&http.Client{}), // 可以不设置
//		option.WithTimeout(2 * time.Second),   // 自行进行超时时间配置
//	}
//	client, err := core.NewClient(ctx, opts...)
//	if err != nil {
//		log.Printf("new wechat pay client err:%s", err.Error())
//		return
//	}
func NewClient(ctx context.Context, opts ...option.ClientOption) (client *Client, err error) {
	settings, err := initSettings(opts)
	if err != nil {
		return nil, fmt.Errorf("init client setting err:%v", err)
	}
	client = &Client{
		validator:  settings.Validator,
		credential: settings.Credential,
		hc:         settings.HTTPClient,
		header:     settings.Header,
	}
	return client, nil
}

func initSettings(opts []option.ClientOption) (*setting.DialSettings, error) {
	var o setting.DialSettings
	for _, opt := range opts {
		opt.Apply(&o)
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}
	if o.HTTPClient == nil {
		o.HTTPClient = &http.Client{}
	}
	if o.Header == nil {
		o.Header = &http.Header{}
	}
	if o.Timeout != 0 {
		o.HTTPClient.Timeout = o.Timeout
	}
	return &o, nil
}

// Get 向微信支付发送一个http get请求
func (client *Client) Get(ctx context.Context, requestURL string) (*http.Response, error) {
	return client.doRequest(ctx, http.MethodGet, requestURL, consts.ApplicationJSON, "", "")
}

// Post 向微信支付发送一个http post请求
func (client *Client) Post(ctx context.Context, requestURL string, requestBody interface{}) (*http.Response, error) {
	return client.do(ctx, http.MethodPost, requestURL, requestBody)
}

// Patch 向微信支付发送一个http patch请求
func (client *Client) Patch(ctx context.Context, requestURL string, requestBody interface{}) (*http.Response, error) {
	return client.do(ctx, http.MethodPatch, requestURL, requestBody)
}

// Put 向微信支付发送一个http put请求
func (client *Client) Put(ctx context.Context, requestURL string, requestBody interface{}) (*http.Response, error) {
	return client.do(ctx, http.MethodPut, requestURL, requestBody)
}

// Delete 向微信支付发送一个http delete请求
func (client *Client) Delete(ctx context.Context, requestURL string, requestBody interface{}) (*http.Response, error) {
	return client.do(ctx, http.MethodDelete, requestURL, requestBody)
}

// Upload 向微信支付发送上传图片或视频文件请求
//
// 上传接口文档https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/tool/chapter3_1.shtml
//
// 你需要传入上传接口的url、文件名、文件流以及meta部分，
//
// 当前meta部分需要你自己定义一个结构体，后续我们会将上传接口以api的形式提供，直接封装好
//
//	pictureByes, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		return err
//	}
//	// 计算文件序列化后的sha256
//	h := sha256.New()
//	if _, err = h.Write(pictureByes); err != nil {
//		return err
//	}
//	type Meta struct {
//		FileName string `json:"filename" binding:"required"` // 商户上传的媒体图片的名称，商户自定义，必须以JPG、BMP、PNG为后缀。
//		Sha256   string `json:"sha256" binding:"required"`   // 图片文件的文件摘要，即对图片文件的二进制内容进行sha256计算得到的值。
//	}
//	meta := &Meta{}
//	pictureSha256 := h.Sum(nil)
//	meta.FileName = fileName
//	meta.Sha256 = fmt.Sprintf("%x", string(pictureSha256))
//	metaByte, _ := json.Marshal(meta)
//	requestBody := &bytes.Buffer{}
//	writer := multipart.NewWriter(requestBody)
//	if err = CreateFormField(writer, "meta", "application/json", metaByte); err != nil {
//		return err
//	}
//	if err = CreateFormFile(writer, fileName, "image/jpg", pictureByes); err != nil {
//		return err
//	}
//	if err = writer.Close(); err != nil {
//		return err
//	}
//	response, err := client.Upload(ctx, uploadURL, string(metaByte), requestBody.String(), writer.FormDataContentType())
func (client *Client) Upload(ctx context.Context, requestURL, meta, reqBody, formContentType string) (*http.Response,
	error) {
	return client.doRequest(ctx, http.MethodPost, requestURL, formContentType, reqBody, meta)
}

// RequestInfo 返回client请求里的http request信息
func (client *Client) RequestInfo() *http.Request {
	return client.request
}

func (client *Client) do(ctx context.Context, method, requestURL string, body interface{}) (*http.Response, error) {
	var reqBody string
	var err error
	if reqBody, err = marshalReqBody(body); err != nil {
		return nil, err
	}
	return client.doRequest(ctx, method, requestURL, consts.ApplicationJSON, reqBody, reqBody)
}

func (client *Client) doRequest(ctx context.Context,
	method, requestURL, contentType, reqBody, signBody string) (*http.Response, error) {
	var err error
	var authorization string
	if client.request, err = http.NewRequestWithContext(ctx, method, requestURL,
		strings.NewReader(reqBody)); err != nil {
		return nil, err
	}
	client.request.Header.Set(consts.Accept, "*/*")
	client.request.Header.Set(consts.ContentType, contentType)
	client.request.Header.Set(consts.UserAgent, consts.UserAgentContent)
	if authorization, err = client.credential.GenerateAuthorizationHeader(ctx, method, client.request.URL.RequestURI(),
		signBody); err != nil {
		return nil, fmt.Errorf("generate authorization err:%s", err.Error())
	}
	client.request.Header.Set(consts.Authorization, authorization)
	for key, value := range *client.header {
		client.request.Header.Set(key, value[0])
	}
	response, err := client.hc.Do(client.request)
	if err != nil {
		return response, err
	}
	if err = CheckResponse(response); err != nil {
		return response, err
	}
	if err = client.validator.Validate(ctx, response); err != nil {
		return response, err
	}
	return response, nil
}

func marshalReqBody(body interface{}) (string, error) {
	if body == nil {
		return "", nil
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("json marshal body err:%v", err)
	}
	return string(bodyBytes), nil
}

// CheckResponse 校验回包是否有错误
//
// 当http回包的状态码的范围不是200-299之间的时候，会返回相应的错误信息，主要包括http状态码、回包错误码、回包错误信息提示
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(res.Body)
	if err == nil {
		res.Body = ioutil.NopCloser(bytes.NewBuffer(slurp))
		jerr := &errors.Error{StatusCode: res.StatusCode}
		err = json.Unmarshal(slurp, jerr)
		if err == nil {
			return jerr
		}
	}
	return &errors.Error{
		StatusCode: res.StatusCode,
		Body:       string(slurp),
		Header:     res.Header,
	}
}

// CreateFormField 设置form-data 中的普通属性
//
//示例内容
//	Content-Disposition: form-data; name="meta";
//	Content-Type: application/json
//
//	{ "filename": "file_test.mp4", "sha256": " hjkahkjsjkfsjk78687dhjahdajhk " }
//
// 如果要设置上述内容
//	CreateFormField(w, "meta", "application/json", meta)
func CreateFormField(w *multipart.Writer, fieldName, contentType string, fieldValue []byte) error {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s";`, fieldName))
	h.Set("Content-Type", contentType)
	part, err := w.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = part.Write(fieldValue)
	return err
}

// CreateFormFile 设置form-data中的文件
//
// 示例内容：
//	Content-Disposition: form-data; name="file"; filename="file_test.mp4";
//	Content-Type: video/mp4
//
//	pic1  //pic1即为媒体视频的二进制内容
//
// 如果要设置上述内容，则CreateFormFile(w, "file_test.mp4", "video/mp4", pic1)
func CreateFormFile(w *multipart.Writer, filename, contentType string, file []byte) error {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set("Content-Type", contentType)
	part, err := w.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = part.Write(file)
	return err
}
