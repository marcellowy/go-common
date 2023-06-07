// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

import (
	"encoding/json"
)

// JSONMarshalString ignore error，return JSON.Marshal string
func JSONMarshalString(v interface{}) string {
	var b, err = json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// JSONMarshalByte ignore error, return JSON.Marshal []byte()
func JSONMarshalByte(v interface{}) []byte {
	var b, err = json.Marshal(v)
	if err != nil {
		return []byte("")
	}
	return b
}
