package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
)

type Doctor interface {
	HealthCheck() error
}

func init() {
	log.Println("main progress init function called")
}

func main() {
	pathOfPlugin := "../plug/plug.so"
	open, err := plugin.Open(pathOfPlugin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	md, err := open.Lookup("MyDoctor")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	my, ok := md.(Doctor)
	if !ok {
		fmt.Println(err)
		os.Exit(3)
	}
	if err := my.HealthCheck(); err != nil {
		log.Println("use plugin doctor failed, ", err)
	}
}
