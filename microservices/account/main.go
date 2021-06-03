package main

import (
	"fmt"
	"oi.io/microservices/account/service"
	"oi.io/microservices/dbclient"
)

var appName = "accountService"

func main() {
	initializeBoltClient()
	service.StartWebServer("6767")
	fmt.Printf("Starting %v\n", appName)
}


// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}