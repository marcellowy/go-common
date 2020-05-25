package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// 打开一个Native连接
//
func Connect(conn **sql.DB, config *Config) (err error) {

	if len(config.Charset) == 0 {
		config.Charset = "UTF8MB4"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?autocommit=true&charset=%s",
		config.Username, config.Password, config.Host, config.Port, config.DbName, config.Charset)

	//fmt.Printf("database connection string: %s\n", dsn)

	// 打开数据库驱动
	if *conn, err = sql.Open("mysql", dsn); err != nil {
		//fmt.Printf("failed to connect database: %v\n", err)
		return
	}

	// 参数设置
	(*conn).SetMaxOpenConns(int(config.ConnectionSize))
	(*conn).SetMaxIdleConns(int(config.MaxIdleConnection))
	(*conn).SetConnMaxLifetime(config.ConnectionLifetime)

	// 连接测试
	if err = (*conn).Ping(); err != nil {
		//fmt.Printf("failed to ping connection: %v\n", err)
		return
	}
	return
}

// 关闭Native连接
//
func Close(conn **sql.DB) error {
	return (*conn).Close()
}

// 通过gorm连接db
//
func ConnectViaGorm(db **gorm.DB, config *Config) (err error) {
	if len(config.Charset) == 0 {
		config.Charset = "UTF8MB4"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&autocommit=true&charset=%s",
		config.Username, config.Password, config.Host, config.Port, config.DbName, config.Charset)

	if *db, err = gorm.Open("mysql", dsn); err != nil {
		return
	}

	(*db).DB().SetMaxOpenConns(int(config.ConnectionSize))
	(*db).DB().SetMaxIdleConns(int(config.MaxIdleConnection))
	(*db).DB().SetConnMaxLifetime(config.ConnectionLifetime)

	return
}

// 关闭gorm连接
//
func CloseViaGorm(db **gorm.DB) {
	if (*db) != nil {
		(*db).Close()
	}
}

func TransactionWithGorm(db *gorm.DB, cb func(tx *gorm.DB) error) error {

	var err error
	var tx *gorm.DB

	// 打开事务
	if tx = db.Begin(); tx.Error != nil {
		return tx.Error
	}
	defer tx.RollbackUnlessCommitted()

	// 执行逻辑
	if err = cb(tx); err != nil {
		return err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
