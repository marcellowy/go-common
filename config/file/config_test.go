package file

import (
	"fmt"
	"testing"

	"gitee.com/marcellos/wyi-common/config"
)

func TestNew(t *testing.T) {

	var (
		c   *config.Config
		err error
	)

	c, err = New("config.yaml", func() {
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
