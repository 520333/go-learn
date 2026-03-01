package service

import (
	"context"
	"customer/api/verityCode"
	"customer/internal/biz"
	"customer/internal/data"
	"fmt"
	"regexp"
	"time"

	pb "customer/api/customer"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	cd *data.CustomerData
}

func NewCustomerService(cd *data.CustomerData) *CustomerService {
	return &CustomerService{cd: cd}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	// 一 校验手机号
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$`
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(req.Telephone) {
		return &pb.GetVerifyCodeResp{Code: 1, Message: "电话号码格式错误"}, nil
	}
	// 二 通过验证码服务生成验证码(服务间通讯 grpc)
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("localhost:9000"))
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
	if err = s.cd.SetVerifyCode(req.Telephone, reply.Code, life); err != nil {
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
	code := s.cd.GetVerifyCode(req.Telephone)
	if code == "" || code != req.VerifyCode {
		fmt.Println(req.VerifyCode, code)
		return &pb.LoginResp{
			Code:    1,
			Message: "验证码不匹配",
		}, nil
	}
	// 二 判断电话号码是否注册 来获取顾客信息
	customer, err := s.cd.GetCustomerByTelephone(req.Telephone)
	if err != nil {
		return &pb.LoginResp{
			Code:    1,
			Message: "顾客信息获取错误",
		}, nil
	}

	// 三 设置token jwt-token
	token, err := s.cd.GenerateTokenAndSave(customer, biz.CustomerDuration*time.Second, biz.CustomerSecret)
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
	return &pb.LogoutResp{}, nil
}
