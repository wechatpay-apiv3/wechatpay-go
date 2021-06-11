// Copyright 2021 Tencent Inc.  All rights reserved.

// Package core 微信支付 API v3 Go SDK HTTPClient 基础库，你可以使用它来创建一个 Client，并向微信支付发送 HTTP 请求
//
// 初始化 Client 时，你需要指定以下参数：
//  - Credential 用于生成 HTTP Header 中的 Authorization 信息，微信支付 API v3依赖该值来保证请求的真实性和数据的完整性
//  - Validator 用于对微信支付的应答进行校验，避免被恶意攻击
package core

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

var (
	regJSONTypeCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	regXMLTypeCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIResult 微信支付API v3 请求结果
type APIResult struct {
	// 本次请求所使用的 HTTPRequest
	Request *http.Request
	// 本次请求所获得的 HTTPResponse
	Response *http.Response
}

// Client 微信支付API v3 基础 Client
type Client struct {
	httpClient    *http.Client
	defaultHeader http.Header
	credential    auth.Credential
	validator     auth.Validator
	signer        auth.Signer
}

// NewClient 初始化一个微信支付API v3 HTTPClient
//
// 初始化的时候你可以传递多个配置信息
func NewClient(ctx context.Context, opts ...ClientOption) (client *Client, err error) {
	settings, err := initSettings(opts)
	if err != nil {
		return nil, fmt.Errorf("init client setting err:%v", err)
	}
	client = &Client{
		signer:        settings.Signer,
		validator:     settings.Validator,
		credential:    settings.Credential,
		httpClient:    settings.HTTPClient,
		defaultHeader: settings.Header,
	}
	return client, nil
}

func NewClientWithValidator(client *Client, validator auth.Validator) *Client {
	return &Client{
		httpClient : client.httpClient,
		defaultHeader : client.defaultHeader,
		credential : client.credential,
		signer : client.signer,
		validator: validator,
	}
}

func initSettings(opts []ClientOption) (*dialSettings, error) {
	var o dialSettings
	for _, opt := range opts {
		opt.Apply(&o)
	}
	if err := o.Validate(); err != nil {
		return nil, err
	}
	if o.HTTPClient == nil {
		o.HTTPClient = &http.Client{}
	}
	if o.Timeout != 0 {
		o.HTTPClient.Timeout = o.Timeout
	}
	if o.Credential == nil {
		o.Credential = &credentials.WechatPayCredentials{
			Signer: o.Signer,
		}
	}
	return &o, nil
}

// Get 向微信支付发送一个 HTTP Get 请求
func (client *Client) Get(ctx context.Context, requestURL string) (*APIResult, error) {
	return client.doRequest(ctx, http.MethodGet, requestURL, nil, consts.ApplicationJSON, nil, "")
}

// Post 向微信支付发送一个 HTTP Post 请求
func (client *Client) Post(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	return client.requestWithJSONBody(ctx, http.MethodPost, requestURL, requestBody)
}

// Patch 向微信支付发送一个 HTTP Patch 请求
func (client *Client) Patch(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	return client.requestWithJSONBody(ctx, http.MethodPatch, requestURL, requestBody)
}

// Put 向微信支付发送一个 HTTP Put 请求
func (client *Client) Put(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	return client.requestWithJSONBody(ctx, http.MethodPut, requestURL, requestBody)
}

// Delete 向微信支付发送一个 Http Delete 请求
func (client *Client) Delete(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	return client.requestWithJSONBody(ctx, http.MethodDelete, requestURL, requestBody)
}

// Upload 向微信支付发送上传文件
func (client *Client) Upload(ctx context.Context, requestURL, meta, reqBody, formContentType string) (*APIResult, error) {
	return client.doRequest(ctx, http.MethodPost, requestURL, nil, formContentType, strings.NewReader(reqBody), meta)
}

func (client *Client) requestWithJSONBody(ctx context.Context, method, requestURL string, body interface{}) (*APIResult, error) {
	reqBody, err := setBody(body, consts.ApplicationJSON)
	if err != nil {
		return nil, err
	}

	return client.doRequest(ctx, method, requestURL, nil, consts.ApplicationJSON, reqBody, reqBody.String())
}

func (client *Client) doRequest(
	ctx context.Context,
	method string,
	requestURL string,
	header http.Header,
	contentType string,
	reqBody io.Reader,
	signBody string) (*APIResult, error) {

	var (
		err           error
		authorization string
		request       *http.Request
	)

	// Construct Request
	if request, err = http.NewRequestWithContext(ctx, method, requestURL, reqBody); err != nil {
		return nil, err
	}

	// Header Setting Priority:
	// Fixed Headers > Per-Request Header Parameters > Client Default Headers

	// Add Client Default Headers
	for key, values := range client.defaultHeader {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	// Add Request Header Parameters
	if header != nil {
		for key, values := range header {
			for _, v := range values {
				request.Header.Add(key, v)
			}
		}
	}

	// Set Fixed Headers
	request.Header.Set(consts.Accept, "*/*")
	request.Header.Set(consts.ContentType, contentType)
	request.Header.Set(consts.UserAgent, consts.UserAgentContent)

	// Set Authentication
	if authorization, err = client.credential.GenerateAuthorizationHeader(ctx, method, request.URL.RequestURI(),
		signBody); err != nil {
		return nil, fmt.Errorf("generate authorization err:%s", err.Error())
	}
	request.Header.Set(consts.Authorization, authorization)

	// Send HTTP Request
	result, err := client.doHTTP(request)
	if err != nil {
		return result, err
	}
	// Check if Success
	if err = CheckResponse(result.Response); err != nil {
		return result, err
	}
	// Validate WechatPay Signature
	if err = client.validator.Validate(ctx, result.Response); err != nil {
		return result, err
	}
	return result, nil
}

// Request 向微信支付发送请求
//
// 相比于 Get / Post / Put / Patch / Delete 方法，本方法可以设置更多内容
func (client *Client) Request(
	ctx context.Context,
	method, requestPath string,
	headerParams http.Header,
	queryParams url.Values,
	postBody interface{},
	contentType string) (result *APIResult, err error) {

	// Setup path and query parameters
	varURL, err := url.Parse(requestPath)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := varURL.Query()
	for k, values := range queryParams {
		for _, v := range values {
			query.Add(k, v)
		}
	}

	// Encode the parameters.
	varURL.RawQuery = query.Encode()

	if postBody == nil {
		return client.doRequest(ctx, method, varURL.String(), headerParams, contentType, nil, "")
	}

	// Detect postBody type and set body content
	if contentType == "" {
		contentType = consts.ApplicationJSON
	}
	var body *bytes.Buffer
	body, err = setBody(postBody, contentType)
	if err != nil {
		return nil, err
	}
	return client.doRequest(ctx, method, varURL.String(), headerParams, contentType, body, body.String())
}

func (client *Client) doHTTP(req *http.Request) (result *APIResult, err error) {
	result = &APIResult{
		Request: req,
	}

	result.Response, err = client.httpClient.Do(req)
	return result, err
}

// CheckResponse 校验请求是否成功
//
// 当http回包的状态码的范围不是200-299之间的时候，会返回相应的错误信息，主要包括http状态码、回包错误码、回包错误信息提示
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err == nil {
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(slurp))
		apiError := &APIError{
			StatusCode: resp.StatusCode,
			Header:     resp.Header,
			Body:       string(slurp),
		}
		err = json.Unmarshal(slurp, apiError)
		if err == nil {
			return apiError
		}
	}
	return &APIError{
		StatusCode: resp.StatusCode,
		Body:       string(slurp),
		Header:     resp.Header,
	}
}

// UnMarshalResponse 将回包组织成结构化数据
func UnMarshalResponse(httpResp *http.Response, resp interface{}) error {
	body, err := ioutil.ReadAll(httpResp.Body)
	httpResp.Body.Close()

	if err != nil {
		return err
	}

	httpResp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	return nil
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

// Set Request body from an interface
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(**os.File); ok {
		_, err = bodyBuf.ReadFrom(*fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if regJSONTypeCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if regXMLTypeCheck.MatchString(contentType) {
		err = xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// SelectHeaderContentType select a content type from the available list.
func SelectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return consts.ApplicationJSON
	}
	if contains(contentTypes, consts.ApplicationJSON) {
		return consts.ApplicationJSON
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// ParameterToString 将参数转换为字符串，并使用指定分隔符分隔列表参数
func ParameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// ParameterToJSON 将参数转换为 Json 字符串
func ParameterToJSON(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}
