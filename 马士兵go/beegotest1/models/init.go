package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var alias = "default"
	var driverName = "mysql"
	dataSource := "root:123456@tcp(192.168.50.100:3306)/beego_test?charset=utf8mb4&loc=Local"
	maxOpenConnections := 100
	maxIdleConnections := 20
	// 注册
	if err := orm.RegisterDataBase(alias, driverName, dataSource,
		orm.MaxOpenConnections(maxOpenConnections),
		orm.MaxIdleConnections(maxIdleConnections)); err != nil {
		log.Println(err)
	}

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true // 开启ORM调试日志
	}

}
