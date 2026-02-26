package service

import (
	"context"
	"math/rand"

	pb "verifyCode/api/verityCode"
)

type VerityCodeService struct {
	pb.UnimplementedVerityCodeServer
}

func NewVerityCodeService() *VerityCodeService {
	return &VerityCodeService{}
}

func (s *VerityCodeService) GetVerityCode(ctx context.Context, req *pb.GetVerityCodeRequest) (*pb.GetVerityCodeReply, error) {
	return &pb.GetVerityCodeReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}

func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		return randCode("0123456789", l)
	case pb.TYPE_LETTER:
		return randCode("abcdefghijklmnopqrstuvwxyz", l)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijklmnopqrstuvwxyz", l)
	default:
	}
	return ""
}

func randCode(chars string, l int) string {
	charsLen := len(chars)

	result := make([]byte, l)
	for i := 0; i < l; i++ {
		randIndex := rand.Intn(charsLen)
		result[i] = chars[randIndex]
	}

	return string(result)
}
