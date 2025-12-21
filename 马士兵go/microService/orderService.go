package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
)

func main() {
	// 接受命令行参数作为服务对外的地址和端口
	addr := flag.String("addr", "127.0.0.1", "The Address for listen. Default is  127.0.0.1")
	port := flag.Int("port", 8080, "The Port for listen. Default is  8080")
	flag.Parse()

	// 定义业务逻辑服务，假设为产品服务
	server := http.NewServeMux()
	server.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "Order Service.")
		if err != nil {
			log.Fatalln(err)
		}
	})
	// 注册服务 作为客户端连接consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.50.100:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// 获取单一服务的信息
	//serviceId := "redis1"
	//ag, qm, err := consulClient.Agent().Service(serviceId, nil)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(ag.ID, ag.Address, ag.Port)
	//log.Println(qm)

	// 获取多个服务的信息
	filter := "Service==product"
	services, err := consulClient.Agent().ServicesWithFilter(filter)
	if err != nil {
		log.Fatalln(err)
	}
	for id, svc := range services {
		log.Println(id, svc.Address, svc.Port)
	}
	serviceID := "someService-01"

	// 注销服务
	err = consulClient.Agent().ServiceDeregister(serviceID)
	if err != nil {
		log.Fatalln(err)
	}

	// 启动监听
	address := fmt.Sprintf("%s:%d", *addr, *port)
	log.Printf("Order Service is Listening on %s", address)
	log.Fatalln(http.ListenAndServe(address, server))
}
