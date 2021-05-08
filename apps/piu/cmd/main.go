package main

import (
	"oi.io/apps/piu/api"
	"oi.io/apps/piu/piu"
)

func main() {
	addr := ":8086"
	engine := piu.NewEngine()
	engine.Use(piu.Logger())
	engine.AddFuncMap("FormatAsDate", api.FormatAsDate)
	engine.LoadHTMLGlob("/Users/anker/develop/person/lab/apps/piu/cmd/templates/*")
	engine.Static("/assets", "./static")
	engine.Get("/students", api.Students)
	engine.Get("/date", api.Date)
	engine.Get("/hello", api.SayHello)

	v1 := engine.Group("/v1")
	//v1.Use(piu.Logger())
	v1.Use(piu.Logger())
	{
		v1.Get("/hello", api.SayHello)
		v1.Get("/:lang/say_hello", api.SayHello)
		v1.Get("/:lang/say_bye/*", api.SayBye)
	}
	v1Say := v1.Group("/say")
	//v1Say.Use(piu.Logger())
	{
		v1Say.Get("/hello", api.SayHello)
		v1Say.Get("/:lang/say_hello", api.SayHello)
		v1Say.Get("/:lang/say_bye/*", api.SayBye)
	}

	v2 := engine.Group("/v2")
	{
		v2.Get("/hello", api.SayHello)
		v2.Get("/:lang/say_hello", api.SayHello)
		v2.Get("/:lang/say_bye/*", api.SayBye)
	}
	err := engine.Run(addr)
	if err != nil {
		//panic(err)
	}

}
