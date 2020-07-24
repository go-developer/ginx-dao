// Package define ...
//
// Author: go_developer@163.com<张德满>
//
// Description: 用户相关表的定义
//
// File: user.go
//
// Version: 1.0.0
//
// Date: 2020/07/19 19:36:13
package define

// User 用户表数据结构
//
// Author : go_developer@163.com<张德满>
type User struct {
	ID         uint64 `json:"id"`          //主键id
	Mail       string `json:"mail"`        //邮件
	Phone      string `json:"phone"`       //手机号
	Status     uint   `json:"status"`      //状态
	Role       uint64 `json:"role"`        //角色ID
	Name       string `json:"name"`        //姓名
	Password   string `json:"password"`    //密码
	Salt       string `json:"salt"`        //私钥
	CreateTime string `json:"create_time"` //创建时间
	ModifyTime string `json:"modify_time"` //更新时间
}

// 状态枚举值
const (
	// UserStatusNew 新加
	UserStatusNew = 0
	// UserStatusNormal 正常
	UserStatusNormal = 1
	// UserStatusForbidden 禁用
	UserStatusForbidden = 2
	// UserStatusDelete 删除
	UserStatusDelete = 3
)

// 状态表
var userStatusTable = map[uint]string{
	UserStatusNew:       "新建",
	UserStatusNormal:    "正常",
	UserStatusForbidden: "禁用",
	UserStatusDelete:    "删除",
}

// GetUserStatusTable 获取完整的用户表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 01:03:51
func GetUserStatusTable() map[uint]string {
	return userStatusTable
}

// GetUserStatusDesc 获取状态的描述
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/25 01:06:31
func GetUserStatusDesc(status uint) string {
	var (
		statusDesc string
		exist      bool
	)
	if statusDesc, exist = GetUserStatusTable()[status]; !exist {
		return "未知"
	}
	return statusDesc
}
