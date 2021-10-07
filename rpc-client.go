package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	var reply string
	err = client.Call("HelloService.Hello", "hehe", &reply)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(reply)
}
