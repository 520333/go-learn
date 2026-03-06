package service

import (
	"context"
	"fmt"
	"valuation/internal/biz"

	pb "valuation/api/valuation"

	"github.com/go-kratos/kratos/v2/errors"
)

type ValuationService struct {
	pb.UnimplementedValuationServer
	vb *biz.ValuationBiz
}

func NewValuationService(vb *biz.ValuationBiz) *ValuationService {
	return &ValuationService{vb: vb}
}

func (s *ValuationService) GetEstimatePrice(ctx context.Context, req *pb.GetEstimatePriceReq) (*pb.GetEstimatePriceReply, error) {
	distance, duration, err := s.vb.GetDrivingInfo(ctx, req.Origin, req.Destination)
	if err != nil {
		return nil, errors.New(200, "Map ERROR", "get driving info")
	}
	fmt.Println(distance, duration)
	return &pb.GetEstimatePriceReply{}, nil
}
