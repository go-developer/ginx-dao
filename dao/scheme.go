// Package dao...
//
// Description : scheme 项目协议表(http/https/tcp等)
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020-08-30 11:06 下午
package dao

var (
	// Desc : scheme 表操作实例
	//
	// Author : go_developer@163.com<张德满>
	Scheme *schemeDao
)

// Desc : 实例化
//
// Author : go_developer@163.com<张德满>
func init() {
	Scheme = &schemeDao{}
}

type schemeDao struct {
	BaseDao
}