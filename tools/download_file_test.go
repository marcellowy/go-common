package tools

import (
	"context"
	"testing"
)

func TestDownloadFileFromUrl(t *testing.T) {
	t.SkipNow()
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantN   int64
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
				url: "http://10.86.4.111:9096/warehouse/ls/8.8.47.windows.0/lightStreamer.zip",
			},
			wantN: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			d := &DownloadFile{}
			d.SaveAsDirectory("./c")
			d.SaveAsFilename("111.zip")

			gotN, err := d.DownloadFromUrl(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFileFromUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotN == 0 {
				t.Errorf("DownloadFileFromUrl() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
