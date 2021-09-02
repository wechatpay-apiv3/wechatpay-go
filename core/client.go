// Copyright 2021 Tencent Inc. All rights reserved.

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
	"runtime"
	"strings"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/auth"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/credentials"
	"github.com/wechatpay-apiv3/wechatpay-go/core/cipher"
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

// ClientOption 微信支付 API v3 HTTPClient core.Client 初始化参数
type ClientOption interface {
	// Apply 将初始化参数应用到 DialSettings 中
	Apply(settings *DialSettings) error
}

// ErrorOption 错误初始化参数，用于返回错误
type ErrorOption struct{ Error error }

// Apply 返回初始化错误
func (w ErrorOption) Apply(*DialSettings) error {
	return w.Error
}

// Client 微信支付API v3 基础 Client
type Client struct {
	httpClient *http.Client
	credential auth.Credential
	validator  auth.Validator
	signer     auth.Signer
	cipher     cipher.Cipher
}

// NewClient 初始化一个微信支付API v3 HTTPClient
//
// 初始化的时候你可以传递多个配置信息
func NewClient(ctx context.Context, opts ...ClientOption) (*Client, error) {
	settings, err := initSettings(opts)
	if err != nil {
		return nil, fmt.Errorf("init client setting err:%v", err)
	}

	client := initClientWithSettings(ctx, settings)
	return client, nil
}

// NewClientWithDialSettings 使用 DialSettings 初始化一个微信支付API v3 HTTPClient
func NewClientWithDialSettings(ctx context.Context, settings *DialSettings) (*Client, error) {
	if err := settings.Validate(); err != nil {
		return nil, err
	}

	client := initClientWithSettings(ctx, settings)
	return client, nil
}

// NewClientWithValidator 使用原 Client 复制一个新的 Client，并设置新 Client 的 validator。
// 原 Client 不受任何影响
func NewClientWithValidator(client *Client, validator auth.Validator) *Client {
	return &Client{
		httpClient: client.httpClient,
		credential: client.credential,
		signer:     client.signer,
		validator:  validator,
		cipher:     client.cipher,
	}
}

func initClientWithSettings(_ context.Context, settings *DialSettings) *Client {
	client := &Client{
		signer:     settings.Signer,
		validator:  settings.Validator,
		credential: &credentials.WechatPayCredentials{Signer: settings.Signer},
		httpClient: settings.HTTPClient,
		cipher:     settings.Cipher,
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: consts.DefaultTimeout,
		}
	}
	return client
}

func initSettings(opts []ClientOption) (*DialSettings, error) {
	var (
		o   DialSettings
		err error
	)
	for _, opt := range opts {
		if err = opt.Apply(&o); err != nil {
			return nil, err
		}
	}
	if err := o.Validate(); err != nil {
		return nil, err
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

// Delete 向微信支付发送一个 HTTP Delete 请求
func (client *Client) Delete(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	return client.requestWithJSONBody(ctx, http.MethodDelete, requestURL, requestBody)
}

// Upload 向微信支付发送上传文件
// 推荐使用 services/fileuploader 中各上传接口的实现
func (client *Client) Upload(ctx context.Context, requestURL, meta, reqBody, formContentType string) (
	*APIResult, error,
) {
	return client.doRequest(ctx, http.MethodPost, requestURL, nil, formContentType, strings.NewReader(reqBody), meta)
}

func (client *Client) requestWithJSONBody(ctx context.Context, method, requestURL string, body interface{}) (
	*APIResult, error,
) {
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
	signBody string,
) (*APIResult, error) {

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
	// Fixed Headers > Per-Request Header Parameters

	// Add Request Header Parameters
	for key, values := range header {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	// Set Fixed Headers
	request.Header.Set(consts.Accept, "*/*")
	request.Header.Set(consts.ContentType, contentType)

	ua := fmt.Sprintf(consts.UserAgentFormat, consts.Version, runtime.GOOS, runtime.Version())
	request.Header.Set(consts.UserAgent, ua)

	// Set Authentication
	if authorization, err = client.credential.GenerateAuthorizationHeader(
		ctx, method, request.URL.RequestURI(),
		signBody,
	); err != nil {
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
// 特别地，如果需要为当前请求设置 Header，可以使用本方法
func (client *Client) Request(
	ctx context.Context,
	method, requestPath string,
	headerParams http.Header,
	queryParams url.Values,
	postBody interface{},
	contentType string,
) (result *APIResult, err error) {

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

// EncryptRequest 使用 cipher 对请求结构进行原地加密，并返回加密所用的平台证书的序列号。
// 未设置 cipher 时将跳过加密，并返回空序列号。
//
// 本方法会对结构中的敏感字段进行原地加密，因此需要传入结构体的指针。
func (client *Client) EncryptRequest(ctx context.Context, req interface{}) (string, error) {
	if client.cipher == nil {
		return "", nil
	}
	return client.cipher.Encrypt(ctx, req)
}

// DecryptResponse 使用 cipher 对应答结构进行原地解密，未设置 cipher 时将跳过解密
//
// 本方法会对结构中的敏感字段进行原地解密，因此需要传入结构体的指针。
func (client *Client) DecryptResponse(ctx context.Context, resp interface{}) error {
	if client.cipher == nil {
		return nil
	}
	return client.cipher.Decrypt(ctx, resp)
}

// Sign 使用 signer 对字符串进行签名
func (client *Client) Sign(ctx context.Context, message string) (result *auth.SignatureResult, err error) {
	return client.signer.Sign(ctx, message)
}

// CheckResponse 校验请求是否成功
//
// 当http回包的状态码的范围不是200-299之间的时候，会返回相应的错误信息，主要包括http状态码、回包错误码、回包错误信息提示
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	slurp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response, read body error: %w", err)
	}
	_ = resp.Body.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(slurp))
	apiError := &APIError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(slurp),
	}
	// 忽略 JSON 解析错误，均返回 apiError
	_ = json.Unmarshal(slurp, apiError)
	return apiError
}

// UnMarshalResponse 将回包组织成结构化数据
func UnMarshalResponse(httpResp *http.Response, resp interface{}) error {
	body, err := ioutil.ReadAll(httpResp.Body)
	_ = httpResp.Body.Close()

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
// 示例内容
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
		if strings.EqualFold(a, needle) {
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
