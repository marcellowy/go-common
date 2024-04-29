package tools

import (
	"archive/zip"
	"context"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	type args struct {
		zipPath string
		paths   []string
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
				zipPath: "test.zip",
				paths:   []string{"test_copy_file"},
			},
			wantErr: false,
		},
		{
			name: "test",
			args: args{
				zipPath: "test1.zip",
				paths:   []string{"E:\\project09\\go-common\\tools"},
			},
			wantErr: false,
		},
	}

	_ = os.RemoveAll("test.zip")
	_ = os.RemoveAll("test1.zip")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Zip(tt.args.zipPath, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnzip(t *testing.T) {
	type args struct {
		ctx      context.Context
		filename string
		saveAs   string
		handler  func(file **zip.File)
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
				ctx:      context.Background(),
				filename: "./test_copy_file/light-streamer-win-8.14.153.2-94e391d10e875b419882e497da9db21f.zip",
				saveAs:   "./test_copy_file/aaa",
				handler:  nil,
			},
			wantErr: false,
		},
		{
			name: "test",
			args: args{
				ctx:      context.Background(),
				filename: "./test_copy_file/light-streamer-win-8.14.153.2-94e391d10e875b419882e497da9db21f.zip",
				saveAs:   "./test_copy_file/bbb",
				handler:  UnzipGBKHandler,
			},
			wantErr: false,
		},
	}

	_ = os.RemoveAll("./test_copy_file/aaa")
	_ = os.RemoveAll("./test_copy_file/bbb")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unzip(tt.args.ctx, tt.args.filename, tt.args.saveAs, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Unzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
