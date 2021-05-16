package gorm

// GORM数据库连接器
// Author: chadwang<chadwang@tencent.com>
// Date: 2021/01/01

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
	"sync"

	"gitee.com/marcellos/wyi-common/database/mysql"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// onceMap 连接排重
var (
	onceMap = make(map[string]*gorm.DB)
	lock    = new(sync.RWMutex)
)

func init() {

}

// Option 连接器选项
type Option func(*Connector)

// WithDialector 设置Dialector
func WithDialector(dialector *driver.Dialector) Option {
	return func(connector *Connector) {
		connector.dialector = dialector
	}
}

// WithGormConfig 设置gorm专用配置
func WithGormConfig(config *gorm.Config) Option {
	return func(connector *Connector) {
		connector.gormConfig = config
	}
}

// NewConnector 创建一个连接器
func NewConnector(config mysql.Config, opts ...Option) Connector {

	var connector = Connector{
		config:     config,
		dialector:  nil,
		gormConfig: nil,
		db:         nil,
	}

	for _, opts := range opts {
		opts(&connector)
	}

	return connector
}

// Connector 连接器
type Connector struct {
	config     mysql.Config      // mysql配置
	dialector  *driver.Dialector // mysql驱动配置
	gormConfig *gorm.Config      // gorm特有配置
	db         *gorm.DB          // gorm连接
}

// Connect 连接数据
func (c *Connector) Open() (*gorm.DB, error) {

	var key = c.createOnceMapKey()
	if c.isRepeated(key) {
		c.db = onceMap[key]
		return c.db, nil
	}

	var (
		sqlDB      *sql.DB
		gormConfig *gorm.Config
		err        error
	)

	if c.gormConfig != nil {
		gormConfig = c.gormConfig
	} else {
		gormConfig = &gorm.Config{}
	}

	// 连接数据库
	if c.dialector != nil {
		c.db, err = gorm.Open(c.dialector, gormConfig)
	} else {
		c.db, err = gorm.Open(driver.Open(c.config.ToDSN()), gormConfig)
	}

	if err != nil {
		return nil, err
	}

	if sqlDB, err = c.db.DB(); err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(c.config.MaxIdleConn)
	sqlDB.SetMaxOpenConns(c.config.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(c.config.ConnMaxLifeTime)
	// Go 1.15 or later
	// sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	c.updateInstance(key, c.db)

	return c.db, nil
}

// Close 关闭数据库连接
func (c *Connector) Close() error {
	if c.db != nil {
		var (
			sqlDB *sql.DB
			err   error
		)
		if sqlDB, err = c.db.DB(); err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// updateInstance 更新不重复的实例
func (c *Connector) updateInstance(key string, db *gorm.DB) {
	lock.Lock()
	defer lock.Unlock()
	onceMap[key] = db
}

// createOnceMapKey 生成排重的key
func (c *Connector) createOnceMapKey() string {
	m := md5.New()
	m.Reset()
	_, err := io.WriteString(m, c.config.ToDSN())
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", m.Sum(nil))
}

// isRepeated 判断是不是重复连接
func (c *Connector) isRepeated(key string) bool {
	if _, ok := onceMap[key]; ok {
		return true
	}
	return false
}
