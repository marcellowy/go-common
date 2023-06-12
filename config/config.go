// Package config
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package config

import (
	"github.com/spf13/viper"
	"time"
)

// 默认读取 manifests/config/config.yml
func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("manifests/config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			panic(err)
		}
	}
	viper.WatchConfig()
}

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func Get(key string) interface{} {
	return viper.Get(key)
}
