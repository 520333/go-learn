package service

import (
	"context"
	"driver/internal/biz"
	"log"
	"time"

	pb "driver/api/driver"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
)

type DriverService struct {
	pb.UnimplementedDriverServer
	Bz *biz.DriverBiz
}

func NewDriverService(bz *biz.DriverBiz) *DriverService {
	return &DriverService{
		Bz: bz,
	}
}

func (s *DriverService) IDNoCheck(ctx context.Context, req *pb.IDNoCheckReq) (*pb.IDNoCheckResp, error) {
	return &pb.IDNoCheckResp{}, nil
}

func (s *DriverService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	code, err := s.Bz.GetVerifyCode(ctx, req.Telephone)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: err.Error(),
		}, nil
	}
	return &pb.GetVerifyCodeResp{
		Code:           0,
		Message:        "SUCCESS",
		VerifyCode:     code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: 60,
	}, nil
}

func (s *DriverService) SubmitPhone(ctx context.Context, req *pb.SubmitPhoneReq) (*pb.SubmitPhoneResp, error) {
	// 校验验证码

	// 司机是否已经注册校验

	// 司机是否在黑名单中校验

	// 将司机信息入库，并设置状态为stop 暂时停用
	driver, err := s.Bz.InitDriverInfo(ctx, req.Telephone)
	if err != nil {
		return &pb.SubmitPhoneResp{
			Code:    1,
			Message: "司机号码提交失败",
		}, nil
	}

	return &pb.SubmitPhoneResp{
		Code:    0,
		Message: "SUCCESS",
		Status:  driver.Status.String,
	}, nil
}

// Login 登录 Service
func (s *DriverService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	token, err := s.Bz.CheckLogin(ctx, req.Telephone, req.VerifyCode)
	if err != nil {
		log.Println(err)
		return &pb.LoginResp{
			Code:    1,
			Message: "司机登录失败",
		}, nil
	}
	return &pb.LoginResp{
		Code:           0,
		Message:        "SUCCESS",
		Token:          token,
		TokenCreatedAt: time.Now().Unix(),
		TokenLife:      biz.DriverTokenLife,
	}, nil
}

func (s *DriverService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
	// 一 获取用户id
	claims, _ := jwt.FromContext(ctx)
	claimsMap := claims.(jwt2.MapClaims)
	if err := s.Bz.DelToken(claimsMap["jti"]); err != nil {
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
