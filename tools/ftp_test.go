package tools

import (
	"context"
	"github.com/jlaffaye/ftp"
	"testing"
	"time"
)

func TestWarehouseDownloadFTP_Download(t *testing.T) {
	t.SkipNow()
	type fields struct {
		conn        *ftp.ServerConn
		FTPAddress  string
		FTPUsername string
		FTPPassword string
		FTPPath     string
		FTPTimeout  time.Duration
	}
	type args struct {
		ctx    context.Context
		path   string
		saveAs string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				conn:        nil,
				FTPAddress:  "", // "127.0.0.1:21",
				FTPUsername: "", // "ftp user
				FTPPassword: "", // "ftp password"
				FTPTimeout:  5 * time.Second,
			},
			args: args{
				ctx:    context.Background(),
				path:   "", // get path
				saveAs: "", // save as path
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &FTP{
				conn:        tt.fields.conn,
				FTPAddress:  tt.fields.FTPAddress,
				FTPUsername: tt.fields.FTPUsername,
				FTPPassword: tt.fields.FTPPassword,
				FTPTimeout:  tt.fields.FTPTimeout,
			}
			if err := w.Download(tt.args.ctx, tt.args.path, tt.args.saveAs); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
