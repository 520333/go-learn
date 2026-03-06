package service

import (
	"context"
	"map/internal/biz"

	pb "map/api/mapService"

	"github.com/go-kratos/kratos/v2/errors"
)

type MapServiceService struct {
	pb.UnimplementedMapServiceServer
	msBiz *biz.MapServiceBiz
}

func NewMapServiceService() *MapServiceService {
	return &MapServiceService{}
}

func NewMapServiceServiceBiz(msBiz *biz.MapServiceBiz) *MapServiceService {
	return &MapServiceService{
		msBiz: msBiz,
	}
}

func (s *MapServiceService) GetDrivingInfo(ctx context.Context, req *pb.GetDrivingInfoRequest) (*pb.GetDrivingInfoReply, error) {
	distance, duration, err := s.msBiz.GetDriverInfo(req.Origin, req.Destination)
	if err != nil {
		return nil, errors.New(200, "LBS_ERROR", "lbs api error")
	}
	return &pb.GetDrivingInfoReply{
		Origin:      req.Origin,
		Destination: req.Destination,
		Distance:    distance,
		Duration:    duration,
	}, nil
}
