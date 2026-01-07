package main

import (
	"log"
	"net/http"
)

func main() {
	// 服务端
	var address = "127.0.0.1:8080"
	log.Println(address)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/hello/world", HelloWorld)
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(address, nil)

}
