package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestCreateFormBody(t *testing.T) {

	uploadFilename := "test-upload-file.txt"
	_ = os.WriteFile(uploadFilename, []byte("test123"), os.ModePerm)
	defer func() {
		_ = os.Remove(uploadFilename)
	}()

	var testKey = "test"
	var testValue = "1"

	type args struct {
		ctx  context.Context
		data map[string]string
	}
	tests := []struct {
		name     string
		args     args
		wantBody *bytes.Buffer
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				data: map[string]string{
					testKey: testValue,
					"test2": "2",
					"file":  "@file:" + uploadFilename,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			svr := g.Server()
			//svr := ghttp.Server{}
			svr.SetAddr("127.0.0.1:47632")
			svr.BindHandler("/test", func(r *ghttp.Request) {
				uploadFile := r.GetUploadFile("file")
				if uploadFile == nil {
					r.Response.WriteJson(map[string]interface{}{
						"code":    -1,
						"message": "no file",
					})
					return
				}
				test := r.GetForm(testKey)
				r.Response.WriteJson(map[string]interface{}{
					"code":    0,
					"message": "ok",
					testKey:   test,
				})
				//fmt.Println(uploadFile.Filename, uploadFile.Size)
			})
			go func() {
				svr.Run()
			}()
			defer func() {
				_ = svr.Shutdown()
			}()
			time.Sleep(1 * time.Second)

			gotBody, writer, err := CreateFormBody(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFormBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//fmt.Println(gotBody)
			url := "http://127.0.0.1:47632/test"
			response, err := http.Post(url, writer.FormDataContentType(), gotBody)
			if err != nil {
				t.Error(err)
				return
			}
			defer Close(response.Body)
			cc, _ := io.ReadAll(response.Body)
			type cs struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
				Test    string `json:"test"`
			}
			var csV = cs{}
			if err = json.Unmarshal(cc, &csV); err != nil {
				t.Error(err)
				return
			}

			if csV.Test != testValue {
				t.Error("value err")
				return
			}
		})
	}
}
