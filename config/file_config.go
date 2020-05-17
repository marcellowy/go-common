package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type FileConfig struct {
	file  string
	viper *viper.Viper
}

func (c *FileConfig) Init(file string, cb func(in fsnotify.Event)) error {

	if c.viper != nil && c.file == file {
		return nil
	} else {
		c.viper = viper.New()
		c.file = file
	}

	var (
		err       error
		splitName = strings.Split(path.Base(file), ".")
		length    = len(splitName)
		dirName   = path.Dir(file)
	)

	if length < 2 {
		return errors.New(fmt.Sprintf("%s no suffix", file))
	}

	// 取出后缀和名字
	var (
		name   = splitName[length-2]
		suffix = splitName[length-1]
	)

	// 设置配置属性
	c.viper.SetConfigName(name)
	c.viper.SetConfigType(suffix)
	c.viper.AddConfigPath(dirName)

	// 读取配置并返回
	if err = c.viper.ReadInConfig(); err != nil {
		// 读取配置出错
		return err
	}

	// 自动监听文件更新
	// 自动监听文件要放在 ReadInConfig() 后面
	c.viper.WatchConfig()
	c.viper.OnConfigChange(cb)

	return nil
}

// 获取字符串; 没有找到时返回空字符串
func (c *FileConfig) GetString(key string) string {

	return c.viper.GetString(key)
}

// 获取整型; 没有找到时返回 0
func (c *FileConfig) GetInt(key string) int {

	return c.viper.GetInt(key)
}

// 获取整型; 没有找到时返回 0
func (c *FileConfig) GetInt32(key string) int32 {

	return c.viper.GetInt32(key)
}

// 获取整型; 没有找到时返回 0
func (c *FileConfig) GetInt64(key string) int64 {

	return c.viper.GetInt64(key)
}

func (c *FileConfig) GetFloat64(key string) float64 {

	return c.viper.GetFloat64(key)
}

func (c *FileConfig) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

func (c *FileConfig) MarshalIndent() []byte {

	b, err := json.MarshalIndent(c.viper.AllSettings(), "", "    ")
	if err != nil {
		return []byte("")
	}

	return b
}

func (c *FileConfig) Marshal() []byte {

	b, err := json.Marshal(c.viper.AllSettings())
	if err != nil {
		return []byte("")
	}

	return b
}
