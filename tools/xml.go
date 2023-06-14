// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

import "encoding/xml"

// XMLMarshalString ignore error，return xml.Marshal string
func XMLMarshalString(v interface{}) string {
	var b, err = xml.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// XMLMarshalByte ignore error, return xml.Marshal []byte()
func XMLMarshalByte(v interface{}) []byte {
	var b, err = xml.Marshal(v)
	if err != nil {
		return []byte("")
	}
	return b
}
