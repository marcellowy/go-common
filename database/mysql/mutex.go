package mysql

import (
	"context"
	"database/sql"
	"time"
)

// 数据库结构
// CREATE TABLE `t_lock` (
// `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
// `key` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '',
// `expire_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// PRIMARY KEY `id` (`id`),
// UNIQUE KEY `key` (`key`) USING HASH
// ) ENGINE = INNODB DEFAULT CHARSET = UTF8MB4;

//  defaultTimeout 查询数据库默认超时时间
const defaultTimeout = time.Second * 3

// Mutex 数据库锁
type Mutex struct {
	tx    *sql.Tx // 数据库连接
	table string  // key
}

// NewMutex 实例化一个数据库锁
func NewMutex(tx *sql.Tx, table string) *Mutex {
	return &Mutex{
		table: table,
		tx:    tx,
	}
}

// Lock 上锁,异常时由 innodb_lock_wait_timeout 决定
// 若未超时,由事务回滚或提交决定
// expires 作为数据库语句超时时间存在
func (m *Mutex) Lock(key string, expires time.Duration) error {
	var s = "SELECT id FROM " + m.table + " WHERE key=? FOR UPDATE NO WAIT"
	ctx, cancel := context.WithTimeout(context.Background(), expires)
	defer cancel()
	_, err := m.tx.QueryContext(ctx, s, key)
	return err
}

// Unlock 释放锁
func (m *Mutex) Unlock() error {
	// 提交或者回滚可以释放这个锁
	return m.tx.Rollback()
}
