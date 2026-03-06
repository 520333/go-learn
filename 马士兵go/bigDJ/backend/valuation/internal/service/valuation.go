package service

import (
	"context"
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
	//fmt.Println(distance, duration)
	// 费用
	price, err := s.vb.GetPrice(ctx, distance, duration, 1, 23)
	if err != nil {
		return nil, errors.New(200, "PRICE ERROR", "call price error")
	}
	return &pb.GetEstimatePriceReply{
		Origin:      req.Origin,
		Destination: req.Destination,
		Price:       price,
	}, nil
}
