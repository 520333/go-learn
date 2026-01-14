package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// 上游真实服务器
	realServer := "http://127.0.0.1:8001?a=1&b=2#container"
	serverURL, err := url.Parse(realServer)
	if err != nil {
		log.Println(err)
	}
	// 反向代理服务器
	proxy := httputil.NewSingleHostReverseProxy(serverURL)
	var addr = "127.0.0.1:8081"
	log.Println("starting proxy http server at:" + addr)
	http.ListenAndServe(addr, proxy)
}
