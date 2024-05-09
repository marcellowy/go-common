// Package tools
package tools

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReverseByte(t *testing.T) {
	var a = []byte("ab")
	if bytes.Compare(ReverseByte(a), []byte("ba")) != 0 {
		t.Error("err")
	}
}

func TestStringToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: "abc",
			},
			want: []byte("abc"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				b: []byte("abc"),
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(tt.args.b); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: "abc",
			},
			want: "cba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubStr(t *testing.T) {
	type args struct {
		s   string
		pos uint
		end uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s:   "abcdefghijklmn",
				pos: 5,
				end: 5 + 3,
			},
			want: "fgh",
		},
		{
			name: "test",
			args: args{
				s:   "abcdefghijklmn",
				pos: 14,
				end: 100,
			},
			want: "abcdefghijklmn",
		},
		{
			name: "test",
			args: args{
				s:   "abcdefghijklmn",
				pos: 13,
				end: 100,
			},
			want: "n",
		},
		{
			name: "test",
			args: args{
				s:   "abcdefghijklmn",
				pos: 5,
				end: 4,
			},
			want: "",
		},
		{
			name: "test",
			args: args{
				s:   "abcdefghijklmn",
				pos: 5,
				end: 5,
			},
			want: "",
		},
		{
			name: "test",
			args: args{
				s:   "abcdefghijklmn",
				pos: 3,
				end: 7,
			},
			want: "defg",
		},
		{
			name: "test",
			args: args{
				s:   "abcdefghijklmn",
				pos: 0,
				end: 6,
			},
			want: "abcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubString(tt.args.s, tt.args.pos, tt.args.end); got != tt.want {
				t.Errorf("SubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
