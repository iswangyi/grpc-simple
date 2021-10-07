package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	client := jsonrpc.NewClient(net.Dial("tcp", "127.0.0.1:9000"))
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	var reply int
	type Arith struct {
		A, B int
	}
	arith := &Arith{
		A: 7,
		B: 8,
	}
	err = client.Call("Arith.Hello", arith, &reply)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(reply)
}
