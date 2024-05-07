// Package tools
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
