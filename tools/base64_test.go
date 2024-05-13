package tools

import "testing"

func TestBase64StdEncoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "test",
			},
			want: "dGVzdA==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64StdEncoding(tt.args.s); got != tt.want {
				t.Errorf("Base64StdEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64StdDecoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "dGVzdA==",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64StdDecoding(tt.args.s); got != tt.want {
				t.Errorf("Base64StdDecoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64UrlEncoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "test+/?",
			},
			want: "dGVzdCsvPw==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64UrlEncoding(tt.args.s); got != tt.want {
				t.Errorf("Base64UrlEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64UrlDecoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "dGVzdCsvPw==",
			},
			want: "test+/?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64UrlDecoding(tt.args.s); got != tt.want {
				t.Errorf("Base64UrlDecoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64RawStdEncoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "test",
			},
			want: "dGVzdA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64RawStdEncoding(tt.args.s); got != tt.want {
				t.Errorf("Base64RawStdEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64RawStdDecoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "dGVzdA",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64RawStdDecoding(tt.args.s); got != tt.want {
				t.Errorf("Base64RawStdDecoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64RawURLEncoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "test",
			},
			want: "dGVzdA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64RawURLEncoding(tt.args.s); got != tt.want {
				t.Errorf("Base64RawURLEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64RawURLDecoding(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "dGVzdA",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64RawURLDecoding(tt.args.s); got != tt.want {
				t.Errorf("Base64RawURLDecoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
