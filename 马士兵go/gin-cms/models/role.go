package models

import "ginCms/utils"

// Role 定义角色模型
type Role struct {
	Model
	Title   string `gorm:"type:varchar(255);unique" json:"title"`
	Key     string `gorm:"type:varchar(255);unique" json:"key"`
	Enabled bool   `gorm:"" json:"enabled"`
	Weight  int    `gorm:"index;" json:"weight"`
	Comment string `gorm:"type:text" json:"comment"`
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
