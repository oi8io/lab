package session

import "oi.io/apps/zorm/log"

type Translation interface {
	Begin() (err error)
	Commit() (err error)
	Rollback() (err error)
}

func (s *Session) Begin() (err error) {
	log.Info("transaction begin")
	s.tx, err = s.db.Begin()
	if err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction Commit")
	err = s.tx.Commit()
	if err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction Rollback")
	err = s.tx.Rollback()
	if err != nil {
		log.Error(err)
	}
	return
}
