package models

import (
	"oi.io/apps/zorm/log"
	"oi.io/apps/zorm/session"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

type Account struct {
	ID       int `geeorm:"PRIMARY KEY"`
	Password string
}

func (account *Account) BeforeInsert(s *session.Session) error {
	log.Info("before inert", account)
	account.ID += 1000
	return nil
}

func (account *Account) AfterQuery(s *session.Session) error {
	log.Info("after query", account)
	account.Password = "******"
	return nil
}
