package gorm

import (
	"fmt"
	"log"
)

func RawSelect() {
	type Result struct {
		ID           uint
		Subject      string
		Likes, Views int
	}
	var rs []Result
	sql := "SELECT `id`,`subject`,`likes`,`views` FROM go_content WHERE `likes` > ? ORDER BY `likes` DESC LIMIT ?"
	// select `id`,`subject`,`likes`,`views` From go_content WHERE `likes` > 99 ORDER BY `likes` DESC LIMIT 12
	// 执行
	if err := DB.Raw(sql, 99, 12).Scan(&rs).Error; err != nil {
		log.Println(err)
	}
	log.Println(rs)
}

// RawExec 执行类
func RawExec() {
	// SQL
	sql := "UPDATE `go_content` SET `subject` = CONCAT(`subject`,'-newappfix')  WHERE `id` BETWEEN ? AND ?"

	// UPDATE `go_content` SET `subject` = CONCAT(`subject`,'-newappfix')  WHERE `id` BETWEEN 10 AND 20;
	// 执行获取结果
	result := DB.Exec(sql, 10, 20)
	if result.Error != nil {
		log.Fatalln(result.Error)
	}
	log.Println(result.RowsAffected)

}

// sql.row或者sql.rows类型的结果处理
func RowAndRows() {
	sql := "SELECT `id`,`subject`,`likes`,`views` FROM go_content WHERE `likes` > ? ORDER BY `likes` DESC LIMIT ?"

	rows, err := DB.Raw(sql, 99, 12).Rows()
	if err != nil {
		log.Fatalln(err)
	}
	// 遍历
	for rows.Next() {
		// 扫描到列独立的变量
		var id uint
		var subject string
		var likes, views int
		_ = rows.Scan(&id, &subject, &likes, &views)
		fmt.Println(id, subject, likes, views)

		// 扫描到结构体
		type Result struct {
			ID           uint
			Subject      string
			Likes, Views int
		}
		var r Result
		err := DB.ScanRows(rows, &r)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(r)
	}
}
