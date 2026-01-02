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

func AddFriend(userId uint, targetId uint) int {
	user := UserBasic{}
	if targetId != 0 {
		user = FindById(targetId)
		if user.Identity != "" {
			tx := utils.DB.Begin()
			var contact = Contact{}
			contact.OwnerId = userId
			contact.TargetId = targetId
			contact.Type = 1
			utils.DB.Create(&contact)

			var contact2 = Contact{}
			contact2.OwnerId = targetId
			contact2.TargetId = userId
			contact2.Type = 1
			utils.DB.Create(&contact2)
			tx.Commit()
			return 0
		}
		return -1
	}
	return -1
}
