package common

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
