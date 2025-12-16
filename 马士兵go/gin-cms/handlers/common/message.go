package common

const (
	PageNumDefault  = 1
	PageSizeDefault = 10
	PageSizeMax     = 100

	SortFieldDefault  = "id"
	SortMethodDefault = "desc"
)

// Clean 整理Filter
func (f *Filter) Clean() {
	if f.Keyword == nil {
		temp := ""
		f.Keyword = &temp
	}
}

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

// Filter 通用的请求过滤类型
type Filter struct {
	// 指针类型表示该字段可以不填
	Keyword *string `form:"keyword" binding:"omitempty,gt=0"` //omitempty 非零值才校验
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

// 整体清理
