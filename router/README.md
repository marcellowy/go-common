## Router Example：

```go
// router/router.go
package router

import (
	"your/controller/path/index" // this your controller path
	"github.com/gin-gonic/gin"
	"github.com/marcellowy/go-common/router" // this router
)

// SetRouter call by main.go
func SetRouter(engine *gin.Engine) {
	
	// no group register
	router.Register(engine, &index.Controller{})

	v1 := engine.Group("/cc")
	{
		// group register
		router.Register(v1, &index.Controller{})
	}
}

// output:
// [GIN-debug] POST   /nginx                    --> github.com/marcellowy/go-common/router.router (2 handlers)
// [GIN-debug] GET    /1/2/3/4                  --> github.com/marcellowy/go-common/router.router (2 handlers)
// [GIN-debug] POST   /cc/nginx                 --> github.com/marcellowy/go-common/router.router (2 handlers)
// [GIN-debug] GET    /cc/1/2/3/4               --> github.com/marcellowy/go-common/router.router (2 handlers)
```
## Controller Exmaple
```go
// controller 
type Controller struct {
}

type TextReq struct {
	router.Meta `path:"/1/2/3/4" method:"get"`  // register to gin framework GET("/1/2/3/4", Controller.Text)
	Name        string `json:"name" form:"name"`
	Age         int
}

type TextRes struct {
	router.Meta `mime:"application/json"`   // response json formatter, and support application/xml
	MyName      string
}

type NginxReq struct {
	router.Meta `path:"/nginx" method:"post"` // register to gin framework POST("/nginx", Controller.Nginx)
	Name        string `json:"name" form:"name"`
	Age         int
}

type NginxRes struct {
	Number int `json:"number"`  // response default json formatter
}

func (*Controller) Text(ctx context.Context, req *TextReq) (*TextRes, error) {
	var res = &TextRes{
		MyName: "123",
	}
	res.MyName = "Server say: your " + req.Name
	return res, nil
}

// like grpc invoke
// we can use ctx.(*gin.Context) read *gin.Context Object
func (*Controller) Nginx(ctx context.Context, req *NginxReq) (*NginxRes, error) {
	var res = &NginxRes{
		Number: 999,
	}
	return res, nil
}

```

it's read controller instance and parse method, parameter and they type.