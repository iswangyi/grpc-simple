package main

import (
	"log"
	"net"
	"net/rpc"
)

//构造一个Hello类型方法

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {

	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(conn)
}
