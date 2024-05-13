package tools

import "encoding/base64"

// Base64StdEncoding ...
func Base64StdEncoding(s string) string {
	return base64.StdEncoding.EncodeToString(StringToBytes(s))
}

// Base64StdDecoding ...
func Base64StdDecoding(s string) string {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return BytesToString(b)
}

// Base64UrlEncoding ...
func Base64UrlEncoding(s string) string {
	return base64.URLEncoding.EncodeToString(StringToBytes(s))
}

// Base64UrlDecoding ...
func Base64UrlDecoding(s string) string {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return BytesToString(b)
}

// Base64RawStdEncoding ...
func Base64RawStdEncoding(s string) string {
	return base64.RawStdEncoding.EncodeToString(StringToBytes(s))
}

// Base64RawStdDecoding ...
func Base64RawStdDecoding(s string) string {
	b, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return BytesToString(b)
}

// Base64RawURLEncoding ...
func Base64RawURLEncoding(s string) string {
	return base64.RawURLEncoding.EncodeToString(StringToBytes(s))
}

// Base64RawURLDecoding ...
func Base64RawURLDecoding(s string) string {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return BytesToString(b)
}
