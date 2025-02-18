package tools

import (
	"bytes"
	"context"
	"github.com/marcellowy/go-common/gogf/vlog"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

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
	defer Close(uploadFile)
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
	Close(writer)
	return
}

func CreateFormData(ctx context.Context, data map[string]string) (body *bytes.Buffer, writer *multipart.Writer, err error) {
	return CreateFormBody(ctx, data)
}
