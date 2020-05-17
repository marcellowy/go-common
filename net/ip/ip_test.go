package ip

import (
	"log"
	"testing"
)

func TestLocalIPv4s(t *testing.T) {
	ip, err := LocalIPv4s()
	if err != nil {
		t.Error(err)
	}
	if len(ip) == 0 {
		t.Error("ip len 0")
	}
}

func TestGetFreePorts(t *testing.T) {
	port, err := GetIdlePort()
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(port)
}
