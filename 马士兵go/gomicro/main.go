package main

import (
	"net/http"

	"github.com/micro/go-micro/v2/web"
)

func main() {
	server := web.NewService(web.Address(":8081"))
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world micro"))
	})
	server.Run()
}
