// Package define ...
//
// Author: go_developer@163.com<张德满>
//
// Description: ginx_scheme 表相关定义
//
// File: scheme.go
//
// Version: 1.0.0
//
// Date: 2020/07/25 12:35:29
package define

// Scheme 请求scheme定义
//
// Author : go_developer@163.com<张德满>
type Scheme struct {
	ID           uint64 `json:"id"`             //主键ID
	Scheme       string `json:"scheme"`         //协议
	Status       uint   `json:"status"`         //状态
	CreateUserID uint64 `json:"create_user_id"` //创建人ID
	ModifyUserID uint64 `json:"modify_user_id"` //修改人ID
	CreateTime   string `json:"create_time"`    //创建时间
	ModifyTime   string `json:"modify_time"`    //修改时间
}

//协议状态枚举
const (
	// SchemeStatusUnused 待启用
	//
	// Author : go_developer@163.com<张德满>
	SchemeStatusUnused = 0

	// SchemeStatusUsing 使用中
	//
	// Author : go_developer@163.com<张德满>
	SchemeStatusUsing = 1

	// SchemeStatusForbbiden 禁止使用
	//
	// Author : go_developer@163.com<张德满>
	SchemeStatusForbbiden = 2
)

var schemeStatusTable = map[uint]string{
	SchemeStatusUnused:    "待启用",
	SchemeStatusUsing:     "使用中",
	SchemeStatusForbbiden: "已禁用",
}

// GetSchemeStatusTable 获取状态表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 12:48:31
func GetSchemeStatusTable() map[uint]string {
	return schemeStatusTable
}

// GetSchemeStatusDesc 获取状态描述
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 12:50:52
func GetSchemeStatusDesc(status uint) string {
	var (
		statusDesc string
		exist      bool
	)
	if statusDesc, exist = GetSchemeStatusTable()[status]; !exist {
		return "未知"
	}
	return statusDesc
}
