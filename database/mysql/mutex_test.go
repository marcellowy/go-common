package mysql

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewMutex(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectCommit()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if m := NewMutex(&tx, "table"); m == nil {
		t.Fatal("instance error")
	}
}

func TestMutex_Lock(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery("SELECT id FROM table WHERE key=\\? FOR UPDATE NO WAIT").
		WithArgs("test_key").WillReturnRows(rows)
	mock.ExpectCommit()
	mock.ExpectClose()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if m := NewMutex(&tx, "table"); m == nil {
		t.Fatal("instance error")
	} else {
		if err = m.Lock("test_key", time.Second*3); err != nil {
			t.Fatal(err)
		}
	}

	if err = tx.Commit(); err != nil {
		t.Fatal(err)
	}

}

func TestMutex_Unlock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery("SELECT id FROM table WHERE key=\\? FOR UPDATE NO WAIT").
		WithArgs("test_key").WillReturnRows(rows)
	mock.ExpectCommit()
	mock.ExpectClose()

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if m := NewMutex(&tx, "table"); m == nil {
		t.Fatal("instance error")
	} else {
		if err = m.Lock("test_key", time.Second*3); err != nil {
			t.Fatal(err)
		}
		if err = m.Unlock(); err != nil {
			t.Fatal(err)
		}
	}

	if err = tx.Commit(); err != nil {
		t.Fatal(err)
	}

}
