package server

import (
	"context"
	"customer/internal/service"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

func customerJWT(customerService *service.CustomerService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
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
	}
}
