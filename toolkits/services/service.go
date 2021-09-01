package services

import "fmt"

type services struct {
}

func GetService() *services {
	return &services{}
}
func (s *services) HttpServer() {
	fmt.Println("httpServre")
}

func (s *services) Daemon() {
	fmt.Println("Daemon")
}
