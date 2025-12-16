package models

import (
	"ginCms/utils"
	"log"
)

func Init() {
	// migrate
	migrate()

}

// 表结构迁移
func migrate() {
	// 自动迁移
	if err := utils.DB().AutoMigrate(&Role{}); err != nil {
		log.Fatalln(err)
	}

}
