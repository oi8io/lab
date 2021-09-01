package services

import (
	"fmt"
	"reflect"
	"testing"
)

func TestServices_Daemon(t *testing.T) {
	cmd := "Daemon"
	in := make([]reflect.Value, 0)
	//in = append(in,reflect.ValueOf(cmd))
	fmt.Println(in)
	service := GetService()
	fun := reflect.ValueOf(service).MethodByName(cmd)
	fun.Call(in)
}

func TestServices_HttpServer(t *testing.T) {
	cmd := "HttpServer"
	in := make([]reflect.Value, 0)
	//in = append(in,reflect.ValueOf(cmd))
	fmt.Println(in)
	service := GetService()
	fun := reflect.ValueOf(service).MethodByName(cmd)
	fun.Call(in)
}
