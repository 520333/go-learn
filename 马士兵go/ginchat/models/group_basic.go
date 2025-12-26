package models

import "gorm.io/gorm"

// GroupBasic 群信息
type GroupBasic struct {
	gorm.Model
	Name    string // 群名
	OwnerId uint   // 拥有者
	Icon    string // 图标
	Type    int    // 群类型
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
