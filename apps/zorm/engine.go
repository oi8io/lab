package zorm

import (
	"database/sql"
	"fmt"
	"oi.io/apps/zorm/dialect"
	"oi.io/apps/zorm/log"
	"oi.io/apps/zorm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Error(err)
		return
	}
	getDialect, ok := dialect.GetDialect(driver)
	if !ok {
		return nil, fmt.Errorf("dialect %s Not Found", driver)
	}
	e = &Engine{db: db, dialect: getDialect}
	log.Infof("Database Connect to [%s][%s] success", driver, source)
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Errorf("Failed to close database %s", err)
	} else {
		log.Info("Close database success")
	}
}

func (e *Engine) NewSession() *session.Session {
	return session.NewSession(e.db, e.dialect)
}
