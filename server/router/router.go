// Package router
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marcellowy/go-common/merr"
	"github.com/marcellowy/go-common/tools"
	"reflect"
	"strings"
)

const defaultRouterMetaName = "Meta"
const defaultMimeJSON = "application/json"
const defaultMimeXML = "application/xml"
const defaultMimeHTML = "text/html"

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
	Marshal(interface{}, error) []byte
}

func init() {
	routerCache = make(map[string]map[string]cache)
}

func ResponseMarshal(response Response) {
	responseMarshal = response
}

// Register router register
func Register(engine gin.IRouter, v ...interface{}) {
	for _, handle := range v {
		register(engine, "", handle)
	}
}

// Group router register
func Group(engine gin.IRouter, relativePath string, handle ...interface{}) {
	g := engine.Group(relativePath)
	for _, v := range handle {
		register(g, relativePath, v)
	}
}

// register
func register(engine gin.IRouter, relativePath string, handle interface{}) {
	var (
		ok              bool
		controllerValue = reflect.ValueOf(handle)
	)

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

		if responseField, ok = firstOutParameter.Elem().FieldByName(defaultRouterMetaName); !ok {

		}

		// parse path and method
		// the path and method must unique by your controllers
		requestPath := requestField.Tag.Get("path")
		cacheRequestPath := relativePath + requestPath
		requestMethod := requestField.Tag.Get("method")
		if sMethod, ok = supportMethod[requestMethod]; !ok {
			panic("method " + requestMethod + " not support")
		}

		if _, ok = routerCache[cacheRequestPath][requestMethod]; !ok {
			routerCache[cacheRequestPath] = make(map[string]cache)
		}

		// write cache
		// used by request happen
		routerCache[cacheRequestPath][requestMethod] = cache{
			controllerInstance: controllerValue,
			method:             controllerValue.Method(i),
			request:            secondInParameter,
			response:           firstOutParameter,
			requestMeta:        requestField,
			responseMeta:       responseField,
		}
		// register router to gin framework
		//fmt.Println(sMethod, cacheRequestPath)
		reflect.ValueOf(engine).MethodByName(sMethod).Call([]reflect.Value{
			reflect.ValueOf(requestPath),
			reflect.ValueOf(router),
		})
	}
}

// router convert request parameter to structs function
func router(ctx *gin.Context) {
	var (
		err    error
		method = strings.ToLower(ctx.Request.Method)
		path   = ctx.Request.URL.Path
		ok     bool
	)

	if _, ok = routerCache[path]; !ok {
		// unregister path
		fmt.Println(routerCache)
		fmt.Println("no cache path:", path)
		return
	}
	if _, ok = routerCache[path][method]; !ok {
		// unregister method
		fmt.Println("no cache path:", path, method)
		return
	}
	c := routerCache[path][method]

	request := reflect.New(c.request.Elem()).Interface()
	if err = ctx.ShouldBind(request); err != nil {
		fmt.Println(err)
	}

	response := c.method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(request),
	})

	mime := c.responseMeta.Tag.Get("mime")
	if !response[1].IsNil() {
		// has error logic process
		writeContent(ctx, mime, response[1].Interface())
		return
	}

	fmt.Println("writer.Size()", ctx.Writer.Size())
	if ctx.Writer.Size() > 0 {
		// write body in function
		return
	}
	writeContent(ctx, mime, response[0].Elem().Interface())
}

func writeErrorContent(ctx *gin.Context, mime string, response interface{}) {

	var (
		b      []byte
		err    error
		ok     bool
		merror *merr.Error
	)

	if err, ok = response.(error); ok {
		b = mimeContent(mime, merr.NewCode(-1, err.Error()))
	} else if merror, ok = response.(*merr.Error); ok {
		b = mimeContent(mime, merror)
	} else {
		panic("no support error type")
	}

	if len(b) > 0 {
		if _, err = ctx.Writer.Write(b); err != nil {
			panic(err)
		}
	}
}

// writeContent
func writeContent(ctx *gin.Context, mime string, response interface{}) {
	var (
		b   []byte
		err error
		ok  bool
	)

	if err, ok = response.(error); ok {
		// error
		if responseMarshal != nil {
			// define marshal
			if _, err = ctx.Writer.Write(responseMarshal.Marshal(response, err)); err != nil {
				panic(err)
			}
			return
		}
		writeErrorContent(ctx, mime, response)
		return
	}

	if responseMarshal != nil {
		// define marshal
		if _, err = ctx.Writer.Write(responseMarshal.Marshal(response, nil)); err != nil {
			panic(err)
		}
		return
	}

	b = mimeContent(mime, merr.NewCode(0, "success", response))

	if len(b) > 0 {
		if _, err = ctx.Writer.Write(b); err != nil {
			panic(err)
		}
	}
}

// mimeContent encode content
func mimeContent(mime string, v interface{}) []byte {
	var b []byte
	switch mime {
	case defaultMimeJSON:
		b = tools.JSONMarshalByte(v)
	case defaultMimeXML:
		b = tools.XMLMarshalByte(v)
	default:
		// if unknown
	}
	return b
}
