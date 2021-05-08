package api

import (
	"fmt"
	"net/http"
	"oi.io/apps/piu/piu"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func SayHello(c *piu.Context) {
	c.Json(http.StatusOK, piu.H{"status": "OK", "code": 0, "message": "you are hello"})
}

func Panic(c *piu.Context) {
	panic("try to Panic")
}

func SayBye(c *piu.Context) {
	c.Json(http.StatusOK, piu.H{"status": "OK", "code": 0, "message": "you are byebye"})
}

func Students(c *piu.Context) () {
	fmt.Println("execting")
	stu1 := &student{Name: "BiuBiu", Age: 20}
	stu2 := &student{Name: "BiuBiu1", Age: 22}
	content := piu.H{
		"title":  "piu",
		"stuArr": [2]*student{stu1, stu2},
	}
	c.HTML(http.StatusOK, "students.html", content)
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func Date(c *piu.Context) () {
	c.HTML(http.StatusOK, "date.html", piu.H{
		"title": "biubiu",
		"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
	})
}
