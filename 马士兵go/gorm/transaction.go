package gorm

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

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

func TXCallback() {
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
	log.Println(a1.ID, a2.ID)

	if err := DB.Transaction(func(tx *gorm.DB) error {
		var p = 200
		// 赠送操作
		a1.Points -= p
		a2.Points += p
		if err := tx.Save(&a1).Error; err != nil {
			return err
		}
		if err := tx.Save(&a2).Error; err != nil {
			return err
		}
		// 业务逻辑：要求author的积分不能为负数
		if a1.Points < 0 || a2.Points < 0 {
			return errors.New("a1.Points < 0 || a2.Points < 0")
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}
}

func TXNested() {
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}
	var a1, a2, a3 Author
	a1.Name = "库里"
	a2.Name = "莫兰特"
	a3.Name = "欧文"

	a1.Points = 1600
	a2.Points = 200
	a3.Points = 4000
	if err := DB.Create([]*Author{&a1, &a2, &a3}).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println(a1.ID, a2.ID, a3.ID)

	if err := DB.Transaction(func(tx *gorm.DB) error {
		var p = 20000
		// 赠送操作
		a2.Points += p
		if err := tx.Save(&a2).Error; err != nil {
			return err
		}
		// a1赠送
		errA1 := DB.Transaction(func(tx *gorm.DB) error {
			a1.Points -= p
			if err := tx.Save(&a1).Error; err != nil {
				return err
			}
			if a1.Points < 0 {
				return errors.New("a1.Points < 0")
			}
			return nil
		})
		if errA1 != nil {
			errA3 := DB.Transaction(func(tx *gorm.DB) error {
				// a3赠送
				a3.Points -= p
				if err := tx.Save(&a3).Error; err != nil {
					return err
				}
				if a3.Points < 0 {
					return errors.New("a3.Points < 0")
				}
				return nil
			})
			if errA3 != nil {
				return errors.New("a1 and a3 all send points failed")
			}
		}
		return nil
	}); err != nil {
		log.Fatalln(err)
	}
}

func TXSavePoint() {
	if err := DB.AutoMigrate(&Author{}); err != nil {
		log.Fatalln(err)
	}
	var a1, a2, a3 Author
	a1.Name = "库里"
	a2.Name = "莫兰特"
	a3.Name = "欧文"

	a1.Points = 1600
	a2.Points = 200
	a3.Points = 4000
	if err := DB.Create([]*Author{&a1, &a2, &a3}).Error; err != nil {
		log.Fatalln(err)
	}
	log.Println(a1.ID, a2.ID, a3.ID)

	// 事务操作 a1赠送a2 2000积分
	var p = 5000
	// 开始事务
	tx := DB.Begin()
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
	// 赠送操作
	a2.Points += p
	if err := tx.Save(&a2).Error; err != nil {
		tx.Rollback()
		return
	}
	// 逻辑记录发送points是否成功
	var flagSend bool
	// a1先给a2 send
	tx.SavePoint("beforeA1") // 设置一个savepoint
	a1.Points -= p
	if err := tx.Save(&a1).Error; err != nil || a1.Points < 0 {
		tx.RollbackTo("beforeA1") //回滚到A1
		// a3 to a2
		tx.SavePoint("beforeA3")
		a3.Points -= p
		if err := tx.Save(&a3).Error; err != nil || a3.Points < 0 {
			tx.RollbackTo("beforeA3") //回滚到A3
		} else {
			flagSend = true
		}
	} else {
		flagSend = true
	}
	// 判断赠送积分是否成功
	if flagSend {
		if err := tx.Commit().Error; err != nil {
			log.Fatalln(err)
		}
	} else {
		tx.Rollback() // 回滚事务
	}
}
