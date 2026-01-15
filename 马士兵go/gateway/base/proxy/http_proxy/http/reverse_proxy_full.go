package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	// 上游真实服务器
	realServer := "http://127.0.0.1:8001?a=1&b=2#container"
	serverURL, err := url.Parse(realServer)
	if err != nil {
		log.Println(err)
	}
	// 反向代理服务器
	proxy := NewSingleHostReverseProxy(serverURL)
	var addr = "127.0.0.1:8081"
	log.Println("starting proxy http server at:" + addr)
	http.ListenAndServe(addr, proxy)
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = JoinPath(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "")
		}
	}
	// 修改返回内容
	modifyResponse := func(res *http.Response) error {
		fmt.Println("here is modifyResponse Function")
		if res.StatusCode == 200 {
			srcBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			newBody := []byte(string(srcBody) + " 海绵宝宝")
			res.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))
			length := int64(len(newBody))
			res.ContentLength = length
			res.Header.Set("Content-Length", strconv.FormatInt(length, 10))
		}
		return nil
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
}

func JoinPath(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case aslash || bslash:
		return a + b
	}
	return a + "/" + b
}
