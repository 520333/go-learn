package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"order/protobufs/compiler"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	grpcAddr = flag.String("grpcAddr", "localhost:5051", "gRPC listen address")
	addr     = flag.String("addr", "127.0.0.1", "http service address")
	port     = flag.Int("port", 8080, "http service port")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	service := http.NewServeMux()
	service.HandleFunc("/orders", func(writer http.ResponseWriter, request *http.Request) {
		//grpc客户端实例化
		client := compiler.NewProductClient(conn)
		// 远程调用 RPC
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		productResp, err := client.ProductInfo(ctx, &compiler.ProductRequest{
			Id: 42,
		})
		if err != nil {
			log.Fatalln(err)
		}
		data := struct {
			ID        int64                       `json:"id"`
			Quantity  int                         `json:"quantity"`
			ProductID []*compiler.ProductResponse `json:"products"`
		}{
			9527,
			1,
			[]*compiler.ProductResponse{productResp},
		}
		dataJson, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		writer.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprintf(writer, string(dataJson)); err != nil {
			log.Fatalln(err)
		}
	})
	address := fmt.Sprintf("%s:%d", *addr, *port)
	fmt.Printf("order http service is listening on %s\n", address)
	log.Fatalln(http.ListenAndServe(address, service))
}
