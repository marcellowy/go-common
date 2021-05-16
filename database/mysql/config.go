package mysql

// 通用数据库连接配置
// Author: chadwang<chadwang@tencent.com>
// Date: 2021/01/01
import (
	"fmt"
	"net/url"
	"time"
)

// 设置默认值
const defaultConfigCharset = "UTF8MB4"
const defaultConfigLoc = "Asia/Shanghai"
const defaultConfigMaxOpenConn = 128
const defaultConfigMaxIdleConn = 8
const defaultConfigConnMaxLifeTime = time.Minute * 5
const defaultConfigConnMaxIdleTime = time.Minute * 5
const defaultConfigProtocol = "tcp"

// Options 数据库选项
type Options func(*Config)

// WithLoc 设置时区
func WithLoc(loc string) Options {
	return func(c *Config) {
		c.loc = url.QueryEscape(loc)
	}
}

// WithCharset 设置字符集
func WithCharset(charset string) Options {
	return func(c *Config) {
		c.charset = charset
	}
}

// WithProtocol 设置协议
func WithProtocol(protocol string) Options {
	return func(c *Config) {
		c.protocol = protocol
	}
}

// WithAllowAllFiles
func WithAllowAllFiles() Options {
	return func(c *Config) {
		c.allowAllFiles = true
	}
}

// WithAllowCleartextPasswords
func WithAllowCleartextPasswords() Options {
	return func(c *Config) {
		c.allowCleartextPasswords = true
	}
}

// WithAllowNativePasswords
func WithAllowNativePasswords() Options {
	return func(config *Config) {
		config.allowNativePasswords = true
	}
}

// WithAllowOldPasswords
func WithAllowOldPasswords() Options {
	return func(config *Config) {
		config.allowOldPasswords = true
	}
}

// WithCheckConnLiveNess
func WithCheckConnLiveNess() Options {
	return func(config *Config) {
		config.checkConnLiveNess = true
	}
}

// WithCollation
func WithCollation(collation string) Options {
	return func(config *Config) {
		config.collation = collation
	}
}

// WithClientFoundRows
func WithClientFoundRows() Options {
	return func(config *Config) {
		config.clientFoundRows = true
	}
}

// WithColumnsWithAlias
func WithColumnsWithAlias() Options {
	return func(config *Config) {
		config.columnsWithAlias = true
	}
}

// WithInterpolateParams
func WithInterpolateParams() Options {
	return func(config *Config) {
		config.interpolateParams = true
	}
}

// WithMaxAllowedPacket
func WithMaxAllowedPacket(maxAllowedPacket uint) Options {
	return func(config *Config) {
		config.maxAllowedPacket = maxAllowedPacket
	}
}

// WithMultiStatements
func WithMultiStatements() Options {
	return func(config *Config) {
		config.multiStatements = true
	}
}

// WithParseTime
func WithParseTime() Options {
	return func(config *Config) {
		config.parseTime = true
	}
}

// WithReadTimeout
func WithReadTimeout(readTimeout time.Duration) Options {
	return func(config *Config) {
		config.readTimeout = readTimeout
	}
}

// WithRejectReadOnly
func WithRejectReadOnly() Options {
	return func(config *Config) {
		config.rejectReadOnly = true
	}
}

// WithServerPubKey
func WithServerPubKey(serverPubKey string) Options {
	return func(config *Config) {
		config.serverPubKey = serverPubKey
	}
}

// WithTimeout
func WithTimeout(timeout time.Duration) Options {
	return func(config *Config) {
		config.timeout = timeout
	}
}

// WithWriteTimeout
func WithWriteTimeout(writeTimeout time.Duration) Options {
	return func(config *Config) {
		config.writeTimeout = writeTimeout
	}
}

// WithMaxOpenConn
func WithMaxOpenConn(maxOpenConn int) Options {
	return func(config *Config) {
		config.MaxOpenConn = maxOpenConn
	}
}

// WithMaxIdleConn
func WithMaxIdleConn(maxIdleConn int) Options {
	return func(config *Config) {
		config.MaxIdleConn = maxIdleConn
	}
}

// WithConnMaxLifeTime
func WithConnMaxLifeTime(connMaxLifeTime time.Duration) Options {
	return func(config *Config) {
		config.ConnMaxLifeTime = connMaxLifeTime
	}
}

