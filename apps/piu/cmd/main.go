package main

import (
	"net/http"
	"oi.io/apps/piu/piu"
)

func hello(c *piu.Context) {
	c.Json(http.StatusOK, piu.H{"status": "OK", "code": 0, "message": "you are crash"})
}

func SayHello(c *piu.Context) {
	c.Json(http.StatusOK, piu.H{"status": "OK", "code": 0, "message": "you are hello"})
}

func SayBye(c *piu.Context) {
	c.Json(http.StatusOK, piu.H{"status": "OK", "code": 0, "message": "you are byebye"})
}

func main() {
	addr := ":8086"
	engine := piu.NewEngine()
	v1 := engine.Group("/v1")
	{
		v1.Get("/hello", hello)
		v1.Get("/:lang/say_hello", SayHello)
		v1.Get("/:lang/say_bye/*", SayBye)
	}
	v1Say := v1.Group("/say")
	{
		v1Say.Get("/hello", hello)
		v1Say.Get("/:lang/say_hello", SayHello)
		v1Say.Get("/:lang/say_bye/*", SayBye)
	}

	v2 := engine.Group("/v2")
	{
		v2.Get("/hello", hello)
		v2.Get("/:lang/say_hello", SayHello)
		v2.Get("/:lang/say_bye/*", SayBye)
	}
	err := engine.Run(addr)
	if err != nil {
		//panic(err)
	}
}
