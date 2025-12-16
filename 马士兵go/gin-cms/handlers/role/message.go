package role

import "ginCms/handlers/common"

// GetRowReq 接口的请求消息类型
type GetRowReq struct {
	ID uint `form:"id" binding:"required,gt=0"`
}

// GetListReq role列表请求参数类型
type GetListReq struct {
	common.Filter // 过滤
	common.Sorter // 排序
	common.Pager  // 翻页
}

// Clean 查询列表参数清理
func (req *GetListReq) Clean() {
	req.Filter.Clean()
	req.Sorter.Clean()
	req.Pager.Clean()
}
