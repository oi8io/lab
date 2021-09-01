package gen

import (
	"fmt"
	"regexp"
	"strings"
)

type Table struct {
	Name    string
	Rows    []Row
	Swagger string
	Csv     string
	Gorm    string
	Json    string
}

type Row struct {
	Name    string
	Type    string
	Column  string
	Comment string
}

var tbl *Table

func NewTable() *Table {
	if tbl != nil {
		return tbl
	}
	return &Table{}
}

func getColumnType(str string) string {
	er := strings.NewReplacer("tinyint", "int", "smallint", "int", "int", "int", "mediumint", "int", "bigint", "int", "float", "float", "double", "float", "decimal", "float", "char", "string", "varchar", "string", "text", "string", "mediumtext", "string", "longtext", "string", "datetime", "int", "timestramp", "int", "enum", "string", "set", "string", "blob", "string")
	return er.Replace(str)
}

func (t *Table) GetRows(str string) {
	a := "\\`(\\w+)\\`\\s+(\\w+)(.*),?"
	r := regexp.MustCompile(a)
	rows := r.FindAllStringSubmatch(str, -1)
	for _, r := range rows {
		rex := "comment '(.*)',?"

		c := regexp.MustCompile(rex)
		x := c.FindAllStringSubmatch(r[3], -1)
		row := Row{Name: camelString(r[1]), Column: r[1], Type: getColumnType(r[2])}
		if len(x) == 1 {
			row.Comment = x[0][1]
		}
		//fmt.Printf("%q\n", row)
		t.Rows = append(t.Rows, row)
	}
}
func (t *Table) GetTableName(str string) {
	a := "(?i)table.*\\`(\\w+)\\`\\s+\\("
	r := regexp.MustCompile(a)
	rows := r.FindAllStringSubmatch(str, -1)
	t.Name = camelString(rows[0][1])
}

func (r *Row) GetFormatString(format string) (str string) {
	switch format {
	case "swaggo":
		str = fmt.Sprintf("%s %s `orm:\"column(%s)\" json:\"%s\" swaggo:\"true,%s\"`", r.Name, r.Type, r.Column, r.Column, r.Comment)
	case "csv":
		str = fmt.Sprintf("| %s | %s | 必填 | %s |", r.Name, r.Type, r.Comment)
	case "gorm":
		str = fmt.Sprintf("%s %s `gorm:\"column:%s\" json:\"%s\" swaggo:\"true,%s\"`", r.Name, r.Type, r.Column, r.Column, r.Comment)
	case "json":
		str = fmt.Sprintf("%s %s `json:\"%s\" mark:\"true,%s\"`", r.Name, r.Type, r.Column, r.Comment)
	}
	return
}
func (t *Table) GetCSV() (str string) {
	str = "| 字段名称 | 字段类型 | 是否必填 | 字段说明 | \n| --- | --- | --- |  --- | \n"
	for _, r := range t.Rows {
		str += r.GetFormatString("csv") + "\n"
	}
	t.Csv = str
	return
}

func (t *Table) GetSwaggo() (str string) {
	str = "type " + t.Name + " struct { \n"
	for _, r := range t.Rows {
		str += "\t" + r.GetFormatString("swaggo") + "\n"
	}
	str += "}\n"
	t.Swagger = str
	return
}
func (t *Table) GetJson() (str string) {
	str = "type " + t.Name + " struct { \n"
	for _, r := range t.Rows {
		str += "\t" + r.GetFormatString("json") + "\n"
	}
	str += "}\n"
	t.Json = str
	return
}
func (t *Table) Format() () {
	t.GetGorm()
	t.GetSwaggo()
	t.GetCSV()
	t.GetJson()
}

func (t *Table) GetGorm() (str string) {
	str = "type " + t.Name + " struct { \n"
	for _, r := range t.Rows {
		str += "\t" + r.GetFormatString("gorm") + "\n"
	}
	str += "}\n"
	t.Gorm = str
	return
}

func toStruct(str string) string {
	a := `"(\w+)":\s+(.*)\,?`
	r := regexp.MustCompile(a)
	body := r.ReplaceAllStringFunc(str, func(s string) string {
		part := r.FindStringSubmatch(s)
		return camelString(part[1]) + ":" + part[2]
	})
	return body
}

func toCamel(str string) string {
	a := "[^|\\n]+\\s+\\w+\\s`"
	r := regexp.MustCompile(a)
	body := r.ReplaceAllStringFunc(str, func(s string) string {
		return camelString(s)
	})
	return body
}

// camel string, xx_yy to XxYy
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