// WithConnMaxIdleTime
func WithConnMaxIdleTime(connMaxIdleTime time.Duration) Options {
	return func(config *Config) {
		config.ConnMaxIdleTime = connMaxIdleTime
	}
}

// WithTls
func WithTls(tls string) Options {
	return func(config *Config) {
		config.tls = tls
	}
}

// Config 数据库配置
type Config struct {
	address  string
	username string
	password string
	name     string
	charset  string // 默认 UTF8MB4
	loc      string // 默认 Asia/Shanghai
	protocol string // 协议 默认tcp

	allowAllFiles           bool
	allowCleartextPasswords bool
	allowNativePasswords    bool
	allowOldPasswords       bool
	checkConnLiveNess       bool
	collation               string // 默认 UTF8MB4_GENERAL_CI
	clientFoundRows         bool
	columnsWithAlias        bool
	interpolateParams       bool
	maxAllowedPacket        uint // 默认 4MB
	multiStatements         bool
	parseTime               bool
	readTimeout             time.Duration // 默认 0
	rejectReadOnly          bool
	serverPubKey            string        // 默认 none
	timeout                 time.Duration // 默认 OS default
	writeTimeout            time.Duration // 默认 0
	tls                     string        // 可选值: true, false, skip-verify, preferred, <name>

	// 以下选项取值参考: https://golang.org/pkg/database/sql/
	MaxOpenConn     int           // 默认: 128
	MaxIdleConn     int           // 默认: 8
	ConnMaxLifeTime time.Duration // 默认: 5min
	ConnMaxIdleTime time.Duration // 默认: 5min, need go 1.15 or later
}

// ToDSN 转为dsn string
func (cc *Config) ToDSN() string {

	var dsn = fmt.Sprintf("%s:%s@%s(%s)/%s?loc=%s&charset=%s",
		cc.username,
		cc.password,
		cc.protocol,
		cc.address,
		cc.name,
		cc.loc,
		cc.charset,
	)

	// 补充选项
	var options string
	if cc.allowAllFiles {
		options += "&allowAllFiles=true"
	}

	if cc.allowCleartextPasswords {
		options += "&allowCleartextPasswords=true"
	}

	if cc.allowNativePasswords {
		options += "&allowNativePasswords=true"
	}

	if cc.allowOldPasswords {
		options += "&allowOldPasswords=true"
	}

	if cc.checkConnLiveNess {
		options += "&checkConnLiveness=true"
	}

	if cc.collation != "" {
		options += "&collation=" + cc.collation
	}

	if cc.clientFoundRows {
		options += "&clientFoundRows=true"
	}

	if cc.columnsWithAlias {
		options += "&columnsWithAlias=true"
	}

	if cc.interpolateParams {
		options += "&interpolateParams=true"
	}

	if cc.maxAllowedPacket > 0 {
		options += fmt.Sprintf("&maxAllowedPacket=%d", cc.maxAllowedPacket)
	}

	if cc.multiStatements {
		options += "&multiStatements=true"
	}

	if cc.parseTime {
		options += "&parseTime=true"
	}

	if cc.readTimeout != 0 {
		options += fmt.Sprintf("&readTimeout=%d", cc.readTimeout.Milliseconds())
	}

	if cc.rejectReadOnly {
		options += "&rejectReadOnly=true"
	}

	if cc.serverPubKey != "" {
		options += "&serverPubKey=" + cc.serverPubKey
	}

	if cc.timeout != 0 {
		options += fmt.Sprintf("&timeout=%d", cc.timeout.Milliseconds())
	}

	if cc.tls != "" {
		options += "&tls=" + cc.tls
	}

	if cc.writeTimeout != 0 {
		options += fmt.Sprintf("&writeTimeout=%d", cc.writeTimeout.Milliseconds())
	}

	return dsn + options
}

// NewConfig
func NewConfig(address, username, password, name string, opts ...Options) Config {

	var config = Config{
		address:         address,
		username:        username,
		password:        password,
		name:            name,
		charset:         defaultConfigCharset,
		loc:             url.QueryEscape(defaultConfigLoc),
		protocol:        defaultConfigProtocol,
		MaxOpenConn:     defaultConfigMaxOpenConn,
		MaxIdleConn:     defaultConfigMaxIdleConn,
		ConnMaxLifeTime: defaultConfigConnMaxLifeTime,
		ConnMaxIdleTime: defaultConfigConnMaxIdleTime,
	}

	// 应用选项
	for _, opt := range opts {
		opt(&config)
	}

	return config
}
