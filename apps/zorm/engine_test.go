package zorm

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"oi.io/apps/zorm/session"
	"testing"
)

type User struct {
	Name string `zorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	engine, _ := NewEngine("sqlite3", "zorm.db")
	s := engine.NewSession().Model(&User{})
	//log.Warnf("s.RefTable().Tostring():%s",s.RefTable().Tostring())
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}
}

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "zero.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return nil, errors.New("Error")
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}
