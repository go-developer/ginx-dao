// Package define ...
//
// Author: go_developer@163.com<张德满>
//
// Description: ginx_project 表相关定义
//
// File: project.go
//
// Version: 1.0.0
//
// Date: 2020/07/25 13:11:08
package define

// Project ginx_project表数据结构
//
// Author : go_developer@163.com<张德满>
type Project struct {
	ID           uint64 `json:"id"`             //主键ID
	Flag         string `json:"flag"`           //项目标识
	Name         string `json:"name"`           //项目名称
	Description  string `json:"description"`    //项目描述
	Status       uint   `json:"status"`         //项目状态
	Domain       string `json:"domain"`         //域名,不包含协议头
	Port         uint   `json:"port"`           //请求端口
	CreateUserID uint64 `json:"create_user_id"` //创建人ID
	ModifyUserID uint64 `json:"modify_user_id"` //修改人ID
	CreateTime   string `json:"create_time"`    //创建时间
	ModifyTime   string `json:"modify_time"`    //修改时间
}

const (
	// ProjectStatusUnused 待启用
	//
	// Author : go_developer@163.com<张德满>
	ProjectStatusUnused = 0

	// ProjectStatusUsing 使用中
	//
	// Author : go_developer@163.com<张德满>
	ProjectStatusUsing = 1

	// ProjectStatusForbidden 禁用
	//
	// Author : go_developer@163.com<张德满>
	ProjectStatusForbidden = 2
)

var projectStatusTable = map[uint]string{
	ProjectStatusUnused:    "未启用",
	ProjectStatusUsing:     "使用中",
	ProjectStatusForbidden: "已禁用",
}

// GetProjectStatusTable 获取项目状态表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 14:16:41
func GetProjectStatusTable() map[uint]string {
	return projectStatusTable
}

// GetProjectStatusDesc ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 14:21:28
func GetProjectStatusDesc(status uint) string {
	var (
		statusDesc string
		exist      bool
	)
	if statusDesc, exist = GetProjectStatusTable()[status]; !exist {
		return "未知"
	}
	return statusDesc
}
