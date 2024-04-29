package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	type args struct {
		dst string
		src string
	}

	var dst = "test_copy_file/dst.txt"
	var src = "test_copy_file/source.txt"

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				dst: dst,
				src: src,
			},
			wantErr: false,
		},
	}

	_ = os.RemoveAll("test_copy_file")
	_ = os.MkdirAll(filepath.Dir(src), os.ModePerm)
	_ = os.WriteFile(src, []byte("test"), os.ModePerm)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyFile(tt.args.dst, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		dst string
		src []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				dst: "./test_copy",
				src: []string{"./"},
			},
		},
		{
			name: "test",
			args: args{
				dst: "./test_copy",
				src: []string{"E:/project09/go-common/tools/md5_test_file.txt"},
			},
		},
	}

	_ = os.RemoveAll("./test_copy123")
	_ = os.RemoveAll("./test_copy")
	_ = os.MkdirAll("./test_copy123", os.ModePerm)
	_ = os.MkdirAll("./test_copy", os.ModePerm)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.dst, tt.args.src...); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	defer func() {
		_ = os.RemoveAll("./test_copy123")
		_ = os.RemoveAll("./test_copy")
	}()
}
