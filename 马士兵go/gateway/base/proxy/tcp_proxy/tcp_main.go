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
	go func() {
		var addr = ":8003"
		// 1.创建tcpServer实例
		tcpServer := &server.TCPServer{
			Addr:    addr,
			Handler: &handler{},
		}
		// 2.启动监听提供服务
		log.Println("Starting TCP Server at " + addr)
		_ = tcpServer.ListenAndServe()
	}()
	// 启动TCP代理
	go func() {
		var tcpServerAddr = ":8003"
		// 1.创建tcpProxy实例
		tcpProxy := proxy.NewSingleHostReverseProxy(tcpServerAddr)
		// 2.启动监听提供服务
		var addr = ":8083"
		log.Println("Starting TCP Server at " + addr)
		server.ListenAndServe(tcpServerAddr, tcpProxy)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type handler struct{}

func (h *handler) ServeTCP(ctx context.Context, conn net.Conn) {
	conn.Write([]byte("hahaha\n"))
}
