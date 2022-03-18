package gorm

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	driver "gorm.io/driver/mysql"

	"gitee.com/marcellos/common/database/mysql"
)

// TestNewConnector
func TestNewConnector(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
		return
	}

	defer func() {
		// all expectations were already fulfilled, call to database Close was not expected, 是正常
		_ = db.Close()
	}()

	var opts []Option
	var a = WithDialector(&driver.Dialector{Config: &driver.Config{Conn: db, SkipInitializeWithVersion: true}})
	opts = append(opts, a)
	connector := NewConnector(mysql.NewConfig("", "", "", ""), opts...)
	_, err = connector.Open()
	if err != nil {
		t.Error(err)
		return
	}
}

// TestConnector_Open
func TestConnector_Open(t *testing.T) {
	var opts []Option
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		// all expectations were already fulfilled, call to database Close was not expected, 是正常
		_ = db.Close()
	}()

	var a = WithDialector(&driver.Dialector{Config: &driver.Config{Conn: db, SkipInitializeWithVersion: true}})
	opts = append(opts, a)
	connector := NewConnector(mysql.NewConfig("", "", "", ""), opts...)
	_, err = connector.Open()
	if err != nil {
		t.Error(err)
		return
	}
}

// TestConnector_Close
func TestConnector_Close(t *testing.T) {
	var opts []Option
	db, _, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		// all expectations were already fulfilled, call to database Close was not expected, 是正常
		_ = db.Close()
	}()

	var a = WithDialector(&driver.Dialector{Config: &driver.Config{Conn: db}})
	opts = append(opts, a)
	connector := NewConnector(mysql.NewConfig("", "", "", ""), opts...)
	_, err = connector.Open()
	if err != nil {
		t.Error(err)
		return
	}

}
