package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "rpc/simple-rpc/grpc-simple"
)

type SimpleService struct{}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello" + req.Data,
	}
	return &res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listen on 9090...")
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	//Serve()阻塞监听
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
