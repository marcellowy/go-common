package vdb

import (
	"context"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestNewConnect(t *testing.T) {
	if os.Getenv("TEST_DB_CONNECT") != "1" {
		t.SkipNow()
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name   string
		args   args
		wantDb *gorm.DB
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
				key: "database",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDb := NewConnect(tt.args.ctx, tt.args.key); gotDb != nil {
				t.Errorf("NewConnect() = %v, want %v", gotDb, tt.wantDb)
			}
		})
	}
}
