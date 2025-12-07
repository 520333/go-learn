package gorm

import "log"

func TXDemo() {
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}
	var a1, a2 Author
	a1.Name = "库里"
	a2.Name = "莫兰特"
	a1.Points = 1600
	a2.Points = 200
	if err := DB.Create([]*Author{&a1, &a2}).Error; err != nil {
		log.Fatalln(err)
	}
	// 事务操作 a1赠送a2 2000积分
	var p = 1000
	// 开始事务
	tx := DB.Begin()
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
	// 赠送操作
	a1.Points -= p
	if err := tx.Save(&a1).Error; err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalln(err)
		}
	}

	a2.Points += p
	if err := tx.Save(&a2).Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			log.Fatalln(err)
		}
	}
	// 业务逻辑：要求author的积分不能为负数
	if a1.Points < 0 || a2.Points < 0 {
		log.Println("a1.Points < 0 || a2.Points < 0")
		if err := tx.Rollback().Error; err != nil {
			log.Fatalln(err)
		}
	}
	if err := tx.Commit().Error; err != nil {
		log.Fatalln(err)
	}
	// 决定回滚还是提交
	//if err1 != nil || err2 != nil {
	//	tx.Rollback()
	//} else {
	//	tx.Commit()
	//}
}
