package buffer

import (
	"bytes"
	"fmt"

	"gitee.com/marcellos/wyi-common/config/internal"
	"github.com/spf13/viper"
)

// New 从buffer实例化配置
func New(buffer *bytes.Buffer, ext string) (*viper.Viper, error) {

	var (
		err error
	)

	if !internal.IsSupportedExt(ext) {
		return nil, fmt.Errorf("configuration file are not supported")
	}

	v := viper.New()
	{
		v.SetConfigType(ext)
	}

	if err = v.ReadConfig(buffer); err != nil {
		return nil, err
	}

	return v, nil
}
