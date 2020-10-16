package dao

import (
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

var (
	// Project 项目表操作实例
	//
	// Author : go_developer@163.com<张德满>
	Project *projectDao
)

// NewProjectDao 获取操作实例(单例) ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/27 23:06:05
func init() {
	Project = &projectDao{}
}

// ProjectDao 项目表管理
//
// Author : go_developer@163.com<张德满>
type projectDao struct {
	BaseDao
}

// GetProjectList 获取项目列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:26 下午 2020/10/16
func (pd *projectDao) GetProjectList(dbClient *godb.DBClient, optionList ...SetSearchOption) ([]*define.Project, error) {
	var (
		err  error
		list []*define.Project
	)
	if err = pd.GetDataList(dbClient, define.DBTableProject, &list, optionList...); nil != err {
		return make([]*define.Project, 0), err
	}
	return list, nil
}
