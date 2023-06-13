// Package router
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package router

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/marcellowy/go-common/merr"
	"reflect"
	"strings"
)

const defaultRouterMetaName = "Meta"
const defaultMimeJSON = "application/json"
const defaultMimeXML = "application/xml"

// supportMethod support method
var supportMethod = map[string]string{
	"get":     "GET",
	"post":    "POST",
	"delete":  "DELETE",
	"patch":   "PATCH",
	"put":     "PUT",
	"options": "OPTIONS",
	"head":    "HEAD",
}

// routerCache METHOD PATH cache one function
var routerCache map[string]map[string]cache

var responseMarshal Response

// cache object
type cache struct {
	controllerInstance reflect.Value
	method             reflect.Value
	request            reflect.Type // request parameter exclude *gin.Context
	response           reflect.Type // response parameter exclude error
	requestMeta        reflect.StructField
	responseMeta       reflect.StructField
}

type Meta struct {
}

type Error interface {
	ToJSON() []byte
}

type Response interface {
	// Marshal input is response struct
	Marshal(interface{}) interface{}
}

func init() {
	routerCache = make(map[string]map[string]cache)
}

func ResponseMarshal(response Response) {
	responseMarshal = response
}

// Register router register
func Register(engine gin.IRoutes, v ...interface{}) {
	var ok bool
	for _, e := range v {
		var controllerValue = reflect.ValueOf(e)
		for i := 0; i < controllerValue.NumMethod(); i++ {
			if controllerValue.Method(i).Type().NumIn() != 2 {
				panic("input parameter length must be 2 for " + controllerValue.Method(i).String())
			}
			if controllerValue.Method(i).Type().NumOut() != 2 {
				panic("output parameter length must be 2" + controllerValue.Method(i).String())
			}
			// check parameter type
			//firstInParameter := controllerValue.Method(i).Type().In(0)
			secondInParameter := controllerValue.Method(i).Type().In(1)
			firstOutParameter := controllerValue.Method(i).Type().Out(0)
			//secondOutParameter := controllerValue.Method(i).Type().Out(1)

			if secondInParameter.Kind() != reflect.Ptr {
				panic("second input parameter must ptr")
			}

			if firstOutParameter.Kind() != reflect.Ptr {
				panic("first output parameter must ptr")
			}

			var (
				requestField, responseField reflect.StructField
				sMethod                     = "GET"
			)

			if requestField, ok = secondInParameter.Elem().FieldByName(defaultRouterMetaName); !ok {
				panic("Request struct no" + defaultRouterMetaName)
			}

			// parse path and method
			// the path and method must unique by your controllers
			requestPath := requestField.Tag.Get("path")
			requestMethod := requestField.Tag.Get("method")
			if sMethod, ok = supportMethod[requestMethod]; !ok {
				panic("method " + requestMethod + " not support")
			}

			if _, ok = routerCache[requestPath][requestMethod]; !ok {
				routerCache[requestPath] = make(map[string]cache)
			}

			// write cache
			// used by request happen
			routerCache[requestPath][requestMethod] = cache{
				controllerInstance: controllerValue,
				method:             controllerValue.Method(i),
				request:            secondInParameter,
				response:           firstOutParameter,
				requestMeta:        requestField,
				responseMeta:       responseField,
			}
			// register router to gin framework
			// fmt.Println(sMethod, requestPath)
			reflect.ValueOf(engine).MethodByName(sMethod).Call([]reflect.Value{
				reflect.ValueOf(requestPath),
				reflect.ValueOf(router),
			})
		}
	}
}

// router convert request parameter to structs function
func router(ctx *gin.Context) {
	var (
		err    error
		method = strings.ToLower(ctx.Request.Method)
		path   = ctx.Request.URL.Path
		b      []byte // response body
		ok     bool
	)

	if _, ok = routerCache[path]; !ok {
		// unregister path
		return
	}
	if _, ok = routerCache[path][method]; !ok {
		// unregister method
		return
	}
	c := routerCache[path][method]

	request := reflect.New(c.request.Elem()).Interface()
	if err = ctx.ShouldBind(request); err != nil {

	}
	response := c.method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(request),
	})

	if !response[1].IsNil() {
		var merror Error
		if merror, ok = response[1].Interface().(Error); ok {
			if _, err = ctx.Writer.Write(merror.ToJSON()); err != nil {
				panic(err)
			}
			return
		}
		// no support type
		return
	}

	if ctx.Writer.Size() > 0 {
		// write body in function
		return
	}

	mime := c.responseMeta.Tag.Get("mime")
	var body interface{}
	if responseMarshal == nil {
		body = merr.NewCode(0, "success", response[0].Elem().Interface())
	} else {
		body = responseMarshal.Marshal(response[0].Elem().Interface())
	}

	switch mime {
	case defaultMimeJSON:
		if b, err = json.Marshal(body); err != nil {
			panic(err)
		}
	case defaultMimeXML:
		if b, err = xml.Marshal(body); err != nil {
			panic(err)
		}
	default:
		if b, err = json.Marshal(body); err != nil {
			panic(err)
		}
	}

	if _, err = ctx.Writer.Write(b); err != nil {
		panic(err)
	}
}
