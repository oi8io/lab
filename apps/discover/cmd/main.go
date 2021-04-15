package main

import (
	"flag"
	"log"
	"net/http"
	"oi.io/apps/discover/api"
	"oi.io/apps/discover/configs"
	"oi.io/apps/discover/global"
	"oi.io/apps/discover/model"
)

func main() {
	//init config
	c := flag.String("c", "", "config file path")
	flag.Parse()
	config, err := configs.LoadConfig(*c)
	if err != nil {
		log.Println("load config error:", err)
		return
	}
	//global discovery
	global.Discovery = model.NewDiscovery(config)
	//init router and start server
	router := api.InitRouter()
	srv := &http.Server{
		Addr:    config.HttpServer,
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen:%s\n", err)
	}
	go func() {

	}()
}
