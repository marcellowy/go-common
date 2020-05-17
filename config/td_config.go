package config

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	vipers "github.com/spf13/viper"
)

// 配置读取
//
type TDConfig struct {
	init bool
	lock sync.RWMutex // data 锁
	data map[string]interface{}
}

func (t *TDConfig) setData(data map[string]interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.data = data
}

func (t *TDConfig) GetString(key string) string {

	var defVal = ""

	if _, ok := t.data[key]; !ok {
		return defVal
	}
	switch t.data[key].(type) {
	case string:
		return fmt.Sprintf("%s", t.data[key])
	}
	return defVal
}

func (t *TDConfig) GetInt(key string) int {

	var defVal int = 0

	if _, ok := t.data[key]; !ok {
		return defVal
	}

	s := t.GetString(key)
	if s == "" {
		return defVal
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defVal
	}
	return int(i)
}

func (t *TDConfig) GetInt32(key string) int32 {

	var defVal int32 = 0

	if _, ok := t.data[key]; !ok {
		return defVal
	}

	s := t.GetString(key)
	if s == "" {
		return defVal
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defVal
	}
	return int32(i)
}

func (t *TDConfig) GetInt64(key string) int64 {

	var defVal int64 = 0

	if _, ok := t.data[key]; !ok {
		return defVal
	}

	s := t.GetString(key)
	if s == "" {
		return defVal
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defVal
	}
	return i
}

func (t *TDConfig) GetFloat64(key string) float64 {

	var defVal float64 = 0

	if _, ok := t.data[key]; !ok {
		return defVal
	}

	s := t.GetString(key)
	if s == "" {
		return defVal
	}

	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defVal
	}
	return i
}

func (t *TDConfig) GetJson(key string) (*vipers.Viper, error) {

	var defVal = new(vipers.Viper)

	if _, ok := t.data[key]; !ok {
		return defVal, nil
	}

	s := t.GetString(key)
	if s == "" {
		return defVal, nil
	}

	return t.parse(s, "json")
}

func (t *TDConfig) GetYaml(key string) (*vipers.Viper, error) {

	var defVal = new(vipers.Viper)

	if _, ok := t.data[key]; !ok {
		return defVal, nil
	}

	s := t.GetString(key)
	if s == "" {
		return defVal, nil
	}

	return t.parse(s, "yaml")
}

func (t *TDConfig) MarshalIndent() []byte {

	if t.data == nil {
		return []byte("")
	}

	r, err := json.MarshalIndent(t.data, "", "    ")
	if err != nil {
		return []byte("")
	}
	return r
}

func (t *TDConfig) Marshal() []byte {

	if t.data == nil {
		return []byte("")
	}

	r, err := json.Marshal(t.data)
	if err != nil {
		return []byte("")
	}
	return r
}

func (t *TDConfig) parse(s, typ string) (*vipers.Viper, error) {
	v := vipers.New()
	v.SetConfigType(typ)
	if err := v.ReadConfig(strings.NewReader(s)); err != nil {
		return nil, err
	}
	return v, nil
}
