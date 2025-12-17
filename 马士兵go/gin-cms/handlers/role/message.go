package role

import (
	"ginCms/models"
)

// GetRowReq 接口的请求消息类型
type GetRowReq struct {
	ID uint `form:"id" binding:"required,gt=0"`
}

// AddReq 添加请求消息
type AddReq struct {
	models.Role
	// 需要额外校验的字段
	Title string `json:"title" binding:"required"`
	Key   string `json:"key" binding:"required"`
}

// GetListReq role列表请求参数类型
type GetListReq struct {
	models.RoleFilter // 过滤
	models.Sorter     // 排序
	models.Pager      // 翻页
}

// Clean 查询列表参数清理
func (req *GetListReq) Clean() {
	req.RoleFilter.Clean()
	req.Sorter.Clean()
	req.Pager.Clean()
}
