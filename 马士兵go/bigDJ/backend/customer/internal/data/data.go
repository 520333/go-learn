package data

import (
	"customer/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewCustomerData)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	data := &Data{}
	// 1.初始化Rdb
	// 连接redis 使用服务配置
	redisUrl := fmt.Sprintf("redis://%s/1?dial_timeout=%d", c.Redis.Addr, 1)
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		data.Rdb = nil
	}
	data.Rdb = redis.NewClient(opt)

	cleanup := func() {
		_ = data.Rdb.Close() // 清理Redis连接
		log.Info("closing the data resources")
	}
	return data, cleanup, nil
}
