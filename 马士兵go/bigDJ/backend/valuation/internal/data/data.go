package data

import (
	"valuation/internal/biz"
	"valuation/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewPriceRuleInterface)

// Data .
type Data struct {
	Mdb *gorm.DB // 操作mysql客户端
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	data := &Data{}
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		data.Mdb = nil
		log.Fatalf(err.Error())
	}
	data.Mdb = db
	// 3 自动迁移表
	migrateTable(db)
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return data, cleanup, nil
}
func migrateTable(db *gorm.DB) {
	if err := db.AutoMigrate(&biz.PriceRule{}); err != nil {
		log.Info("auto migrate valuation table err", err)
	}
	// 插入一些priceRule测试数据
	rules := []biz.PriceRule{
		{
			Model: gorm.Model{ID: 1},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      1,
				StartFee:    300,
				DistanceFee: 35,
				DurationFee: 10,
				StartAt:     7,
				EndAt:       23,
			},
		},
		{
			Model: gorm.Model{ID: 2},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      1,
				StartFee:    350,
				DistanceFee: 35,
				DurationFee: 10,
				StartAt:     23,
				EndAt:       24,
			},
		},
		{
			Model: gorm.Model{ID: 3},
			PriceRuleWork: biz.PriceRuleWork{
				CityID:      1,
				StartFee:    400,
				DistanceFee: 35,
				DurationFee: 10,
				StartAt:     0,
				EndAt:       7,
			},
		},
	}
	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&rules)
}
