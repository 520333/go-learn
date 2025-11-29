package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HttpServerCustom() {
	addr := ":8080"
	handler := &CustomInfoHandler{"Custom HTTP Server."}
	log.Println("Listening on", addr)

	server := http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		MaxHeaderBytes:    1 << 10,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(addr)
	}
}

type CustomInfoHandler struct {
	info string
}

func (h *CustomInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, h.info)
}
func mian() {
	HttpServerCustom()
}
