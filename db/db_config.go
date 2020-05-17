package db

import "time"

type Config struct {
	Host               string        // 主机或域名
	Port               int32         // 端口
	Username           string        // 用户名
	Password           string        // 密码
	DbName             string        // 数据库名
	Charset            string        // 字符集
	ConnectionSize     int32         // 连接数量
	ConnectionLifetime time.Duration // 连接生命周期
	MaxIdleConnection  int32         // 连接最大空闲时间
}

func NewConfig(host string, port int32, username, password, dbName, charset string) *Config {

	return &Config{
		Host:               host,
		Port:               port,
		Username:           username,
		Password:           password,
		DbName:             dbName,
		Charset:            charset,
		ConnectionSize:     20,
		ConnectionLifetime: time.Second * 20,
		MaxIdleConnection:  0,
	}
}
