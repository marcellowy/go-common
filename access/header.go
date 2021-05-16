package access

import (
	"context"
	"encoding/json"
)

const key = "common_header"

// Header 公共头
type Header struct {
	TraceID string `json:"trace_id"`
	Version int    `json:"version"` // 头版本,主要用于序列化
}

func (h *Header) Marshal() ([]byte, error) {
	return json.Marshal(h)
}

func (h *Header) Unmarshal(data []byte) error {
	return json.Unmarshal(data, h)
}

func WriteHeader(ctx *context.Context, header *Header) {
	val, err := header.Marshal()
	if err != nil {
		return
	}
	*ctx = context.WithValue(*ctx, key, val)
}

func ReadHeader(ctx context.Context) *Header {
	data := ctx.Value(key).([]byte)
	var h = &Header{}
	_ = h.Unmarshal(data)
	return h
}
