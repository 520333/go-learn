package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func Client() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second, // 探活时间
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,              // 最大空闲
		IdleConnTimeout:       90 * time.Second, // 空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}
	// 客户端
	var client = http.Client{Timeout: 30 * time.Second, Transport: transport}
	resp, err := client.Get("http://localhost:8080/hello/world")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	bds, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bds))
}
