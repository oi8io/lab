package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}


func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Println("TCP listen Error", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Println("Accept Error", err)
	}
	rpc.ServeConn(conn)

}
