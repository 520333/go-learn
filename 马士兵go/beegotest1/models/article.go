package models

import "time"

type Article struct {
	// 基础字段
	Id        uint64     `orm:"description(自增主键)" json:"id"`
	CreatedAt time.Time  `orm:"auto_now_add" json:"createdAt"`
	UpdatedAt time.Time  `orm:"auto_now" json:"updatedAt"`
	DeletedAt *time.Time `orm:"NULL" json:"-"`

	// 业务字段
	Subject    string  `orm:"" json:"subject"`
	Summary    string  `orm:"type(varchar);size(255);" json:"summary"`
	Content    string  `orm:"text" json:"content"`
	ContentMd5 string  `orm:"type(char);size(32)" json:"content_md5"`
	Views      uint    `orm:"index" json:"views"`
	BuyPrice   float64 `orm:"" json:"buyPrice"`
	ViewPrice  float32 `orm:"digits(8);decimals(2);" json:"viewPrice"`
	Published  bool    `orm:"" json:"published"`

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
