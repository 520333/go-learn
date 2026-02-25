package router

import (
	"context"
	"fmt"
	"gateway/base/proxy"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestSliceRouter(t *testing.T) {
	var addr = "127.0.0.1:8006"
	log.Println("Starting httpserver at " + addr)

	sliceRouter := NewSliceRouter()
	routerRoot := sliceRouter.Group("/")
	routerRoot.Use(handle, func(c *SliceRouteContext) {
		fmt.Println("reverse proxy")
		reverseProxy(c.Ctx).ServeHTTP(c.Rw, c.Req)
	})

	routeBase := sliceRouter.Group("/base")
	routeBase.Use(handle, func(c *SliceRouteContext) {
		c.Rw.Write([]byte("test function"))
	})

	var routerHandler http.Handler = NewSliceRouterHandler(nil, sliceRouter)
	http.ListenAndServe(addr, routerHandler)
}

func handle(c *SliceRouteContext) {
	log.Println("trace_in")
	c.Next()
	log.Println("trace_out")
}

func reverseProxy(c context.Context) http.Handler {
	rs1 := "http://127.0.0.1:8001/"
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	rs2 := "http://127.0.0.1:8002/haha"
	url2, err2 := url.Parse(rs2)
	if err2 != nil {
		log.Println(err2)
	}
	urls := []*url.URL{url1, url2}
	return proxy.NewMultipleHostReverseProxy(c, urls)
}
