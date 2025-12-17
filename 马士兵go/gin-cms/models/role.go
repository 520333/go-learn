package models

import (
	"ginCms/utils"
)

// Role 定义角色模型
type Role struct {
	Model
	Title   string `gorm:"type:varchar(255);unique" json:"title"`
	Key     string `gorm:"type:varchar(255);unique" json:"key"`
	Enabled bool   `gorm:"" json:"enabled"`
	Weight  int    `gorm:"index;" json:"weight"`
	Comment string `gorm:"type:text" json:"comment"`
}

// RoleFilter 通用的请求过滤类型
type RoleFilter struct {
	// 指针类型表示该字段可以不填
	Keyword *string `form:"keyword" binding:"omitempty,gt=0"` //omitempty 非零值才校验
}

// Clean 整理Filter
func (f *RoleFilter) Clean() {
	if f.Keyword == nil {
		temp := ""
		f.Keyword = &temp
	}
}

// RoleFetchRow 根据条件查询单条 assoc 是否查询管理数据 where,args 查询条件
func RoleFetchRow(assoc bool, where any, args ...any) (*Role, error) {
	row := &Role{}
	if err := utils.DB().Where(where, args...).First(&row).Error; err != nil {
		return nil, err
	}
	// 关联查询
	if assoc {
	}
	return row, nil
}

func RoleFetchList(assoc bool, filter RoleFilter, sorter Sorter, pager Pager) ([]*Role, error) {
	return nil, nil
}

// 填充数据
func roleSeed() {
	// 构建数据
	rows := []Role{
		{
			Title:   "管理员",
			Key:     "administrator",
			Enabled: true,
			Model:   Model{ID: 1},
		},
		{
			Title:   "常规用户",
			Key:     "regular",
			Enabled: true,
			Model:   Model{ID: 2},
		},
	}
	// 插入
	for _, row := range rows {
		if err := utils.DB().FirstOrCreate(&row, row.ID).Error; err != nil {
			utils.Logger().Warn(err.Error())
		}
	}
}
