package main

import (
	"fmt"
	"log"
	"net/http"
)

func HttpServerSimple() {
	http.HandleFunc("/ping", handlePing)
	infoHandler := InfoHandler{
		info: "<h1>Welcome to Go HTTP Server.</h1>",
	}
	http.Handle("/info", &infoHandler)
	addr := ":8088"
	log.Println("Listening on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

type InfoHandler struct {
	info string
}

func (h *InfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.info)
}

func main() {
	//HttpServerSimple()
	HttpServerCustom()

}
