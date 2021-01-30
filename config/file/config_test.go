package file

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestNew(t *testing.T) {

	var (
		c   *viper.Viper
		err error
	)

	c, err = New("H:/tmp/config.yaml", func() {
		fmt.Println(c.GetString("redis.password"))
	})
	if err != nil {
		t.Error(err)
		return
	}
	b := c.GetBool("debug")
	if b == false {
		t.Errorf("bool value error")
		return
	}
}
