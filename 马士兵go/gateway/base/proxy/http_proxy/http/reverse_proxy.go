package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var port = "8080"
	http.HandleFunc("/", handler)
	fmt.Println("反向代理服务器启动：" + port)
	http.ListenAndServe(":"+port, nil)
}

var (
	proxyAddr = "http://192.168.1.240:8001"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 1.解析下游服务器地址，更改请求地址
	realServer, _ := url.Parse(proxyAddr)
	r.URL.Scheme = realServer.Scheme
	r.URL.Host = realServer.Host
	// 2.请求下游(真实服务器)
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(r)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}
	// 3.把下游请求内容做一些处理，然后返回给上游(客户端)
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	bufio.NewReader(resp.Body).WriteTo(w)
}
