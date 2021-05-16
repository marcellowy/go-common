package internal

import (
	"github.com/spf13/viper"
)

// IsSupportedExt 判断配置后缀是否支持
func IsSupportedExt(ext string) bool {
	for _, v := range viper.SupportedExts {
		if ext == v {
			return true
		}
	}
	return false
}
