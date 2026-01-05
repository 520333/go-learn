package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Img     string
	Desc    string
}

func (c *Community) TableName() string {
	return "community"
}

func CreateCommunity(c Community) (int, string) {
	tx := utils.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if len(c.Name) == 0 {
		return -1, "群名称不能为空"
	}
	if c.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := utils.DB.Create(&c).Error; err != nil {
		fmt.Println(err)
		tx.Rollback()
		return -1, "建群失败"
	}
	contact := Contact{}
	contact.OwnerId = c.OwnerId
	contact.TargetId = c.ID
	contact.Type = 2
	if err := tx.Create(&contact).Error; err != nil {
		tx.Rollback()
		return -1, "添加群关系失败"
	}

	tx.Commit()
	return 0, "建群成功"
}

func LoadCommunity(ownerId uint) ([]*Community, string) {
	contacts := make([]Contact, 0)
	objIds := make([]uint64, 0)
	utils.DB.Where("owner_id=? and type = 2", ownerId).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, uint64(v.TargetId))
	}
	data := make([]*Community, 10)
	utils.DB.Where("id in (?)", objIds).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data, "查询成功"
}
