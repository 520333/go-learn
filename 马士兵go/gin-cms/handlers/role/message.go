package role

import (
	"ginCms/models"
)

// GetRowReq 接口的请求消息类型
type GetRowReq struct {
	ID uint `form:"id" binding:"required,gt=0"`
}

// RestoreReq 还原请求消息
type RestoreReq struct {
	IDList []uint `form:"id" binding:"gt=0"`
}

// DeleteReq 删除请求消息
type DeleteReq struct {
	IDList []uint `form:"id" binding:"gt=0"`
	Force  bool   `form:"force" binding:""`
}

// AddReq 添加请求消息
type AddReq struct {
	models.Role
	// 需要额外校验的字段
	Title string `json:"title" binding:"required"`
	Key   string `json:"key" binding:"required"`
}

// ToRole AddReq to Role
func (req AddReq) ToRole() *models.Role {
	row := req.Role
	row.Title = req.Title
	row.Key = req.Key
	return &row
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
