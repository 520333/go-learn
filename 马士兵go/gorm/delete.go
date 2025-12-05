package gorm

import (
	"fmt"
	"log"
)

func DeleteWhere() {
	// 逻辑软删除
	// UPDATE `go_content` SET `deleted_at`='2025-12-06 00:10:24.941' WHERE likes < 100 AND `go_content`.`deleted_at` IS NULL
	result := DB.Delete(&Content{}, "likes < ?", 100)
	if err := result.Error; err != nil {
		log.Fatal(err)
	}
	// UPDATE `go_content` SET `deleted_at`='2025-12-06 00:10:24.95' WHERE likes < 100 AND `go_content`.`deleted_at` IS NULL
	result2 := DB.Where("likes < ?", 100).Delete(&Content{})
	if err := result2.Error; err != nil {
		log.Fatal(err)
	}
}

func FindDeleted() {
	var c Content
	// SELECT * FROM `go_content` WHERE `go_content`.`id` = 13 AND `go_content`.`deleted_at` IS NULL ORDER BY `go_content`.`id` LIMIT 1

	DB.Delete(&c, 13)
	// SELECT * FROM `go_content` WHERE `go_content`.`id` = 13 ORDER BY `go_content`.`id` LIMIT 1
	if err := DB.Unscoped().First(&c, 13).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}

func DeleteHard() {
	var c Content
	// DELETE FROM `go_content` WHERE `go_content`.`id` = 14
	if err := DB.Unscoped().Delete(&c, 14).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}
