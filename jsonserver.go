package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//构造一个Hello类型方法

type Args struct {
	A, B int
}
type Arith int

func (p *Arith) Hello(request *Args, reply *int) error {
	*reply = request.B * request.A
	return nil
}

func main() {

	//	rpc.RegisterName("HelloService",new(HelloService))
	arith := new(Arith)
	rpc.Register(arith)
	listener, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	go jsonrpc.NewServerCodec(conn)
}
