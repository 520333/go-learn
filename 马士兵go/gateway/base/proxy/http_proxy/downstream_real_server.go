package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server1 := &RealServer{Addr: ":8001"}
	server1.Run()
	server2 := &RealServer{Addr: ":8002"}
	server2.Run()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

type RealServer struct {
	Addr string
}

func (r *RealServer) Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/realserver", r.HelloHandler)
	server := &http.Server{Addr: r.Addr, Handler: mux, WriteTimeout: time.Second * 3}
	go func() {
		server.ListenAndServe()
	}()
}

func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
	newPath := fmt.Sprintf("Here is real server http://%s%s", req.RemoteAddr, req.URL.Path)
	w.Write([]byte(newPath))
}
