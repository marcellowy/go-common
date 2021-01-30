package internal

import (
	"github.com/spf13/viper"
)

// IsSupportedExt
func IsSupportedExt(ext string) bool {
	for _, v := range viper.SupportedExts {
		if ext == v {
			return true
		}
	}
	return false
}
