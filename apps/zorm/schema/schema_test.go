package schema

import (
	"fmt"
	"oi.io/apps/zorm/dialect"
	"testing"
)

func TestParse(t *testing.T) {
	type User struct {
		Name string `geeorm:"PRIMARY KEY"`
		Age  int
	}
	getDialect, _ := dialect.GetDialect("sqlite3")
	parse := Parse(&User{}, getDialect)
	fmt.Print(parse.Tostring())
}


var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse2(t *testing.T) {
	type User struct {
		Name string `geeorm:"PRIMARY KEY"`
		Age  int
	}
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}


func TestSchema_RecordValues(t *testing.T) {
	type User struct {
		Name string `geeorm:"PRIMARY KEY"`
		Age  int
	}
	schema := Parse(&User{}, TestDial)

	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
	values := schema.RecordValues(&User{"tom", 18})
	fmt.Println(values)
}

