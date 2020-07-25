// Package define ...
//
// Author: go_developer@163.com<张德满>
//
// Description: ginx_project_api 表相关定义
//
// File: project_api.go
//
// Version: 1.0.0
//
// Date: 2020/07/25 14:36:48
package define

// ProjectAPI ginx_project_api 表数据结构
//
// Author : zhangdeman001@ke.com<张德满>
type ProjectAPI struct {
	ID           uint64 `json:"id"`             //主键ID
	ProjectID    uint64 `json:"project_id"`     //项目ID
	Status       uint   `json:"status"`         //api状态
	Name         string `json:"name"`           //API名称
	URI          string `json:"uri"`            //请求uri
	MapURI       string `json:"map_uri"`        //后端真实服务URI
	SchemeID     uint64 `json:"scheme_id"`      //请求协议ID
	MethodID     uint64 `json:"method_id"`      //请求方法ID
	Description  string `json:"description"`    //项目描述
	CreateUserID uint64 `json:"create_user_id"` //创建人ID
	ModifyUserID uint64 `json:"modify_user_id"` //修改人ID
	CreateTime   string `json:"create_time"`    //创建时间
	ModifyTime   string `json:"modify_time"`    //修改时间
}

const (
	// ProjectAPIStatusUnused 待启用
	//
	// Author : go_developer@163.com<张德满>
	ProjectAPIStatusUnused = 0

	// ProjectAPIStatusUsing 使用中
	//
	// Author : go_developer@163.com<张德满>
	ProjectAPIStatusUsing = 1

	// ProjectAPIStatusForbidden 禁用
	//
	// Author : go_developer@163.com<张德满>
	ProjectAPIStatusForbidden = 2
)

var projectAPIStatusTable = map[uint]string{
	ProjectAPIStatusUnused:    "未启用",
	ProjectAPIStatusUsing:     "使用中",
	ProjectAPIStatusForbidden: "停用",
}

// GetProjectAPIStatusTable 获取api状态表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 18:24:35
func GetProjectAPIStatusTable() map[uint]string {
	return projectAPIStatusTable
}

// GetProjectAPIStatusDesc 状态描述
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 18:26:42
func GetProjectAPIStatusDesc(status uint) string {
	var (
		statusDesc string
		exist      bool
	)
	if statusDesc, exist = GetProjectAPIStatusTable()[status]; !exist {
		return "未知"
	}
	return statusDesc
}
