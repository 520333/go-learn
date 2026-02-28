package biz

import (
	"database/sql"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerWork
	CustomerToken
	gorm.Model
}

type CustomerWork struct {
	Telephone string `gorm:"type:varchar(15);uniqueIndex;" json:"telephone"`
	Name      string `gorm:"type:varchar(255);uniqueIndex"  json:"name"`
	Email     string `gorm:"type:varchar(255);uniqueIndex"  json:"email"`
	Wechat    string `gorm:"type:varchar(255);uniqueIndex" json:"wechat"`
}

type CustomerToken struct {
	Token          string       `gorm:"type:varchar(4095)"  json:"token"`
	TokenCreatedAt sql.NullTime `gorm:""  json:"token_created_at"`
}
