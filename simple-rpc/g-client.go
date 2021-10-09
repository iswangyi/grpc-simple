package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "rpc/simple-rpc/grpc-simple"
)

func main() {
	//链接服务器
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net connect err: %v", err)
	}
	defer conn.Close()
	//建立grpc链接
	grpcClient := pb.NewSimpleClient(conn)
	//创建发送结构体
	req := pb.SimpleRequest{Data: "grpccc"}
	//调用服务Route
	res, err := grpcClient.Route(context.Background(), &req)
	log.Println(res)
	if err != nil {
		log.Fatalf("call route err :%v", err)
	}
}
