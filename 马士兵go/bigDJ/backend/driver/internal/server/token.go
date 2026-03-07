package server

import (
	"context"
	"driver/internal/service"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv4 "github.com/golang-jwt/jwt/v5"
)

// DriverToken 中间件
func DriverToken(service *service.DriverService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			//1 校验jwt 获取其中的司机标识tel
			claims, ok := jwt.FromContext(ctx)
			if !ok {
				return nil, errors.Unauthorized("Unauthorized", "Claims not found")
			}
			claimsMap := claims.(jwtv4.MapClaims)
			tel := claimsMap["jti"]

			//2 利用tel 获取存储在司机表中的token
			token, err := service.Bz.Di.GetToken(ctx, tel.(string))
			if err != nil {
				return nil, errors.Unauthorized("Unauthorized", "driver token not found")
			}

			//3 比对两个token和请求头中比对
			header, _ := transport.FromServerContext(ctx)
			auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
			reqToken := auths[1]
			if token != reqToken {
				return nil, errors.Unauthorized("Unauthorized", "token was updated")
			}

			//4 记录司机信息
			driver, err := service.Bz.Di.FetchInfoByTel(ctx, tel.(string))
			if err != nil {
				return nil, errors.Unauthorized("Unauthorized", "driver was found")
			}
			ctxWithDriver := context.WithValue(ctx, "driver", driver)
			//ctxWithDriver.Value("driver")

			//5 jwt校验通过
			return handler(ctxWithDriver, req)
		}
	}
}
