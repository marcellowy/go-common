// Package config
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package config

import (
	"github.com/spf13/viper"
	"sync"
	"time"
)

// defaultConfigName default config filename
const defaultConfigName = "config"

// defaultConfigPath default config file directory
const defaultConfigPath = "manifests/config"

// defaultConfig default
var defaultConfig *Config

// defaultConfigInit default config initial flag
var defaultConfigInit bool

// Config default instance
type Config struct {
	v *viper.Viper
}

func (c *Config) GetStringSlice(key string) []string {
	return c.v.GetStringSlice(key)
}

func (c *Config) GetStringMap(key string) map[string]interface{} {
	return c.v.GetStringMap(key)
}

func (c *Config) GetStringMapString(key string) map[string]string {
	return c.v.GetStringMapString(key)
}

func (c *Config) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *Config) GetInt64(key string) int64 {
	return c.v.GetInt64(key)
}

func (c *Config) GetFloat64(key string) float64 {
	return c.v.GetFloat64(key)
}

func (c *Config) GetBool(key string) bool {
	return c.v.GetBool(key)
}

func (c *Config) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *Config) GetDuration(key string) time.Duration {
	return c.v.GetDuration(key)
}

func (c *Config) Get(key string) interface{} {
	return c.v.Get(key)
}

// cache global
var cache sync.Map

// default read manifests/config/config.yml
func init() {
	initByName(defaultConfigName)
}

func initByName(name string, path ...string) {
	if _, ok := cache.Load(name); ok {
		// exists
		return
	}
	def := viper.New()
	def.SetConfigName(name) // name of config file (without extension)
	def.AddConfigPath(defaultConfigPath)
	def.AddConfigPath(".")
	for _, v := range path {
		def.AddConfigPath(v)
	}
	def.SetConfigType("yaml")
	if err := def.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			panic(err)
		}
	}

	def.WatchConfig()
	c := &Config{v: def}
	cache.Store(name, c)
	if name == defaultConfigName && defaultConfig == nil {
		defaultConfig = c
	}
}

// Instance from new instance or cache
func Instance(name ...string) *Config {
	if len(name) == 0 || name[0] == "" {
		return defaultConfig
	}

	initByName(name[0])

	if c := readConfig(name[0]); c != nil {
		return c
	}

	return defaultConfig
}

func readConfig(name string) *Config {
	var (
		ok bool
		v  any
		c  *Config
	)
	if v, ok = cache.Load(name); !ok {
		return nil
	} else {
		if c, ok = v.(*Config); ok {
			return c
		}
	}
	return nil
}

func GetStringSlice(key string) []string {
	return defaultConfig.v.GetStringSlice(key)
}

func GetStringMap(key string) map[string]interface{} {
	return defaultConfig.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return defaultConfig.GetStringMapString(key)
}

func GetInt(key string) int {
	return defaultConfig.GetInt(key)
}

func GetInt64(key string) int64 {
	return defaultConfig.GetInt64(key)
}

func GetFloat64(key string) float64 {
	return defaultConfig.GetFloat64(key)
}

func GetBool(key string) bool {
	return defaultConfig.GetBool(key)
}

func GetString(key string) string {
	return defaultConfig.GetString(key)
}

func GetDuration(key string) time.Duration {
	return defaultConfig.GetDuration(key)
}

func Get(key string) interface{} {
	return defaultConfig.Get(key)
}
