// Package db
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/marcellowy/go-common/gogf/config"
	"github.com/marcellowy/go-common/gogf/vlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// NewConnect 数据库实例
func NewConnect(ctx context.Context, key string) (db *gorm.DB) {

	var (
		charset      = "utf8mb4"
		newCharset   = config.Get(key + ".charset")
		parseTime    = "True"
		newParseTime = config.Get(key + ".parseTime")
	)

	if !newCharset.IsEmpty() {
		charset = newCharset.String()
	}

	if !newParseTime.Bool() {
		parseTime = "False"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=Local",
		config.Get(key+".user").String(),
		config.Get(key+".pwd").String(),
		config.Get(key+".host").String(),
		config.Get(key+".port").Int(),
		config.Get(key+".dbName").String(),
		charset,
		parseTime,
	)
	var (
		err error
		sDB *sql.DB
	)

	vlog.Debug(ctx, dsn)

	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		vlog.Error(ctx, err)
		return
	}

	if sDB, err = db.DB(); err != nil {
		vlog.Error(ctx, err)
		return
	}

	sDB.SetMaxOpenConns(config.Get(key + ".maxOpenConn").Int())
	var maxIdleConn = config.Get(key + ".maxIdleConn").Int()
	if maxIdleConn > 0 {
		sDB.SetMaxIdleConns(maxIdleConn)
	}

	var connMaxLifetime = config.Get(key + ".connMaxLifetime").Int()
	if connMaxLifetime > 0 {
		sDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	}

	return db
}
