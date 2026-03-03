package server

import (
	"context"
	"customer/api/customer"
	v1 "customer/api/helloworld/v1"
	"customer/internal/biz"
	"customer/internal/conf"
	"customer/internal/service"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, customerService *service.CustomerService, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			// 自己设置的中间件
			selector.Server(jwt.Server(func(token *jwt2.Token) (interface{}, error) {
				return []byte(biz.CustomerSecret), nil
			}), func(handler middleware.Handler) middleware.Handler {
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					// 一、获取存储在jwt中的用户（顾客）id
					claims, ok := jwt.FromContext(ctx)
					if !ok {
						// 没有获取到claims
						return nil, errors.Unauthorized("UNAUTHORIZED", "claims not found")
					}
					// 1.2 断言使用
					claimsMap := claims.(jwt2.MapClaims)
					log.Info(claimsMap)
					id := claimsMap["jti"]
					// 二 获取id对应顾客的token
					token, err := customerService.CD.GetToken(id)
					if err != nil {
						return nil, errors.Unauthorized("UNAUTHORIZED", "customer not found")
					}

					// 三 比对数据表中的token与请求的token是否一致
					header, _ := transport.FromServerContext(ctx)
					auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
					jwtToken := auths[1]
					if jwtToken != token {
						return nil, errors.Unauthorized("UNAUTHORIZED", "token was updated")
					}

					// 四 校验通过放行继续执行
					return handler(ctx, req)
				}
			}).Match(func(ctx context.Context, operation string) bool {
				// 根据
				noJWT := map[string]struct{}{
					"/api.customer.Customer/GetVerifyCode": {},
					"/api.customer.Customer/Login":         {},
				}
				if _, exists := noJWT[operation]; exists {
					return false
				}
				return true
			}).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	customer.RegisterCustomerHTTPServer(srv, customerService)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
