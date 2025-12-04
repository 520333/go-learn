package gorm

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func UpdatePK() {
	var c Content
	// INSERT INTO `go_content` (`created_at`,`updated_at`,`deleted_at`,`subject`,`likes`,`views`,`publish_time`,`author_id`) VALUES ('2025-12-05 01:05:25.663','2025-12-05 01:05:25.663',NULL,'',0,0,'2025-12-05 01:05:25.662',0) ON DUPLICATE KEY UPDATE `updated_at`='2025-12-05 01:05:25.663',`deleted_at`=VALUES(`deleted_at`),`subject`=VALUES(`subject`),`likes`=VALUES(`likes`),`views`=VALUES(`views`),`publish_time`=VALUES(`publish_time`),`author_id`=VALUES(`author_id`)
	if err := DB.Save(&c).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", c)

	// UPDATE `go_content` SET `created_at`='2025-12-05 01:10:04.527',`updated_at`='2025-12-05 01:10:04.539',`deleted_at`=NULL,`subject`='',`likes`=0,`views`=0,`publish_time`='2025-12-05 01:10:04.526',`author_id`=0 WHERE `go_content`.`deleted_at` IS NULL AND `id` = 27
	if err := DB.Save(&c).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", c)
}

func UpdateWhere() {
	//UPDATE `go_content` SET `likes`=10001,`subject`='Where Update Row',`updated_at`='2025-12-05 01:20:19.73' WHERE likes > 100 AND `go_content`.`deleted_at` IS NULL
	// 更新的字段值数据 推荐map结构
	values := map[string]any{
		"subject": "Where Update Row",
		"likes":   10001,
	}
	// 执行带有条件的更新
	result := DB.Model(&Content{}).
		//Omit("updated_at").
		Where("likes > ?", 100).
		Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	// 获取更新结果，更新的记录数量(受影响的ROWS)
	log.Println("Update rows num:", result.RowsAffected)
}

func UpdateNoWhere() {
	values := map[string]any{
		"subject": "Where Update Row",
		"likes":   1001,
	}
	// 执行带有条件的更新 1=1 全局更新
	result := DB.Model(&Content{}).Where("1 = 1").Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	// 获取更新结果，更新的记录数量(受影响的ROWS)
	log.Println("Update rows num:", result.RowsAffected)
}

func UpdateExpr() {
	// UPDATE `go_content` SET `likes`=likes + 10,`subject`='Where Update Row',`updated_at`='2025-12-05 01:35:56.448' WHERE likes > 100 AND `go_content`.`deleted_at` IS NULL
	// 更新的字段值数据 推荐map结构
	values := map[string]any{
		"subject": "Where Update Row",
		"likes":   gorm.Expr("likes + ?", 10),
	}
	// 执行带有条件的更新
	result := DB.Model(&Content{}).
		Where("likes > ?", 100).
		Updates(values)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	// 获取更新结果，更新的记录数量(受影响的ROWS)
	log.Println("Update rows num:", result.RowsAffected)
}
