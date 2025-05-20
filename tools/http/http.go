package http

import (
	"bytes"
	"context"
	"fmt"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/tools"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type sendFormUpload struct {
	FieldName string
	Path      string
}

type OptionFunc func(*Client)

func WithTimeout(timeout time.Duration) OptionFunc {
	return func(client *Client) {
		client.timeout = timeout
	}
}

func WithProxy(addr string) OptionFunc {
	return func(client *Client) {
		client.proxy = true
		client.proxyAddr = addr
	}
}

func WithProxyFromEnvironment() OptionFunc {
	return func(client *Client) {
		client.proxy = true
		client.proxyAddr = ""
	}
}

func WithHeader(header map[string]string) OptionFunc {
	return func(client *Client) {
		if client.header == nil {
			client.header = make(map[string]string)
		}
		for k, v := range header {
			client.header[k] = v
		}
	}
}

func WithUserAgent(userAgent string) OptionFunc {
	return func(client *Client) {
		if client.header == nil {
			client.header = make(map[string]string)
		}
		client.header["User-Agent"] = userAgent
	}
}

type Client struct {
	client    *http.Client
	proxy     bool
	proxyAddr string
	header    map[string]string
	timeout   time.Duration
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

func NewHttpClient(opts ...OptionFunc) *Client {
	client := &Client{}
	for _, opt := range opts {
		opt(client)
	}
	if client.timeout == 0 {
		client.timeout = time.Second * 5
	}
	client.client = &http.Client{Timeout: client.timeout}
	client.configProxy()
	return client
}

func (h *Client) SetHeader(key, value string) {
	if h.header == nil {
		h.header = make(map[string]string)
	}
	h.header[key] = value
}

func (h *Client) DelHeader(key string) {
	if h.header == nil {
		return
	}
	delete(h.header, key)
}

func (h *Client) ClearHeader(key string) {
	h.header = nil
}

func (h *Client) configProxy() {
	if !h.proxy {
		return
	}
	var (
		proxy *url.URL
		err   error
	)
	if proxy, err = url.Parse(h.proxyAddr); err != nil {
		return
	}
	var pu = http.ProxyURL(proxy)
	if h.proxyAddr == "" {
		// use system proxy
		pu = http.ProxyFromEnvironment
	}
	h.client.Transport = &http.Transport{
		Proxy: pu,
	}
}

// Post
// Don't forget close response body
func (h *Client) Post(ctx context.Context, url string, data io.Reader) (*Response, error) {
	var request, err = http.NewRequestWithContext(ctx, http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}
	for k, v := range h.header {
		request.Header.Set(k, v)
	}
	var response *http.Response
	if response, err = h.client.Do(request); err != nil {
		return nil, err
	}
	defer tools.Close(response.Body)
	return h.makeResponse(ctx, response)
}

// Get
// Don't forget close response body
func (h *Client) Get(ctx context.Context, url string) (*Response, error) {
	var request, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range h.header {
		request.Header.Set(k, v)
	}
	var response *http.Response
	if response, err = h.client.Do(request); err != nil {
		return nil, err
	}
	defer tools.Close(response.Body)
	return h.makeResponse(ctx, response)
}

func (h *Client) makeResponse(ctx context.Context, response *http.Response) (resp *Response, err error) {
	if response.StatusCode != http.StatusOK {
		return &Response{
				StatusCode: response.StatusCode,
				Header:     response.Header,
			},
			fmt.Errorf("http status code: %d", response.StatusCode)
	}
	var (
		body []byte
	)
	if body, err = io.ReadAll(response.Body); err != nil {
		return nil, err
	}
	resp = &Response{
		StatusCode: http.StatusOK,
		Header:     response.Header,
		Body:       body,
	}
	return
}

// PostForm
// support upload file
func (h *Client) PostForm(ctx context.Context, url string, formData map[string]string) (response *Response, err error) {
	bb := &bytes.Buffer{}
	writer := multipart.NewWriter(bb)
	var uploadFiles []*sendFormUpload
	for k, v := range formData {
		if strings.Contains(v, "@file:") {
			uploadFiles = append(uploadFiles, &sendFormUpload{
				FieldName: k,
				Path:      strings.ReplaceAll(v, "@file:", ""),
			})

			continue
		}
		_ = writer.WriteField(k, v)
	}

	// create upload form
	for _, v := range uploadFiles {
		if err = createFormFile(ctx, v.FieldName, v.Path, &writer); err != nil {
			vlog.Error(ctx, err)
			return
		}
	}
	_ = writer.Close()
	h.SetHeader("Content-Type", writer.FormDataContentType())

	return h.Post(ctx, url, bb)
}

func createFormFile(ctx context.Context, fieldName, filename string, writer **multipart.Writer) (err error) {
	var uploadWriter io.Writer
	if uploadWriter, err = (*writer).CreateFormFile(fieldName, filepath.Base(filename)); err != nil {
		vlog.Error(ctx, err)
		return
	}
	var uploadFile *os.File
	if uploadFile, err = os.Open(filename); err != nil {
		vlog.Error(ctx, err)
		return
	}
	defer tools.Close(uploadFile)
	if _, err = io.Copy(uploadWriter, uploadFile); err != nil {
		vlog.Error(ctx, err)
		return
	}

	return
}

// CreateFormBody create http form body from map[string]string
func CreateFormBody(ctx context.Context, data map[string]string) (body *bytes.Buffer, writer *multipart.Writer, err error) {
	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	for k, v := range data {
		if strings.Contains(v, "@file:") {
			// upload file
			// the "v" is file absolute path
			file := strings.ReplaceAll(v, "@file:", "")
			if err = createFormFile(ctx, k, file, &writer); err != nil {
				vlog.Error(ctx, err)
				return
			}
			continue
		}
		if err = writer.WriteField(k, v); err != nil {
			vlog.Error(ctx, err)
			return
		}
	}
	tools.Close(writer)
	return
}

func CreateFormData(ctx context.Context, data map[string]string) (body *bytes.Buffer, writer *multipart.Writer, err error) {
	return CreateFormBody(ctx, data)
}
