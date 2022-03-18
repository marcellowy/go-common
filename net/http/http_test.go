package http

import (
	"testing"
)

func TestClient_Get(t *testing.T) {
	client := Client{}
	body, err := client.Get("https://www.fourfire.wang")
	if err != nil {
		t.Error(err)
	}

	if len(body) > 0 {
	}

	//fmt.Println("body length:", len(body))
}
