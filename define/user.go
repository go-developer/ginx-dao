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
	UpdateTime string `json:"update_time"` //更新时间
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
