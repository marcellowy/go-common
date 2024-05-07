// Package tools
package tools

import (
	"fmt"
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
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				dir: "./test_recreate_directory",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReCreateDirectory(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("ReCreateDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
