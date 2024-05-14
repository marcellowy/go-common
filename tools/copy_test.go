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

	defer func() {
		_ = os.RemoveAll("test_copy_file")
	}()

}

func TestCopyDir(t *testing.T) {
	type args struct {
		dst string
		src []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				dst: "./test_copy_dir",
				src: []string{"./test_copy_dir_1"},
			},
		},
		{
			name: "test",
			args: args{
				dst: "./test_copy_dir",
				src: []string{"./md5_test_file.txt"},
			},
		},
	}

	_ = os.RemoveAll("./test_copy_dir")
	_ = os.MkdirAll("./test_copy_dir", os.ModePerm)

	_ = os.RemoveAll("./test_copy_dir_1")
	{
		_ = os.MkdirAll("./test_copy_dir_1/test", os.ModePerm)
		_ = os.WriteFile("./test_copy_dir_1/test/test1.txt", []byte("test1"), os.ModePerm)
		_ = os.WriteFile("./test_copy_dir_1/test/test2.txt", []byte("test2"), os.ModePerm)
		_ = os.WriteFile("./test_copy_dir_1/test/test3.txt", []byte("test3"), os.ModePerm)
		_ = os.WriteFile("./test_copy_dir_1/test/test4.txt", []byte("test4"), os.ModePerm)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.dst, tt.args.src...); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	defer func() {
		_ = os.RemoveAll("./test_copy_dir")
		_ = os.RemoveAll("./test_copy_dir_1")
	}()
}
