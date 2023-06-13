// Package merr
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package merr

import "github.com/marcellowy/go-common/tools"

// New err
func New(message string) *Error {
	return &Error{
		Code:    -1,
		Message: message,
	}
}

// NewCode err and code
func NewCode(code int, message string, data ...interface{}) *Error {
	var retData interface{}
	if len(data) > 0 {
		retData = data[0]
	}
	return &Error{
		Code:    code,
		Message: message,
		Data:    retData,
	}
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) ToJSON() []byte {
	return tools.JSONMarshalByte(e)
}
