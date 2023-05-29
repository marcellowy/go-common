// Package vlog
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package vlog

import (
	"database/sql"
	"fmt"
	"gitlab.vrviu.com/inviu/engine/go-common.git/models/gorm/db_lsstat/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormLog(t *testing.T) {
	t.SkipNow()

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"viu@1234",
		"192.168.44.131",
		3306,
		"db_lsstat_dev",
	)

	var (
		DB  *gorm.DB
		err error
	)
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewGormLog(),
	}); err != nil {
		t.Error(err)
		return
	}

	var s *sql.DB
	if s, err = DB.DB(); err != nil {
		t.Error(err)
		return
	}

	if err = s.Ping(); err != nil {
		t.Error(err)
		return
	}

	var table = model.ReportRec{}

	if err = DB.Table("t_report_rec_20230529").First(&table).Error; err != nil {
		t.Error(err)
		return
	}

}
