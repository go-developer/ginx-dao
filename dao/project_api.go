// Package dao...
//
// Author: go_developer@163.com<张德满>
//
// File:  project_api.go
//
// Description: project_api 表操作
//
// Date: 2020/10/10 3:30 下午
package dao

import (
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

var (
	// ProjectAPI 操作实例
	//
	// Author : go_developer@163.com<张德满>
	//
	// Date : 3:31 下午 2020/10/10
	ProjectAPI *projectAPIDao
)

func init() {
	ProjectAPI = &projectAPIDao{}
}

type projectAPIDao struct {
	BaseDao
}

// CreateProjectAPI 创建一个API
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:58 下午 2020/10/10
func (pad *projectAPIDao) CreateProjectAPI(dbClient *godb.DBClient, apiData *define.ProjectAPI) error {
	return pad.Create(dbClient, define.DBTableProjectAPI, apiData)
}

// GetProjectAPIByID 根据 API ID 获取API 信息
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:02 下午 2020/10/10
func (pad *projectAPIDao) GetProjectAPIByID(dbClient *godb.DBClient, apiID int64) (*define.ProjectAPI, error) {
	var (
		apiData define.ProjectAPI
		err     error
	)
	err = dbClient.GormDB.Table(define.DBTableProjectAPI).Where("id = ?", apiID).First(&apiData).Error
	return &apiData, err
}

// GetProjectAPIByProjectID 根据projectID获取api列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 5:08 下午 2020/10/10
func (pad *projectAPIDao) GetProjectAPIByProjectID(dbClient *godb.DBClient, projectID int64) ([]*define.ProjectAPI, error) {
	var (
		err     error
		apiList []*define.ProjectAPI
	)
	if err = dbClient.GormDB.Table(define.DBTableProjectAPI).Where("project_id = ?", projectID).Find(&apiList).Error; nil != err {
		return make([]*define.ProjectAPI, 0), err
	}
	return apiList, nil
}

// GetProjectAPIList 查询api列表
//
// Author : zhangdeman001@ke.com<张德满>
//
// Date : 5:17 下午 2020/10/10
func (pad *projectAPIDao) GetProjectAPIList(dbClient *godb.DBClient, optionList ...SetSearchOption) ([]*define.ProjectAPI, error) {

	var (
		apiList []*define.ProjectAPI
		err     error
	)
	if err = pad.GetDataList(dbClient, define.DBTableProjectAPI, &apiList, optionList...); nil != err {
		return make([]*define.ProjectAPI, 0), err
	}
	return apiList, nil
}
