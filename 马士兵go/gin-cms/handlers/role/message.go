package role

import (
	"ginCms/models"
	"reflect"
)

// GetRowReq 接口的请求消息类型
type GetRowReq struct {
	ID uint `form:"id" binding:"required,gt=0"`
}

// EditUriReq URI上的id参数
type EditUriReq struct {
	ID uint `uri:"id" binding:"required,gt=0"`
}

// EditBodyReq 更新主体参数
type EditBodyReq struct {
	Title   *string `json:"title" field:"title"`
	Key     *string `json:"key" field:"key"`
	Enabled *bool   `json:"enabled" field:"enabled"`
	Weight  *int    `json:"weight" field:"weight"`
	Comment *string `json:"comment" field:"comment"`
}

func (req EditBodyReq) ToFieldMap() models.FieldMap {
	// 1.初始化Map
	m := models.FieldMap{}

	// 2.利用反射reflect来遍历req结构的全部字段
	reqType := reflect.TypeOf(req)
	reqVale := reflect.ValueOf(req)
	// 通过字段数量进行遍历
	for i, nums := 0, reqType.NumField(); i < nums; i++ {
		// 获取field tag
		fieldTag := reqType.Field(i).Tag.Get("field")
		// 存在 field tag才自动处理
		if fieldTag == "" {
			continue
		}
		// 判定字段是否为nil,这个值的判定
		if !reqVale.Field(i).IsNil() {
			if fieldTag == "some_field" {
			} else {
				m[fieldTag] = reqVale.Field(i).Elem().Interface() //放入map
			}
		}
	}
	return m

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
