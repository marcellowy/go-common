package config

import "github.com/spf13/viper"

// Config 继承viper
type Config struct {
	*viper.Viper
}

// NewConfig 实例化配置
func New() *Config {
	return &Config{
		viper.New(),
	}
}
