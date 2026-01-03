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
	if len(c.Name) == 0 {
		return -1, "群名称不能为空"
	}
	if c.OwnerId == 0 {
		return -1, "请先登录"
	}
	if err := utils.DB.Create(&c).Error; err != nil {
		fmt.Println(err)
		return -1, "建群失败"
	}
	return 0, "建群成功"
}

func LoadCommunity(ownerId uint) ([]*Community, string) {
	data := make([]*Community, 0)
	utils.DB.Where("owner_id = ?", ownerId).Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}

	return data, "查询成功"
}
