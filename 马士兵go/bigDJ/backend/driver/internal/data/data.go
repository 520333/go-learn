package data

import (
	"driver/internal/biz"
	"driver/internal/conf"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDriverInterface)

// Data .
type Data struct {
	Mdb *gorm.DB
	Rdb *redis.Client
	cs  *conf.Service
}

// NewData .
func NewData(c *conf.Data, cs *conf.Service, logger log.Logger) (*Data, func(), error) {
	data := &Data{
		cs: cs,
	}
	redisURL := fmt.Sprintf("redis://%s/1?dial_timeout=%d", c.Redis.Addr, 1)
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		data.Rdb = nil
		log.Fatal(err)
	}
	data.Rdb = redis.NewClient(options)

	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		data.Rdb = nil
		log.Fatal(err)
	}
	data.Mdb = db
	migrateTable(db)
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return data, cleanup, nil
}

func migrateTable(db *gorm.DB) {
	if err := db.AutoMigrate(&biz.Driver{}); err != nil {
		log.Fatal(err)
	}
}
