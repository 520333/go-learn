package biz

import (
	"context"
	"valuation/api/mapService"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

type ValuationBiz struct{}

func NewValuationBiz() *ValuationBiz {
	return &ValuationBiz{}
}

// GetDrivingInfo 获取市场和距离
func (*ValuationBiz) GetDrivingInfo(ctx context.Context, origin, destination string) (distance string, duration string, err error) {
	// 一 发出grpc请求
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.1.178:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return
	}
	// 2.服务发现管理器
	dis := consul.New(consulClient)

	// 2.1连接目标grpc服务器
	endpoint := "discovery:///map"
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(endpoint), // 目标服务的名字
		grpc.WithDiscovery(dis),     // 使用服务发现
	)

	if err != nil {
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	// 2.2 发送获取驾驶距离和市场请求，RPC调用
	client := mapService.NewMapServiceClient(conn)
	reply, err := client.GetDrivingInfo(context.Background(), &mapService.GetDrivingInfoRequest{
		Origin:      origin,
		Destination: destination,
	})
	if err != nil {
		return
	}
	distance, duration = reply.Distance, reply.Duration
	return
}
