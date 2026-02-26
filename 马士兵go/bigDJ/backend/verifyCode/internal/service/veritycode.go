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
		return randCode("0123456789", l, 4)
	case pb.TYPE_LETTER:
		return randCode("abcdefghijklmnopqrstuvwxyz", l, 5)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijklmnopqrstuvwxyz", l, 6)
	default:
	}
	return ""
}

// 随机数核心方法（优化）
func randCode(chars string, l, idxBits int) string {
	//idxBits = len(fmt.Sprintf("%b", len(chars)))
	idxMask := 1<<idxBits - 1
	idxMax := 63 / idxBits
	result := make([]byte, l)
	// 生成随机数字符
	for i, cache, remain := 0, rand.Int63(), idxMax; i < l; {
		if 0 == remain {
			cache, remain = rand.Int63(), idxMax
		}
		if randIndex := int(cache & int64(idxMask)); randIndex < len(chars) {
			result[i] = chars[randIndex]
			i++
		}
		cache >>= idxBits
		remain--
	}
	return string(result)
}

//func randCode(chars string, l int) string {
//	charsLen := len(chars)
//
//	result := make([]byte, l)
//	for i := 0; i < l; i++ {
//		randIndex := rand.Intn(charsLen)
//		result[i] = chars[randIndex]
//	}
//
//	return string(result)
//}
