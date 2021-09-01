package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"oi.io/gokit/napo"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8082", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our napo service
	srv := napo.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 映射端点
	endpoints := napo.Endpoints{
		GetEndpoint:      napo.MakeGetEndpoint(srv),
		StatusEndpoint:   napo.MakeStatusEndpoint(srv),
		ValidateEndpoint: napo.MakeValidateEndpoint(srv),
	}

	// HTTP 传输
	go func() {
		log.Println("napo is listening on port:", *httpAddr)
		handler := napo.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
