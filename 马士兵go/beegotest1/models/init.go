package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 1.注册数据库

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

	// 2.注册表
	if beego.BConfig.RunMode == "dev" {
		orm.RegisterModel(&Article{})
		//orm.RegisterModel(&Article{},&other{})
		//orm.RegisterModelWithPrefix("beegotest_", &Article{})
	}

	// 3.自动同步表结构
	if beego.BConfig.RunMode == "dev" {
		dbName, force, verbose := "default", true, false
		if err := orm.RunSyncdb(dbName, force, verbose); err != nil {
			log.Println(err)
		}
	}

	// 4.初始化Ormer
	OrmDft = orm.NewOrm()
	//OrmDft = orm.NewOrmWithDB("default")
}

var OrmDft orm.Ormer
