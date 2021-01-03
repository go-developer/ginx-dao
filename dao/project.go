package dao

import (
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
	"github.com/pkg/errors"
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
		return make([]*define.Project, 0), errors.Wrap(err, "获取项目列表失败")
	}
	return list, nil
}

// GetProjectDetailByFlag 根据flag获取项目详情
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:32 下午 2021/1/3
func (pd *projectDao) GetProjectDetailByFlag(dbClient *godb.DBClient, flag string) (*define.Project, error) {
	var (
		err    error
		result define.Project
	)
	if err = dbClient.GormDB.Table(define.DBTableProject).Where("flag = ?", flag).First(&result).Error; nil != err {
		return nil, errors.Wrap(err, "根据flag查询项目详情失败")
	}
	if result.ID == 0 {
		return nil, nil
	}
	return &result, nil
}
