// Package db
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
		charset          = "utf8mb4"
		newCharset       = config.Get(key + ".charset")
		parseTime        = "True"
		disableParseTime = config.Get(key + ".disableParseTime")
	)

	if !newCharset.IsEmpty() {
		charset = newCharset.String()
	}

	if disableParseTime.Bool() {
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
	var gConfig = &gorm.Config{}
	if config.Get(key+".debug").Bool() == true {
		vlog.Debug(ctx, "gorm debug open")
		//gConfig.Logger = logger.Default.LogMode(logger.Info)
		gConfig.Logger = NewGormLog()
	}

	if db, err = gorm.Open(mysql.Open(dsn), gConfig); err != nil {
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
