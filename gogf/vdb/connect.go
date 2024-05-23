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

// Config connect config
type Config struct {
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

const (
	// DefaultMaxIdleConn default max idle conns
	DefaultMaxIdleConn = 1
	DefaultMaxOpenConn = 10
)

// NewConnectWithConfig via config connect new
func NewConnectWithConfig(ctx context.Context, config *Config) (db *gorm.DB) {

	var parseTime = "True"
	if !config.ParseTime {
		parseTime = "False"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Schema, config.Charset, parseTime,
	)
	var (
		err error
		sDB *sql.DB
	)

	vlog.Debug(ctx, dsn)
	var gConfig = &gorm.Config{}
	if config.Debug {
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

	if config.MaxOpenConn > 0 {
		sDB.SetMaxOpenConns(config.MaxOpenConn)
	} else {
		sDB.SetMaxOpenConns(DefaultMaxOpenConn)
	}

	if config.MaxIdleConn > 0 {
		sDB.SetMaxIdleConns(config.MaxIdleConn)
	} else {
		sDB.SetMaxIdleConns(DefaultMaxIdleConn)
	}

	if config.ConnMaxLifeTime > 0 {
		sDB.SetConnMaxLifetime(config.ConnMaxLifeTime)
	}

	return
}

// NewQuickConnect via framework config connect new
func NewQuickConnect(ctx context.Context, key string) (db *gorm.DB) {
	var (
		config = &Config{
			User:        vconfig.Get(key + ".user").String(),
			Password:    vconfig.Get(key + ".password").String(),
			Host:        vconfig.Get(key + ".host").String(),
			Port:        vconfig.Get(key + ".port").Int(),
			Schema:      vconfig.Get(key + ".schema").String(),
			Charset:     vconfig.Get(key + ".charset").String(),
			MaxOpenConn: vconfig.Get(key + ".maxOpenConn").Int(),
			MaxIdleConn: vconfig.Get(key + ".maxIdleConn").Int(),
			Debug:       vconfig.Get(key + ".debug").Bool(),
		}
	)

	disableParseTime := vconfig.Get(key + ".disableParseTime").Bool()
	connMaxLifeTime := vconfig.Get(key + ".connMaxLifetime").Int()

	if config.Charset == "" {
		config.Charset = "utf8"
	}

	if disableParseTime {
		config.ParseTime = false
	}

	if connMaxLifeTime > 0 {
		config.ConnMaxLifeTime = time.Duration(connMaxLifeTime) * time.Second
	}

	return NewConnectWithConfig(ctx, config)
}

// NewConnect establishes a new database connection based on the provided context and key.
//
// Parameters:
// - ctx: the context used for the database connection.
// - key: the key to retrieve database connection configuration details.
//
// Returns a *gorm.DB representing the established database connection.
// Deprecated: use NewQuickConnect instead.
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

	return
}
