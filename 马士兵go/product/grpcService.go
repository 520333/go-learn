package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"protobufs/protobufs/compiler"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5051, "The gRpc server port")
)

func main() {
	flag.Parse()
	listenner, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 实例化grpc服务器
	grpcServer := grpc.NewServer()
	// 注册到grpc服务器中
	compiler.RegisterProductServer(grpcServer, &compiler.UnimplementedProductServer{})
	// 启动监听
	log.Println("grpc server listening on :", listenner.Addr())
	if err := grpcServer.Serve(listenner); err != nil {
		log.Fatalln("failed to serve:", err)
	}

}
