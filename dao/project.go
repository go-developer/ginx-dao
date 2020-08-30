package dao

var (
	// Project 项目表操作实例
	//
	// Author : go_developer@163.com<张德满>
	Project *ProjectDao
)

// NewProjectDao 获取操作实例(单例) ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/27 23:06:05
func init()  {
	Project = &ProjectDao{}
}

// ProjectDao 项目表管理
//
// Author : go_developer@163.com<张德满>
type ProjectDao struct {
	BaseDao
}
