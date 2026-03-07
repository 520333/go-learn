package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NWCors 跨域资源共享中间件
func NWCors() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				ht := tr.(http.Transporter)
				ht.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
				ht.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
				ht.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
				ht.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,Access-Control-Allow-Credentials,User-Agent, Content-Length,, Authorization")
				defer func() {
				}()
			}
			return handler(ctx, req)
		}
	}
}
