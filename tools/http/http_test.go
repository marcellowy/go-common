package http

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/marcellowy/go-common/tools"
	"io"
	"net/http"
	"os"
	"strings"
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

			svr.BindHandler("/test_post", func(r *ghttp.Request) {
				name := r.GetRequest("name", "").String()
				r.Response.WriteJson(map[string]interface{}{
					"code":    0,
					"message": "ok",
					"name":    name,
				})
			})

			svr.BindHandler("/test_get", func(r *ghttp.Request) {
				if r.Request.Method == http.MethodGet {
					name := r.GetRequest("name", "").String()
					r.Response.WriteJson(map[string]interface{}{
						"code":    0,
						"message": "ok",
						"name":    name,
					})
				}
			})

			go func() {
				svr.Run()
			}()
			defer func() {
				_ = svr.Shutdown()
			}()
			time.Sleep(1 * time.Second)

			{
				//fmt.Println(gotBody)
				url := "http://127.0.0.1:47632/test"
				client := NewHttpClient()
				response, err := client.PostForm(gctx.New(), url, tt.args.data)
				if err != nil {
					t.Error(err)
					return
				}
				defer tools.Close(response.Body)
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
			}

			{
				//fmt.Println(gotBody)
				url := "http://127.0.0.1:47632/test_post"
				client := NewHttpClient()
				response, err := client.Post(gctx.New(), url, strings.NewReader(`{"name":"123"}`))
				if err != nil {
					t.Error(err)
					return
				}
				defer tools.Close(response.Body)
				cc, _ := io.ReadAll(response.Body)
				type cs struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
					Name    string `json:"name"`
				}
				var csV = cs{}
				if err = json.Unmarshal(cc, &csV); err != nil {
					t.Error(err)
					return
				}

				if csV.Name != "123" {
					t.Error("value err")
					return
				}
			}

			{
				//fmt.Println(gotBody)
				url := "http://127.0.0.1:47632/test_post?name=123"
				client := NewHttpClient()
				response, err := client.Get(gctx.New(), url)
				if err != nil {
					t.Error(err)
					return
				}
				defer tools.Close(response.Body)
				cc, _ := io.ReadAll(response.Body)
				type cs struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
					Name    string `json:"name"`
				}
				var csV = cs{}
				if err = json.Unmarshal(cc, &csV); err != nil {
					t.Error(err)
					return
				}

				if csV.Name != "123" {
					t.Error("value err")
					return
				}
			}
		})
	}
}
