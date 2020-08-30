// Package dao...
//
// Description : method 请求方法表操作
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020-08-30 11:12 下午
package dao

var (
	// Method method 表操作实例
	//
	// Author : go_developer@163.com<张德满>
	Method *methodDao
)

func init()  {
	Method = &methodDao{}
}

type methodDao struct {
	BaseDao
}
