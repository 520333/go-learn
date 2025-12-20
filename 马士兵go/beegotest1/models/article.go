package models

import "time"

type Article struct {
	// 基础字段
	Id        uint64     `orm:"description(自增主键)"`
	CreatedAt time.Time  `orm:"auto_now_add"`
	UpdatedAt time.Time  `orm:"auto_now"`
	DeletedAt *time.Time `orm:"NULL"`

	// 业务字段
	Subject    string  `orm:""`
	Summary    string  `orm:"type(varchar);size(255);"`
	Content    string  `orm:"text"`
	ContentMd5 string  `orm:"type(char);size(32)"`
	Views      uint    `orm:"index"`
	BuyPrice   float64 `orm:""`
	ViewPrice  float32 `orm:"digits(8);decimals(2);"`
	Published  bool    `orm:""`

	// 关联字段
}

// TableIndex 索引配置
func (*Article) TableIndex() [][]string {
	return [][]string{
		[]string{"UpdatedAt"},
		[]string{"Subject", "Summary"},
	}
}
func (*Article) TableUnique() [][]string {
	return [][]string{}
}

// TableName 表名
func (*Article) TableName() string {
	return "article"
}

// TableEngine 存储引擎
func (*Article) TableEngine() string {
	return "INNODB"
}
