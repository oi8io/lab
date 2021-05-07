package main

import (
	"net/http"
	"oi.io/apps/piu/piu"
)

func hello(c *piu.Context) {
	c.Json(http.StatusOK, map[string]interface{}{"status": "OK", "code": 0, "message": "you are crash"})
}

func SayHello(c *piu.Context) {
	c.Json(http.StatusOK, map[string]interface{}{"status": "OK", "code": 0, "message": "you are hello"})
}

func SayBye(c *piu.Context) {
	c.Json(http.StatusOK, map[string]interface{}{"status": "OK", "code": 0, "message": "you are byebye"})
}

func main() {
	addr := ":8086"
	engine := piu.NewEngine()
	engine.Get("/hello", hello)
	engine.Get("/:lang/say_hello", SayHello)
	engine.Get("/:lang/say_bye/*", SayBye)
	err := engine.Run(addr)
	if err != nil {
		//panic(err)
	}
}
