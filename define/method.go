// Package define ...
//
// Author: go_developer@163.com<张德满>
//
// Description: ginx_method 表相关定义
//
// File: method.go
//
// Version: 1.0.0
//
// Date: 2020/07/25 12:52:15
package define

// Method 表结构
//
// Author : go_developer@163.com<张德满>
type Method struct {
	ID           uint64 `json:"id"`             //主键ID
	Method       string `json:"method"`         //请求方法
	Status       uint   `json:"status"`         //方法状态
	CreateUserID uint64 `json:"create_user_id"` //创建人ID
	ModifyUserID uint64 `json:"modify_user_id"` //修改人ID
	CreateTime   string `json:"create_time"`    //创建时间
	ModifyTime   string `json:"modify_time"`    //修改时间
}

const (
	// MethodStatusUnused 待启用
	//
	// Author : go_developer@163.com<张德满>
	MethodStatusUnused = 0

	// MethodStatusUsing 使用中
	//
	// Author : go_developer@163.com<张德满>
	MethodStatusUsing = 1

	// MethodStatusForbbiden 禁用
	//
	// Author : go_developer@163.com<张德满>
	MethodStatusForbbiden = 2
)

var methodStatusTable = map[uint]string{
	MethodStatusUnused:    "待启用",
	MethodStatusUsing:     "使用中",
	MethodStatusForbbiden: "已禁用",
}

// GetMethodStatusTable 获取请求方法状态表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 13:00:27
func GetMethodStatusTable() map[uint]string {
	return methodStatusTable
}

// GetMethodStatusDesc 获取方法状态描述
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 13:06:37
func GetMethodStatusDesc(status uint) string {
	var (
		statusDesc string
		exist      bool
	)
	if statusDesc, exist = GetMethodStatusTable()[status]; !exist {
		return "未知"
	}
	return statusDesc
}
