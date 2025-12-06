package gorm

import (
	"log"
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

//type Author struct {
//	gorm.Model
//	Status int
//
//	Name  string
//	Email string
//}

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
	log.Println("Content BeforeCreate Hook")
	return nil
}

func (c *Content) AfterCreate(db *gorm.DB) error {
	//return errors.New("custom error")
	return nil
}

func (c *Content) AfterFind(db *gorm.DB) error {
	if c.AuthorID == 0 {
		c.AuthorID = 1
	}
	return nil
}

// Author 作者
type Author struct {
	gorm.Model
	Status int
	Name   string
	Email  string
	Essay  []Essay // 拥有多个论文内容
	//EssayMate EssayMate // 拥有一个论文元信息

}

// Essay 论文内容
// Create Table: CREATE TABLE `go_essay` (
// `id` bigint unsigned NOT NULL AUTO_INCREMENT,
// `created_at` datetime(3) DEFAULT NULL,
// `updated_at` datetime(3) DEFAULT NULL,
// `deleted_at` datetime(3) DEFAULT NULL,
// `subject` longtext,
// `content` longtext,
// `author_id` bigint unsigned DEFAULT NULL,
// PRIMARY KEY (`id`),
// KEY `idx_go_essay_deleted_at` (`deleted_at`),
// KEY `fk_go_author_essay` (`author_id`),
// CONSTRAINT `fk_go_author_essay` FOREIGN KEY (`author_id`) REFERENCES `go_author` (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
type Essay struct {
	gorm.Model
	Subject   string
	Content   string
	AuthorID  *uint  // 外键字段
	Author    Author //属于某个作者
	EssayMate EssayMate
	Tag       []Tag `gorm:"many2many:essay_tag;"`
}

// EssayMate 论文元信息
type EssayMate struct {
	gorm.Model
	Keyword     string
	Description string
	EssayID     *uint // 外键字段
	//Essay       *Essay // 属于一个论文内容
}

type Tag struct {
	gorm.Model
	Title string
	Essay []Essay `gorm:"many2many:essay_tag;"` // 拥有多个Essay
}
