package router

import (
	"context"
	"net/http"
	"strings"
)

const abortIndex int8 = 63

type HandlerFunc func(*SliceRouteContext)

type SliceRouter struct {
	groups []*sliceRoute
}

type sliceRoute struct {
	*SliceRouter
	path     string
	handlers []HandlerFunc
}

//type SliceRouterContext struct{}

type SliceRouteContext struct {
	*sliceRoute

	index int8

	Ctx context.Context
	Req *http.Request
	Rw  http.ResponseWriter
}

func NewSliceRouter() *SliceRouter {
	return &SliceRouter{}
}

func (g *SliceRouter) Group(path string) *sliceRoute {
	return &sliceRoute{
		SliceRouter: g,
		path:        path,
	}
}

func (route *sliceRoute) Use(middlewares ...HandlerFunc) *sliceRoute {
	route.handlers = append(route.handlers, middlewares...)
	flag := false
	for _, r := range route.SliceRouter.groups {
		if route == r {
			flag = true
			break
		}
	}
	if !flag {
		route.SliceRouter.groups = append(route.SliceRouter.groups, route)
	}
	return route
}

type handler func(*SliceRouteContext) http.Handler

type SliceRouterHandler struct {
	h      handler
	router *SliceRouter
}

func NewSliceRouterHandler(h handler, router *SliceRouter) *SliceRouterHandler {
	return &SliceRouterHandler{h: h, router: router}
}

func (rh *SliceRouterHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	c := newSliceRouterContext(rw, req, rh.router)
	if rh.h != nil {
		c.handlers = append(c.handlers, func(*SliceRouteContext) {
			rh.h(c).ServeHTTP(c.Rw, c.Req)
		})
	}
	c.Reset()
	c.Next()
}

func newSliceRouterContext(rw http.ResponseWriter, req *http.Request, r *SliceRouter) *SliceRouteContext {
	sr := &sliceRoute{}
	matchUrlLen := 0
	for _, route := range r.groups {
		if strings.HasPrefix(req.RequestURI, route.path) {
			pathLen := len(route.path)
			if pathLen > matchUrlLen {
				matchUrlLen = pathLen
				*sr = *route
			}
		}
	}
	c := &SliceRouteContext{
		Rw:         rw,
		Req:        req,
		Ctx:        req.Context(),
		sliceRoute: sr}
	c.Reset()
	return c

}

func (c *SliceRouteContext) Get(key interface{}) interface{} {
	return c.Ctx.Value(key)
}

func (c *SliceRouteContext) Set(key, val interface{}) {
	c.Ctx = context.WithValue(c.Ctx, key, val)
}

func (c *SliceRouteContext) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *SliceRouteContext) Abort() {
	c.index = abortIndex
}

func (c *SliceRouteContext) IsAborted() bool {
	return c.index >= abortIndex
}

func (c *SliceRouteContext) Reset() {
	c.index = -1
}
