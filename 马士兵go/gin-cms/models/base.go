package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

const (
	PageNumDefault  = 1
	PageSizeDefault = 10
	PageSizeMax     = 100

	SortFieldDefault  = "id"
	SortMethodDefault = "desc"
)

// Clean 整理Sorter
func (s *Sorter) Clean() {
	if s.SortField == nil {
		temp := SortFieldDefault
		s.SortField = &temp

	}
	if s.SortMethod == nil {
		temp := SortMethodDefault
		s.SortMethod = &temp
	}
}

// Clean 整理Pager
func (p *Pager) Clean() {
	if p.PageNum == nil {
		temp := PageNumDefault
		p.PageNum = &temp
	}
	if p.PageSize == nil {
		temp := PageSizeDefault
		p.PageSize = &temp
	}
	if *p.PageSize > PageSizeMax {
		temp := PageSizeMax
		p.PageSize = &temp
	}
}

// Sorter 通用的查询列表排序类型
type Sorter struct {
	SortField  *string `form:"sortField" binding:"omitempty,gt=0"`            // 排序字段
	SortMethod *string `form:"sortMethod" binding:"omitempty,oneof=asc desc"` // 排序方式 asc,desc
}

// Pager 通用的查询列表翻页类型
type Pager struct {
	PageNum  *int `form:"pageNum" binding:"omitempty,gt=1"`
	PageSize *int `form:"pageSize" binding:"omitempty,gt=0"`
}
