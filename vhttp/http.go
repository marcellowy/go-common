package vhttp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/marcellowy/go-common/gogf/vlog"
	"github.com/marcellowy/go-common/tools"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FormFileBuffer struct {
	Filename string // filename, example: sample.7z
	Buffer   io.Reader
}

type FormFile struct {
	Filename string // full path filename, example: /path/to/file.7z
}

type Client struct {
	http.Client
	header map[string]string
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

func NewHttpClient() *Client {
	return &Client{}
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

// PostData post data
func (h *Client) PostData(ctx context.Context, url string, data io.Reader) (*Response, error) {
	var request, err = http.NewRequest(http.MethodPost, url, data)
	if err != nil {
		return nil, err
	}
	for k, v := range h.header {
		request.Header.Set(k, v)
	}
	var response *http.Response
	if response, err = h.Do(request); err != nil {
		return nil, err
	}
	defer tools.Close(response.Body)
	return h.makeResponse(ctx, response)
}

// GetData from url
func (h *Client) GetData(ctx context.Context, url string) (*Response, error) {
	var request, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range h.header {
		request.Header.Set(k, v)
	}
	var response *http.Response
	if response, err = h.Do(request); err != nil {
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

// PostUploadForm
// support upload file
//
//		 use FormFileBuffer
//
//			PostUploadForm(ctx, "http:/example.org/upload.php", map[string]any{
//				"fieldName": &FormFileBuffer{
//					"Filename": "test.zip", // filename
//					"Buffer": bytes.NewBuffer([]byte(...)), // buffer from file byte
//				}
//			})
//
//	 use FormFile
//
//		PostUploadForm(ctx, "http:/example.org/upload.php", map[string]any{
//			"fieldName": &FormFile{Filename:"/path/to/filename"},
//		})
func (h *Client) PostUploadForm(ctx context.Context, url string, formData map[string]any) (response *Response, err error) {
	var (
		bb     *bytes.Buffer
		writer *multipart.Writer
	)
	if bb, writer, err = CreateFormBody(ctx, formData); err != nil {
		vlog.Warning(ctx, err)
		return
	}
	_ = writer.Close()
	h.SetHeader("Content-Type", writer.FormDataContentType())

	return h.PostData(ctx, url, bb)
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

func createFormFileFromBuffer(ctx context.Context, fieldName, filename string,
	buffer io.Reader, writer **multipart.Writer) (err error) {
	var uploadWriter io.Writer
	if uploadWriter, err = (*writer).CreateFormFile(fieldName, filepath.Base(filename)); err != nil {
		vlog.Error(ctx, err)
		return
	}
	if _, err = io.Copy(uploadWriter, buffer); err != nil {
		vlog.Error(ctx, err)
		return
	}
	return
}

// CreateFormBody create http form body from map[string]string
func CreateFormBody(ctx context.Context, data map[string]any) (body *bytes.Buffer, writer *multipart.Writer, err error) {
	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	for k, v := range data {
		switch v.(type) {
		case bool:
			var vv = "false"
			vv_ := v.(bool)
			if vv_ {
				vv = "true"
			}
			if err = writer.WriteField(k, vv); err != nil {
				vlog.Error(ctx, err, k, vv)
				return
			}
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			vv := fmt.Sprintf("%d", v)
			if err = writer.WriteField(k, vv); err != nil {
				vlog.Error(ctx, err, k, vv)
				return
			}
		case string, []byte:
			var vv string
			switch v.(type) {
			case string:
				vv = v.(string)
			case []byte:
				vv = tools.BytesToString(v.([]byte))
			}

			if err = writer.WriteField(k, vv); err != nil {
				vlog.Error(ctx, err, k, vv)
				return
			}
		case *FormFile, FormFile:
			var vv *FormFile
			switch v.(type) {
			case *FormFile:
				vv = v.(*FormFile)
			case FormFile:
				vv_ := v.(FormFile)
				vv = &vv_
			}
			if err = createFormFile(ctx, k, vv.Filename, &writer); err != nil {
				vlog.Error(ctx, err)
				return
			}
		case *FormFileBuffer, FormFileBuffer:
			// 处理buffer上传
			var vv *FormFileBuffer
			switch v.(type) {
			case *FormFileBuffer:
				vv = v.(*FormFileBuffer)
			case FormFileBuffer:
				vv_ := v.(FormFileBuffer)
				vv = &vv_
			}
			if err = createFormFileFromBuffer(ctx, k, vv.Filename, vv.Buffer, &writer); err != nil {
				vlog.Error(ctx, err)
				return
			}
		default:
			err = fmt.Errorf("unsupported type: %T", v)
			return
		}

	}
	tools.Close(writer)
	return
}
