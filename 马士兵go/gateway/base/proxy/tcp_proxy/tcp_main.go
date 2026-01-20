package main

import (
	"context"
	"gateway/base/proxy/tcp_proxy/proxy"
	"gateway/base/proxy/tcp_proxy/server"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 启动TCP服务器
	// 后端 TCP Server
	go func() {
		addr := "192.168.1.240:8003"
		tcpServer := &server.TCPServer{
			Addr:    addr,
			Handler: &handler{},
		}
		log.Println("Starting TCP Server at", addr)
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	// 启动TCP代理
	go func() {
		backendAddr := "192.168.1.240:8003"
		proxyHandler := proxy.NewSingleHostReverseProxy(backendAddr)

		listenAddr := "192.168.1.240:8083"
		log.Println("Starting TCP Proxy at", listenAddr)

		if err := server.ListenAndServe(listenAddr, proxyHandler); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type handler struct{}

func (h *handler) ServeTCP(ctx context.Context, conn net.Conn) {
	conn.Write([]byte("hahaha\n"))
	conn.Close()
}
