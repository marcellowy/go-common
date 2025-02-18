// Package tools
package tools

import (
	"fmt"
	"os"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T) {
	dir := GetCurrentDirectory()
	if dir == "" {
		t.Error("GetCurrentDirectory err")
		return
	}
}

func TestPathExists(t *testing.T) {
	ok, err := PathExists(".")
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("path not exists")
		return
	}

	ok, err = PathExists("/aaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		t.Error(err)
		return
	}

	if ok {
		// directory not exists and ok is false it's true
		t.Error(err)
		return
	}

}

func TestRemoveLastSeparator(t *testing.T) {
	fmt.Println(RemoveLastSeparator(""))
	fmt.Println(RemoveLastSeparator("a/b/c"))
	fmt.Println(RemoveLastSeparator("a/b/c/"))
	fmt.Println(RemoveLastSeparator("\\a\\b\\"))
}

func TestReCreateDirectory(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				dir: "./test_recreate_directory",
			},
			wantErr: false,
		},
	}

	defer func() {
		_ = os.RemoveAll("./test_recreate_directory")
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReCreateDirectory(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("ReCreateDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateEmptyFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				filename: "./test_create_empty_file/test_create_empty_file",
			},
		},
		{
			name: "test",
			args: args{
				filename: "./test_create_empty_file/test_create_empty_file.txt",
			},
		},
		{
			name: "test",
			args: args{
				filename: "./test_create_empty_file/test_create_empty_file/",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateEmptyFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("CreateEmptyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	if err := os.RemoveAll("./test_create_empty_file"); err != nil {
		t.Error(err)
		return
	}
}
