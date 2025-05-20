package tools

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				d: 345 * time.Second,
			},
			want: "5 min 45 s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.args.d); got != tt.want {
				t.Errorf("FormatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatBytes(t *testing.T) {
	type args struct {
		bytes uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				bytes: 1024,
			},
			want: "1.00 KB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBytes(tt.args.bytes); got != tt.want {
				t.Errorf("FormatBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSizeUnit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		wantByte int64
		wantErr  bool
	}{
		{
			name:     "1",
			args:     args{s: "10M"},
			wantByte: 1024 * 1024 * 10,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotByte, err := ParseSizeUnit(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSizeUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotByte != tt.wantByte {
				t.Errorf("ParseSizeUnit() gotByte = %v, want %v", gotByte, tt.wantByte)
			}
		})
	}
}
