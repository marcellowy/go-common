package file

import (
	"fmt"
	"path/filepath"
	"strings"

	"gitee.com/marcellos/wyi-common/crypto"

	"github.com/fsnotify/fsnotify"

	"gitee.com/marcellos/wyi-common/config/internal"

	common "gitee.com/marcellos/wyi-common"

	"github.com/spf13/viper"
)

// New 从文件初始化配置
func New(file string, cb func()) (*viper.Viper, error) {

	var (
		dir           = filepath.Dir(file)
		filename      = filepath.Base(file)
		err           error
		ok            bool
		hash, newHash string
	)

	if ok, err = common.PathExists(file); err != nil || !ok {
		return nil, fmt.Errorf("file not exists or no permission")
	}

	f := strings.Split(filename, ".")
	if len(f) != 2 {
		return nil, fmt.Errorf("file must contain path, name and suffix; like: /path/to/conf.yaml")
	}

	if !internal.IsSupportedExt(f[1]) {
		return nil, fmt.Errorf("configuration file are not supported")
	}

	v := viper.New()
	{
		v.SetConfigName(f[0])
		v.AddConfigPath(dir)
		v.SetConfigType(f[1])
	}

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}

	hash, err = crypto.FileMd5(file)
	if err != nil {
		return nil, err
	}

	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		if newHash, err = crypto.FileMd5(file); err != nil {
			return
		}

		if hash != newHash {
			hash = newHash
			cb()
		}
	})

	return v, nil
}
