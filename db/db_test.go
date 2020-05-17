package db

import (
	"database/sql"
	"testing"
)

func TestConnect(t *testing.T) {

	//t.SkipNow()

	var (
		conn   *sql.DB = nil
		config *Config = NewConfig("127.0.0.1", 3306, "root",
			"marcello123", "", "utf8mb4")
	)

	if err := Connect(&conn, config); err != nil {
		t.Error(err)
	}
}
