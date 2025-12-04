package gorm

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Content struct {
	gorm.Model
	Subject string
	Likes   uint `gorm:""`
	Views   uint `gorm:""`
	//Likes       uint `gorm:"default:99"`
	//Views       *uint `gorm:"default:99"`
	PublishTime *time.Time

	// 禁用写操作
	Sv       string `gorm:"-:migration;<-:false"`
	AuthorID uint
}

type Author struct {
	gorm.Model
	Status int

	Name  string
	Email string
}

type ContentStrPK struct {
	ID          string `gorm:"primary_key"`
	Subject     string
	Likes       uint
	Views       uint
	PublishTime *time.Time
}

const (
	DefaultLikes = 99
	DefaultViews
)

func NewContent() Content {
	return Content{
		Likes: DefaultLikes,
		Views: DefaultViews,
	}
}

func (c *Content) BeforeCreate(db *gorm.DB) error {
	// 业务
	if c.PublishTime == nil {
		now := time.Now()
		c.PublishTime = &now
	}
	// 配置
	db.Statement.AddClause(clause.OnConflict{UpdateAll: true})

	return nil
}

//	func (c *Content) AfterCreate(db *gorm.DB) error {
//		return errors.New("custom error")
//	}

func (c *Content) AfterFind(db *gorm.DB) error {
	if c.AuthorID == 0 {
		c.AuthorID = 1
	}
	return nil
}
