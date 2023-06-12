// Package db
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/marcellowy/go-common/config"
	"github.com/marcellowy/go-common/log"
	"github.com/marcellowy/go-common/tools"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Connector *gorm.DB // default database
)

// Config 连接参数
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// dbConf database 配置内容
type dbConf map[string]Config

// autoConnectInstance 自动连接保存的实例
var autoConnectInstance = make(map[string]*gorm.DB)

// InitConnector init
func InitConnector(db *gorm.DB) {
	Connector = db
}

// GetInstance 获取初始化后的实例
func GetInstance(key string) *gorm.DB {
	if v, ok := autoConnectInstance[key]; ok {
		return v
	}
	return nil
}

func Connect(ctx context.Context, c Config) *gorm.DB {

	var (
		db  *gorm.DB
		err error
	)

	gc := gorm.Config{}
	if l, e := zapcore.ParseLevel(config.GetString("logger.level")); e == nil && l < zapcore.ErrorLevel {
		gc.Logger = log.NewGormLog()
	}

	// Connect db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)
	if db, err = gorm.Open(mysql.Open(dsn), &gc); err != nil {
		panic(err)
	}

	var s *sql.DB
	if s, err = db.DB(); err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}
	return db
}

func connectWithMap(ctx context.Context, cc dbConf) {
	for key, cconf := range cc {
		if _, ok := autoConnectInstance[key]; !ok {
			autoConnectInstance[key] = Connect(ctx, cconf)
		}
		if key == "default" {
			Connector = autoConnectInstance[key]
		}
	}
}

// init 初始化所有DB
func init() {
	var (
		err error
		cc  = make(dbConf)
		b   = config.GetStringMap("database")
		ctx = context.Background()
	)

	if err = json.Unmarshal(tools.JSONMarshalByte(b), &cc); err != nil {
		log.Error(ctx, err)
		return
	}

	if len(cc) == 0 {
		// 没有配置
		return
	}

	connectWithMap(ctx, cc)
}
