package service

import (
	"context"

	pb "valuation/api/valuation"
)

type ValuationService struct {
	pb.UnimplementedValuationServer
}

func NewValuationService() *ValuationService {
	return &ValuationService{}
}

func (s *ValuationService) GetEstimatePrice(ctx context.Context, req *pb.GetEstimatePriceReq) (*pb.GetEstimatePriceReply, error) {
	return &pb.GetEstimatePriceReply{}, nil
}
