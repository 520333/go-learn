package gorm

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

func SessionIssue() {
	// SELECT * FROM `go_content` WHERE views > 100 AND likes > 9 AND `go_content`.`deleted_at` IS NULL
	// 连续查询
	db := DB.Model(&Content{}).Where("views > ?", 100)
	db.Where("likes > ?", 9)
	var cs []Content
	db.Find(&cs)

	// SELECT * FROM `go_content` WHERE views > 100 AND likes > 9 AND `go_content`.`deleted_at` IS NULL AND likes > 19
	db.Where("likes > ?", 19)
	var cs2 []Content
	db.Find(&cs2)
}

func SessionDB() {
	// 连续查询
	// SELECT * FROM `go_content` WHERE views > 100 AND likes > 9 AND `go_content`.`deleted_at` IS NULL
	db1 := DB.Model(&Content{}).Where("views > ?", 100)
	db1.Where("likes > ?", 9)
	var cs []Content
	db1.Find(&cs)

	// SELECT * FROM `go_content` WHERE views > 100 AND likes > 19 AND `go_content`.`deleted_at` IS NULL
	db2 := DB.Model(&Content{}).Where("views > ?", 100)
	db2.Where("likes > ?", 19)
	var cs2 []Content
	db2.Find(&cs2)
}

func SessionNew() {
	// 连续查询
	// SELECT * FROM `go_content` WHERE views > 100 AND likes > 9 AND `go_content`.`deleted_at` IS NULL
	db := DB.Model(&Content{}).Where("views > ?", 100).Session(&gorm.Session{})
	var cs []Content
	db.Where("likes > ?", 9).Find(&cs)

	// SELECT * FROM `go_content` WHERE views > 100 AND likes > 19 AND `go_content`.`deleted_at` IS NULL
	var cs2 []Content
	db.Where("likes > ?", 19).Find(&cs2)
}

func SessionOptions() {
	//db := DB.Session(&gorm.Session{
	//	SkipHooks: true,
	//})
	//db.Save(&Content{Subject: "NO CREATE HOOK"})

	//db := DB.Session(&gorm.Session{DryRun: true})
	//stmt := db.Save(&Content{}).Statement

	// INSERT INTO `go_content` (`created_at`,`updated_at`,`deleted_at`,`subject`,`likes`,`views`,`publish_time`,`author_id`) VALUES (?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE `updated_at`=?,`deleted_at`=VALUES(`deleted_at`),`subject`=VALUES(`subject`),`likes`=VALUES(`likes`),`views`=VALUES(`views`),`publish_time`=VALUES(`publish_time`),`author_id`=VALUES(`author_id`)
	//fmt.Println(stmt.SQL.String()) //输出执行语句
	//fmt.Println(stmt.Vars)

	db1 := DB.Session(&gorm.Session{
		PrepareStmt: true,
	})
	stmManager, ok := db1.ConnPool.(*gorm.PreparedStmtDB)
	if !ok {
		log.Println("*gorm.PreparedStmtDB assert failed")
	}
	fmt.Println(stmManager.Stmts)
	var c1 Content
	db1.First(&c1, 13)
	fmt.Println(stmManager.Stmts)
	var c2 Content
	db1.First(&c2, 14)
	var c3 Content
	db1.First(&c3, 15)
	keys := stmManager.Stmts.Keys()
	fmt.Println("Prepared SQL count:", len(keys))
	for _, sql := range keys {
		fmt.Println("Prepared SQL:", sql)
	}
	var c4 Content
	db1.Find(&c4, []uint{15, 16})
	fmt.Println("Prepared SQL count:", stmManager.Stmts.Keys())

}
