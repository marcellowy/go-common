package config

import (
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestFileConfig_Init(t *testing.T) {
	config := &FileConfig{}
	err := config.Init("config.yaml", func(in fsnotify.Event) {

	})
	if err != nil {
		t.Error(err)
		return
	}

	//fmt.Printf("%s\n", config.MarshalIndent())
}
