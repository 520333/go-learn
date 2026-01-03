package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

// Contact 人员关系
type Contact struct {
	gorm.Model
	OwnerId  uint // 谁的关系
	TargetId uint // 对应的谁
	Type     int  // 对应的类型 1:好友 2:群组 3:广播
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}

func SearchFriend(userId uint) []UserBasic {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id = ? and type = 1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(v)
		objIds = append(objIds, uint64(v.TargetId))
	}
	users := make([]UserBasic, 0)
	utils.DB.Where("id in (?)", objIds).Find(&users)
	//fmt.Println(users)
	//for _, v := range users {
	//	fmt.Println("好友:", v.Name)
	//}
	return users
}

func AddFriend(userId uint, targetId uint) (int, string) {
	user := UserBasic{}
	if targetId != 0 {
		user = FindById(targetId)
		if user.Identity != "" {
			if userId == user.ID {
				return -1, "无法添加自己"
			}
			contact0 := Contact{}

			utils.DB.Where("owner_id = ? and target_id = ? and type = 1", userId, targetId).Find(&contact0)
			if contact0.ID != 0 {
				return -1, "不能重复添加"
			}
			tx := utils.DB.Begin()
			// 事务一旦开始无论期间什么异常都会rollback
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()
			var contact = Contact{}
			contact.OwnerId = userId
			contact.TargetId = targetId
			contact.Type = 1
			if err := utils.DB.Create(&contact).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}

			var contact2 = Contact{}
			contact2.OwnerId = targetId
			contact2.TargetId = userId
			contact2.Type = 1
			if err := utils.DB.Create(&contact2).Error; err != nil {
				tx.Rollback()
				return -1, "添加好友失败"
			}
			tx.Commit()
			return 0, "添加好友成功"
		}
		return -1, "用户未找到"
	}
	return -1, "好友ID不能为空"
}
