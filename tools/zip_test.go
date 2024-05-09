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
		{
			name: "test",
			args: args{
				zipPath: "test_zip_file/test.zip",
				paths:   []string{"test_zip_file"},
			},
			wantErr: false,
		},
	}

	_ = os.RemoveAll("test_zip_file")

	_ = os.MkdirAll("test_zip_file", os.ModePerm)
	_ = os.WriteFile("test_zip_file/test.txt", []byte("test"), os.ModePerm)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Zip(tt.args.zipPath, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnzip(t *testing.T) {

	{
		type args struct {
			zipPath string
			paths   []string
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "test",
				args: args{
					zipPath: "test_unzip_file/test.zip",
					paths:   []string{"test_unzip_file"},
				},
				wantErr: false,
			},
		}

		_ = os.RemoveAll("test.zip")
		_ = os.RemoveAll("test_unzip_file")

		_ = os.MkdirAll("test_unzip_file", os.ModePerm)
		_ = os.WriteFile("test_unzip_file/test.txt", []byte("test"), os.ModePerm)

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := Zip(tt.args.zipPath, tt.args.paths...); (err != nil) != tt.wantErr {
					t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}

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
		{
			name: "test",
			args: args{
				ctx:      context.Background(),
				filename: "test_unzip_file/test.zip",
				saveAs:   "test_unzip_file/unzip/test",
				handler:  nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unzip(tt.args.ctx, tt.args.filename, tt.args.saveAs, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Unzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
