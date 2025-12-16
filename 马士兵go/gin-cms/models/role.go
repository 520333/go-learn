package models

// Role 定义角色模型
type Role struct {
	Model
	Title   string `gorm:"type:varchar(255);unique" json:"title"`
	Key     string `gorm:"type:varchar(255);unique" json:"key"`
	Enabled bool   `gorm:"" json:"enabled"`
	Weight  int    `gorm:"index;" json:"weight"`
	Comment string `gorm:"type:text" json:"comment"`
}
