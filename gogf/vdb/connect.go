// Package vdb
package vdb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/marcellowy/go-common/gogf/vconfig"
	"github.com/marcellowy/go-common/gogf/vlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewConnect establishes a new database connection based on the provided context and key.
//
// Parameters:
// - ctx: the context used for the database connection.
// - key: the key to retrieve database connection configuration details.
//
// Returns a *gorm.DB representing the established database connection.
func NewConnect(ctx context.Context, key string) (db *gorm.DB) {

	var (
		user             = vconfig.Get(key + ".user").String()
		password         = vconfig.Get(key + ".password").String()
		host             = vconfig.Get(key + ".host").String()
		port             = vconfig.Get(key + ".port").Int()
		schema           = vconfig.Get(key + ".schema").String()
		charset          = "utf8mb4"
		newCharset       = vconfig.Get(key + ".charset")
		parseTime        = "True"
		disableParseTime = vconfig.Get(key + ".disableParseTime")
		maxOpenConn      = vconfig.Get(key + ".maxOpenConn").Int()
		maxIdleConn      = vconfig.Get(key + ".maxIdleConn").Int()
		connMaxLifetime  = vconfig.Get(key + ".connMaxLifetime").Int()
		debug            = vconfig.Get(key + ".debug").Bool()
	)

	if !newCharset.IsEmpty() {
		charset = newCharset.String()
	}

	if disableParseTime.Bool() {
		parseTime = "False"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=Local",
		user, password, host, port, schema, charset, parseTime,
	)
	var (
		err error
		sDB *sql.DB
	)

	vlog.Debug(ctx, dsn)
	var gConfig = &gorm.Config{}
	if debug {
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

	sDB.SetMaxOpenConns(maxOpenConn)
	if maxIdleConn > 0 {
		sDB.SetMaxIdleConns(maxIdleConn)
	} else {
		sDB.SetMaxIdleConns(1)
	}

	if connMaxLifetime > 0 {
		sDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	}

	return db
}
