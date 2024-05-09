package tools

import "testing"

func TestIsNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "123",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s: "abc",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s: "1bb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumber(tt.args.s); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
