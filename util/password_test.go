package util

import (
	"fmt"
	"testing"
)

func TestPasswordHash(t *testing.T) {
	fmt.Println(PasswordHash("marcello"))
}
