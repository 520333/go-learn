package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

func main() {
	// 接受命令行参数作为服务对外的地址和端口
	addr := flag.String("addr", "127.0.0.1", "The Address for listen. Default is  127.0.0.1")
	port := flag.Int("port", 8080, "The Port for listen. Default is  8080")
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *addr, *port)

	// 定义业务逻辑服务，假设为产品服务
	server := http.NewServeMux()
	server.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "Product Service.")
		if err != nil {
			log.Fatalln(err)
		}
	})

	// 增加 /health 路由处理器
	server.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Health Check.")
		_, err := fmt.Fprintf(writer, "Healthy.")
		if err != nil {
			log.Fatalln(err)
		}
	})
	// 定义注册中心的服务
	serviceReg := new(api.AgentServiceRegistration)
	serviceReg.Name = "product"
	serviceReg.ID = "Product-" + uuid.NewString()
	serviceReg.Address = *addr
	serviceReg.Port = *port
	serviceReg.Tags = []string{"product", "golang"}
	serviceReg.Check = &api.AgentServiceCheck{
		CheckID:                        serviceReg.ID + "-check",
		Name:                           serviceReg.ID + "-check",
		Interval:                       "5s",
		Timeout:                        "1s",
		HTTP:                           fmt.Sprintf("http://%s/health", address),
		Method:                         "GET",
		Notes:                          "Product Server Health",
		SuccessBeforePassing:           0,
		FailuresBeforeCritical:         0,
		FailuresBeforeWarning:          0,
		DeregisterCriticalServiceAfter: "10s", // 10秒不健康会被取消注册
	}
	// 注册服务 作为客户端连接consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.50.100:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalln(err)
	}

	if err := consulClient.Agent().ServiceRegister(serviceReg); err != nil {
		log.Fatalln(err)
	}

	// 启动监听
	log.Printf("Product Service is Listening on %s", address)
	log.Fatalln(http.ListenAndServe(address, server))
}
