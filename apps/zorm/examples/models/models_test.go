package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"oi.io/apps/zorm/dialect"
	"oi.io/apps/zorm/session"
	"testing"
)

var (
	driver, source = "sqlite3", "/tmp/zorm.db"
	user1          = &User{"Tom", 18}
	user2          = &User{"Sam", 25}
	user3          = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *session.Session {
	t.Helper()
	db, _ := sql.Open(driver, source)
	_dialect, _ := dialect.GetDialect(driver)
	s := session.NewSession(db, _dialect).Model(&Account{})
	//err1 := s.DropTable()
	//err2 := s.CreateTable()
	//_, err3 := s.Insert(user1, user2)
	//if err1 != nil || err2 != nil || err3 != nil {
	//	t.Fatal("failed init test records")
	//}
	return s
}

func TestSession_CallMethod(t *testing.T) {
	s := testRecordInit(t)
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "123456"}, &Account{2, "qwerty"})

	u := &Account{}

	err := s.First(u)
	if err != nil || u.ID != 1001 || u.Password != "******" {
		t.Fatal("Failed to call hooks after query, got", u)
	}
}
