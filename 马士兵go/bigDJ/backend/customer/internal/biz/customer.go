package biz

import (
	"database/sql"

	"gorm.io/gorm"
)

const (
	CustomerSecret   = "bigDJ"
	CustomerDuration = 2 * 30 * 24 * 3600
)

type Customer struct {
	CustomerWork
	CustomerToken
	gorm.Model
}

type CustomerWork struct {
	Telephone string         `gorm:"type:varchar(15);uniqueIndex;" json:"telephone"`
	Name      sql.NullString `gorm:"type:varchar(255);uniqueIndex"  json:"name"`
	Email     sql.NullString `gorm:"type:varchar(255);uniqueIndex"  json:"email"`
	Wechat    sql.NullString `gorm:"type:varchar(255);uniqueIndex" json:"wechat"`
	CityID    uint           `gorm:"index;" json:"city_id"`
}

type CustomerToken struct {
	Token          string       `gorm:"type:varchar(4095)"  json:"token"`
	TokenCreatedAt sql.NullTime `gorm:""  json:"token_created_at"`
}
