package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"oi.io/apps/zorm"
	"oi.io/apps/zorm/log"
)

func main() {
	log.SetLevel(log.LevelDebug)
	engine, _ := zorm.NewEngine("sqlite3", "zorm.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
