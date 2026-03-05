package service

import (
	"context"
	"customer/api/verityCode"
	"customer/internal/biz"
	"customer/internal/data"
	"fmt"
	"log"
	"regexp"
	"time"

	pb "customer/api/customer"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/consul/api"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	CD *data.CustomerData
}

func NewCustomerService(cd *data.CustomerData) *CustomerService {
	return &CustomerService{CD: cd}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	// 一 校验手机号
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$`
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(req.Telephone) {
		return &pb.GetVerifyCodeResp{Code: 1, Message: "电话号码格式错误"}, nil
	}

	// 二 通过验证码服务生成验证码(服务间通讯 grpc)
	// 使用服务发现
	// 1.获取consul客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.1.178:8500"
	consulClient, err := api.NewClient(consulConfig)
	// 2.服务发现管理器
	dis := consul.New(consulClient)
	if err != nil {
		log.Fatal(err)
	}
	// 2.1连接目标grpc服务器
	endpoint := "discovery:///verifyCode"
	conn, err := grpc.DialInsecure(context.Background(),
		//grpc.WithEndpoint("localhost:9000"),
		grpc.WithEndpoint(endpoint), // 目标服务的名字
		grpc.WithDiscovery(dis),     // 使用服务发现
	)

	if err != nil {
		return &pb.GetVerifyCodeResp{Code: 1, Message: "验证码服务不可用"}, nil
	}
	defer func() {
		_ = conn.Close()
	}()

	// 2.2 发送获取验证码请求
	client := verityCode.NewVerityCodeClient(conn)
	reply, err := client.GetVerityCode(context.Background(), &verityCode.GetVerityCodeRequest{Length: 6, Type: 1})
	if err != nil {
		return &pb.GetVerifyCodeResp{Code: 1, Message: "验证码获取错误"}, nil
	}

	// 三 redis的临时存储
	// 3.1 连接redis
	const life = 60
	if err = s.CD.SetVerifyCode(req.Telephone, reply.Code, life); err != nil {
		return &pb.GetVerifyCodeResp{Code: 1, Message: "验证码存储错误(redis配置解析错误)"}, nil
	}

	//opt, err := redis.ParseURL("redis://192.168.1.178:6379/1?dial_timeout=1")
	//if err != nil {
	//	return &pb.GetVerifyCodeResp{Code: 1, Message: "验证码存储错误(redis配置解析错误)"}, nil
	//}
	//rdb := redis.NewClient(opt)
	//status := rdb.Set(context.Background(), "CVC:"+req.Telephone, reply.Code, life*time.Second)
	//if _, err = status.Result(); err != nil {
	//	return &pb.GetVerifyCodeResp{Code: 1, Message: "验证码存储错误(redis Set操作错误)"}, nil
	//}

	return &pb.GetVerifyCodeResp{
		Code:           0,
		VerifyCode:     reply.Code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: life,
	}, nil
}

func (s *CustomerService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	// 一 校验电话和验证码
	// redis
	code := s.CD.GetVerifyCode(req.Telephone)
	if code == "" || code != req.VerifyCode {
		fmt.Println(req.VerifyCode, code)
		return &pb.LoginResp{
			Code:    1,
			Message: "验证码不匹配",
		}, nil
	}
	// 二 判断电话号码是否注册 来获取顾客信息
	customer, err := s.CD.GetCustomerByTelephone(req.Telephone)
	if err != nil {
		return &pb.LoginResp{
			Code:    1,
			Message: "顾客信息获取错误",
		}, nil
	}

	// 三 设置token jwt-token
	token, err := s.CD.GenerateTokenAndSave(customer, biz.CustomerDuration*time.Second, biz.CustomerSecret)
	if err != nil {
		return &pb.LoginResp{
			Code:    1,
			Message: "Token生成失败",
		}, nil
	}

	// 四 响应token
	return &pb.LoginResp{
		Code:           0,
		Message:        "login success",
		Token:          token,
		TokenCreatedAt: time.Now().Unix(),
		TokenLife:      int32(biz.CustomerDuration),
	}, nil
}

func (s *CustomerService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
	// 一 获取用户id
	claims, _ := jwt.FromContext(ctx)
	claimsMap := claims.(jwt2.MapClaims)
	if err := s.CD.DelToken(claimsMap["jti"]); err != nil {
		return &pb.LogoutResp{
			Code:    1,
			Message: "Token删除失败",
		}, nil
	}
	return &pb.LogoutResp{
		Code:    0,
		Message: "logout success",
	}, nil
}
