// Package tools
package tools

import (
	"testing"
)

func TestRemoveSameFromStringSlice(t *testing.T) {
	RemoveSameFromStringSlice([]string{"1", "1", "2", "2", "3"})
}

func TestSliceTrimSame(t *testing.T) {
	SliceTrimSame([]string{"1", "1", "2", "2", "3", "4", "4"})
	SliceTrimSame([]string{})
}

func TestSliceRemove(t *testing.T) {
	SliceRemove([]string{"1"}, []string{"1", "1", "2", "2", "3", "4", "4"})
}

func TestInSlice(t *testing.T) {
	type args[T comparable] struct {
		need T
		arr  []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "test",
			args: args[string]{
				need: "1",
				arr:  []string{"1", "1", "2", "2", "3", "4", "4"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.need, tt.args.arr); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
