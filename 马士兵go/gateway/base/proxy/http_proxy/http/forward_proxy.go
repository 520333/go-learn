package main

import (
	"fmt"
	"io"
	"net/http"
)

// 正向代理
func main() {
	fmt.Println("正向代理服务器启动:8081")
	http.Handle("/", &Pxy{})
	http.ListenAndServe(":8081", nil)
}

type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodConnect {
		fmt.Println("[HTTPS]", req.Host)
		return
	}
	host := req.URL.Host
	if host == "" {
		host = req.Host
	}

	fmt.Printf(
		"客户端：%s 访问：[HTTP] %s %s://%s%s\n",
		req.RemoteAddr,
		req.Method,
		req.URL.Scheme,
		host,
		req.URL.Path,
	)

	// 1.代理服务器接受客户端请求，复制 封装成新请求
	outReq := &http.Request{}
	*outReq = *req

	// 2.发送新请求到下游真实服务器，接受响应
	transport := http.DefaultTransport
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	// 3.处理响应并返回上游客户端
	for k, v := range res.Header {
		for _, q := range v {
			rw.Header().Add(k, q)
		}
	}
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
	res.Body.Close()
}
