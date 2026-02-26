package service

import (
	"context"

	pb "verifyCode/api/verityCode"
)

type VerityCodeService struct {
	pb.UnimplementedVerityCodeServer
}

func NewVerityCodeService() *VerityCodeService {
	return &VerityCodeService{}
}

func (s *VerityCodeService) GetVerityCode(ctx context.Context, req *pb.GetVerityCodeRequest) (*pb.GetVerityCodeReply, error) {
	return &pb.GetVerityCodeReply{}, nil
}
