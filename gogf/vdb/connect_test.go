package vdb

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"os"
	"reflect"
	"testing"
	"time"
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

func TestNewQuickConnect(t *testing.T) {
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
		// TODO: Add test cases.
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
			if gotDb := NewQuickConnect(tt.args.ctx, tt.args.key); !reflect.DeepEqual(gotDb, tt.wantDb) {
				t.Errorf("NewQuickConnect() = %v, want %v", gotDb, tt.wantDb)
			}
		})
	}
}

func TestConfig_Hash(t *testing.T) {
	type fields struct {
		User            string
		Password        string
		Host            string
		Port            int
		Schema          string
		Charset         string
		ParseTime       bool
		MaxOpenConn     int
		MaxIdleConn     int
		ConnMaxLifeTime time.Duration
		Debug           bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				User:            "root",
				Password:        "123456",
				Host:            "127.0.0.1",
				Port:            3306,
				Schema:          "test123",
				Charset:         "utf8",
				ParseTime:       true,
				MaxOpenConn:     10,
				MaxIdleConn:     10,
				ConnMaxLifeTime: 10 * time.Second,
				Debug:           true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				User:            tt.fields.User,
				Password:        tt.fields.Password,
				Host:            tt.fields.Host,
				Port:            tt.fields.Port,
				Schema:          tt.fields.Schema,
				Charset:         tt.fields.Charset,
				ParseTime:       tt.fields.ParseTime,
				MaxOpenConn:     tt.fields.MaxOpenConn,
				MaxIdleConn:     tt.fields.MaxIdleConn,
				ConnMaxLifeTime: tt.fields.ConnMaxLifeTime,
				Debug:           tt.fields.Debug,
			}
			got, err := c.Hash()
			fmt.Println(got, err)
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			//if got != tt.want {
			//	t.Errorf("Hash() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
