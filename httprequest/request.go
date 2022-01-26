package httprequest

import (
	"encoding/json"
	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type httpRequest struct {
	client    *httpclient.Client
	url       string
	headers   http.Header
	body      io.Reader
	urlParams url.Values
	response  *http.Response
	err       error
}

type Options struct {
	Timeout    time.Duration
	RetryCount int
}

func NewRequest(url string, opts ...Options) *httpRequest {
	var c *httpclient.Client
	if len(opts) == 0 {
		c = httpclient.NewClient()
	} else {
		for _, opt := range opts {
			c = httpclient.NewClient(
				httpclient.WithHTTPTimeout(opt.Timeout),
				httpclient.WithRetryCount(opt.RetryCount),
				)
		}
	}
	return &httpRequest{
		client:  c,
		headers: http.Header{},
		url:     url,
	}
}

/**
 * @title Header
 * @description 添加HTTP头
 * @param name string
 * @param value string
 * @return *httpRequest
 **/
func (r *httpRequest) Header(name string, value string) *httpRequest {
	r.headers.Add(name, value)
	return r
}

/**
 * @title Headers
 * @description 批量添加HTTP头
 * @param headers map[string]string
 * @return *httpRequest
 **/
func (r *httpRequest) Headers(headers map[string]string) *httpRequest {
	for k, v := range headers {
		r.headers.Add(k, v)
	}
	return r
}

/**
 * @title UrlParams
 * @description 添加Request Url参数
 * @param params map[string]interface{}
 * @return *httpRequest
 **/
func (r *httpRequest) UrlParams(params map[string]interface{}) *httpRequest {
	p := url.Values{}
	for k, v := range params {
		p.Add(k, cast.ToString(v))
	}
	r.urlParams = p
	return r
}

/**
 * @title BodyParamsByMap
 * @description 添加Request Body参数，内容来自于map
 * @param params map[string]interface{}
 * @return *httpRequest
 **/
func (r *httpRequest) BodyParamsByMap(params map[string]interface{}) *httpRequest {
	p := url.Values{}
	for k, v := range params {
		p.Add(k, cast.ToString(v))
	}
	r.body = strings.NewReader(p.Encode())
	return r
}

/**
 * @title BodyParamsByJson
 * @description 添加Request Body参数，内容来自于json
 * @param json string
 * @return *httpRequest
 **/
func (r *httpRequest) BodyParamsByJson(json string) *httpRequest {
	r.body = strings.NewReader(json)
	r.Header("Content-Type", "application/json")
	return r
}

/**
 * @title Get
 * @description 进行httpRequest GET 请求
 * @return *httpRequest
 **/
func (r *httpRequest) Get() *httpRequest {
	requestUrl := r.buildURL()
	r.response, r.err = r.client.Get(requestUrl, r.headers)
	return r
}

/**
 * @title Post
 * @description 进行httpRequest Post 请求
 * @return *httpRequest
 **/
func (r *httpRequest) Post() *httpRequest {
	requestUrl := r.buildURL()
	r.response, r.err = r.client.Post(requestUrl, r.body, r.headers)
	return r
}

/**
 * @title Put
 * @description 进行httpRequest Put 请求
 * @return *httpRequest
 **/
func (r *httpRequest) Put() *httpRequest {
	requestUrl := r.buildURL()
	r.response, r.err = r.client.Put(requestUrl, r.body, r.headers)
	return r
}

/**
 * @title Delete
 * @description 进行httpRequest Delete 请求
 * @return *httpRequest
 **/
func (r *httpRequest) Delete() *httpRequest {
	requestUrl := r.buildURL()
	r.response, r.err = r.client.Delete(requestUrl, r.headers)
	return r
}

/**
 * @title Patch
 * @description 进行httpRequest Patch 请求
 * @return *httpRequest
 **/
func (r *httpRequest) Patch() *httpRequest {
	requestUrl := r.buildURL()
	r.response, r.err = r.client.Patch(requestUrl, r.body, r.headers)
	return r
}

/**
 * @title Response
 * @description 返回请求的响应体和错误信息
 * @return *http.Response, error
 **/
func (r *httpRequest) Response() (*http.Response, error) {
	return r.response, r.err
}

/**
 * @title ResponseJson
 * @description 将响应按Json的形式解析，并返回对应错误
 * @return *http.Response, error
 **/
func (r *httpRequest) ResponseJson(jsonStruct interface{}) (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}
	body, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return r.response, err
	}
	return r.response, json.Unmarshal(body, &jsonStruct)
}

/**
 * @title ResponseString
 * @description 将响应按字符串形式返回，并返回对应错误
 * @return string, error
 **/
func (r *httpRequest) ResponseString() (string, *http.Response, error) {
	if r.err != nil {
		return "", nil, r.err
	}
	body, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return "", r.response, err
	}
	return string(body), r.response, nil
}

/**
 * @title buildURL
 * @description 返回带url参数的链接地址
 * @return string
 **/
func (r *httpRequest) buildURL() string {
	encodeParams := r.urlParams.Encode()
	requestUrl := r.url
	if encodeParams != "" {
		requestUrl = requestUrl + "?" + encodeParams
	}
	return requestUrl
}
