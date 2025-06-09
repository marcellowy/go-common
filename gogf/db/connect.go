// Package db
package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/marcellowy/go-common/gogf/config"
	"github.com/marcellowy/go-common/gogf/vconfig"
	"github.com/marcellowy/go-common/gogf/vlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// NewConnect establishes a new database connection based on the provided context and key.
//
// Parameters:
// - ctx: the context used for the database connection.
// - key: the key to retrieve database connection configuration details.
//
// Returns a *gorm.DB representing the established database connection.
// Deprecate: use `vdb.NewConnect` instead
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
		vconfig.Get(key+".user").String(),
		vconfig.Get(key+".pwd").String(),
		vconfig.Get(key+".host").String(),
		vconfig.Get(key+".port").Int(),
		vconfig.Get(key+".dbName").String(),
		charset,
		parseTime,
	)
	var (
		err error
		sDB *sql.DB
	)

	vlog.Debug(ctx, dsn)
	var gConfig = &gorm.Config{}
	if vconfig.Get(key+".debug").Bool() == true {
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

	sDB.SetMaxOpenConns(vconfig.Get(key + ".maxOpenConn").Int())
	var maxIdleConn = vconfig.Get(key + ".maxIdleConn").Int()
	if maxIdleConn > 0 {
		sDB.SetMaxIdleConns(maxIdleConn)
	}

	var connMaxLifetime = vconfig.Get(key + ".connMaxLifetime").Int()
	if connMaxLifetime > 0 {
		sDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	}

	return db
}
