// Package define ...
//
// Author: go_developer@163.com<张德满>
//
// Description: ginx_project_api_param 表相关定义
//
// File: project_api_param.go
//
// Version: 1.0.0
//
// Date: 2020/07/25 18:30:07
package define

// ProjectAPIParam 数据结构定义
//
// Author : go_developer@163.com<张德满>
type ProjectAPIParam struct {
	ID           uint64      `json:"id"`             // 主键ID
	ProjectID    uint64      `json:"project_id"`     // 项目ID
	APIID        uint64      `json:"api_id"`         // 接口ID
	ParamName    string      `json:"param_name"`     // 参数名称
	ParamType    uint        `json:"param_type"`     // 参数类型
	IsRequired   uint        `json:"is_required"`    // 是否必填
	DefaultValue interface{} `json:"default_value"`  // 非必传参数,没传时候的默认值
	ExampleValue string      `json:"example_value"`  // 请求参数示例值
	Status       uint        `json:"status"`         // 方法状态
	CreateUserID uint64      `json:"create_user_id"` // 创建人ID
	ModifyUserID uint64      `json:"modify_user_id"` // 修改人ID
	CreateTime   string      `json:"create_time"`    // 创建时间
	ModifyTime   string      `json:"modify_time"`    // 修改时间
}

const (
	// APIParamStatusUnused 未使用
	//
	// Author : go_developer@163.com<张德满>
	APIParamStatusUnused = 0
	// APIParamStatusUsing 生效中
	//
	// Author : go_developer@163.com<张德满>
	APIParamStatusUsing = 1
	// APIParamStatusForbidden 禁用
	//
	// Author : go_developer@163.com<张德满>
	APIParamStatusForbidden = 2
)

var apiParamStatusTable = map[uint]string{
	APIParamStatusUnused:    "待启用",
	APIParamStatusUsing:     "使用中",
	APIParamStatusForbidden: "已停用",
}

// GetAPIParamStatusTable 获取api参数状态表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 18:57:36
func GetAPIParamStatusTable() map[uint]string {
	return apiParamStatusTable
}

// GetAPIParamStatusDesc 获取参数状态描述
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 18:59:56
func GetAPIParamStatusDesc(status uint) string {
	var (
		statusDesc string
		exist      bool
	)
	if statusDesc, exist = GetAPIParamStatusTable()[status]; !exist {
		return "未知"
	}
	return statusDesc
}
