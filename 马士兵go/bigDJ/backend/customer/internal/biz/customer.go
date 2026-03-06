package biz

import (
	"context"
	"customer/api/valuation"
	"database/sql"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
)

const (
	CustomerSecret   = "bigDJ"
	CustomerDuration = 2 * 30 * 24 * 3600
)

type Customer struct {
	CustomerWork
	CustomerToken
	gorm.Model
}

type CustomerWork struct {
	Telephone string         `gorm:"type:varchar(15);uniqueIndex;" json:"telephone"`
	Name      sql.NullString `gorm:"type:varchar(255);uniqueIndex"  json:"name"`
	Email     sql.NullString `gorm:"type:varchar(255);uniqueIndex"  json:"email"`
	Wechat    sql.NullString `gorm:"type:varchar(255);uniqueIndex" json:"wechat"`
	CityID    uint           `gorm:"index;" json:"city_id"`
}

type CustomerToken struct {
	Token          string       `gorm:"type:varchar(4095)"  json:"token"`
	TokenCreatedAt sql.NullTime `gorm:""  json:"token_created_at"`
}

type CustomerBiz struct {
}

func NewCustomerBiz() *CustomerBiz {
	return &CustomerBiz{}
}

func (cb *CustomerBiz) GetEstimatePrice(origin, destination string) (int64, error) {
	// 一 发出grpc请求
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "192.168.1.178:8500"
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return 0, err
	}
	// 2.服务发现管理器
	dis := consul.New(consulClient)

	// 2.1连接目标grpc服务器
	endpoint := "discovery:///Valuation"
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(endpoint), // 目标服务的名字
		grpc.WithDiscovery(dis),     // 使用服务发现
	)

	if err != nil {
		return 0, err
	}
	defer func() {
		_ = conn.Close()
	}()
	client := valuation.NewValuationClient(conn)
	reply, err := client.GetEstimatePrice(context.Background(),
		&valuation.GetEstimatePriceReq{
			Origin:      origin,
			Destination: destination,
		})
	if err != nil {
		return 0, err
	}

	return reply.Price, nil
}
