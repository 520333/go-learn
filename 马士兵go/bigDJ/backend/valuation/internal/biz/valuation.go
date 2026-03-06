package biz

import (
	"context"
	"strconv"
	"valuation/api/mapService"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
)

type PriceRule struct {
	gorm.Model
	PriceRuleWork
}

type PriceRuleWork struct {
	CityID      uint  `gorm:"" json:"city_id"`
	StartFee    int64 `gorm:"" json:"start_fee"`
	DistanceFee int64 `gorm:"" json:"distance_fee"`
	DurationFee int64 `gorm:"" json:"duration_fee"`
	StartAt     int   `gorm:"type:int" json:"start_at"` // 0-[0
	EndAt       int   `gorm:"type:int" json:"end_at"`   // 7 0
}

// PriceRuleInterface 定义操作priceRule的接口
type PriceRuleInterface interface {
	// GetRule 获取规则
	GetRule(cityid uint, curr int) (*PriceRule, error)
}

type ValuationBiz struct {
	pri PriceRuleInterface
}

func NewValuationBiz(pri PriceRuleInterface) *ValuationBiz {
	return &ValuationBiz{
		pri: pri,
	}
}

func (vb *ValuationBiz) GetPrice(ctx context.Context, distance, duration string, cityid uint, curr int) (int64, error) {
	// 一 获取规则
	rule, err := vb.pri.GetRule(cityid, curr)
	if err != nil {
		return 0, err
	}
	// 将距离和时长转换为int64
	distanceInt64, err := strconv.ParseInt(distance, 10, 64)
	if err != nil {
		return 0, err
	}
	durationInt64, err := strconv.ParseInt(duration, 10, 64)
	if err != nil {
		return 0, err
	}

	// 二 基于rule计算
	distanceInt64 /= 1000
	durationInt64 /= 60
	var startDistance int64 = 5

	var distancePrice int64
	if distanceInt64 > startDistance {
		distancePrice = rule.DistanceFee * (distanceInt64 - startDistance)
	}
	total := rule.StartFee +
		distancePrice +
		rule.DurationFee*durationInt64
	//total := rule.StartFee +
	//	rule.DistanceFee*(distanceInt64-startDistance) +
	//	rule.DurationFee*durationInt64
	return total, nil
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
	endpoint := "discovery:///Map"
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(endpoint),           // 目标服务的名字
		grpc.WithDiscovery(dis),               // 使用服务发现
		grpc.WithMiddleware(tracing.Client()), // 客户端tracing
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
