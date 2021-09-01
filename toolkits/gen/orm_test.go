package gen

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewTable(t *testing.T) {
	by, _ := ioutil.ReadFile("/Users/ocean/develop/go/zelo/gen/orm.sql")

	text := string(by)
	text = strings.ToLower(text)
	arr:= strings.Split(text,";")
	for _,t:= range arr {
		if strings.TrimSpace(t) == "" {
			break
		}
		tb := NewTable()
		tb.GetTableName(string(t))
		tb.GetRows(string(t))
		tb.Format()
		//println(tb.Csv)
		println(tb.Gorm)
	}


	//println(tb.Json)
}

func TestOJ(t *testing.T)  {


}