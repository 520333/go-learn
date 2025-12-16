package models

import (
	"ginCms/utils"
	"log"
)

func Init() {
	// migrate
	migrate()
	seed()
}

// 表结构迁移
func migrate() {
	// 自动迁移
	if err := utils.DB().AutoMigrate(&Role{}); err != nil {
		log.Fatalln(err)
	}
}

// 数据填充
func seed() {
	roleSeed()
}
