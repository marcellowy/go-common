package gorm

import (
	"testing"

	"gitee.com/marcellos/wyi-common/database/mysql"
	"gorm.io/gorm"
)

// TestNewConnector
func TestNewConnector(t *testing.T) {

	// t.SkipNow()

	var (
		connector Connector
		db        *gorm.DB
		err       error
	)

	connector = NewConnector(mysql.NewConfig("9.135.126.130:3306", "mariadb_admin",
		"mEurV@3457cA", ""))

	if db, err = connector.Open(); err != nil {
		t.Error(err)
		return
	}

	if db == nil {
		t.Error("init db error")
		return
	}

	if tx := db.Exec("SHOW DATABASES"); tx.Error != nil {
		t.Error(tx.Error)
		return
	}
}

// TestConnector_Open
func TestConnector_Open(t *testing.T) {

}

// TestConnector_Close
func TestConnector_Close(t *testing.T) {

}
