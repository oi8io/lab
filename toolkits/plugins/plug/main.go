//+build  go build --buildmode=plugin
package main

import (
	"log"
	"time"
)

type GoodDoctor string

func init() {
	log.Println("plugin init function called")
}
func (g *GoodDoctor) HealthCheck() error {
	log.Println("now is", *g)
	log.Println("status is OK")
	return nil
}

var MyDoctor = GoodDoctor(time.Now().Format(time.RFC3339))
