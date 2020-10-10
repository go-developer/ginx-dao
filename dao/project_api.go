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
	"errors"

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
func (pad *projectAPIDao) GetProjectAPIList(dbClient *godb.DBClient, page int64, limit int64, where map[string]interface{}, whereIn map[string][]interface{}) ([]*define.ProjectAPI, error) {
	whereSql := ""
	bindDataList := make([]interface{}, 0)
	for field, value := range where {
		bindDataList = append(bindDataList, value)
		if len(whereSql) > 0 {
			whereSql = whereSql + " AND " + field + " = ? "
		} else {
			whereSql = field + " = ?"
		}
	}

	for field, valueList := range whereIn {
		if len(valueList) == 0 {
			return make([]*define.ProjectAPI, 0), errors.New(field + "设置in条件，但是无任何取值")
		}
		bindDataList = append(bindDataList, valueList...)
		if len(whereSql) > 0 {
			whereSql = " AND " + field + " IN ( "

		} else {
			whereSql = field + " IN ( "
		}
		for i := 0; i < len(valueList)-1; i++ {
			whereSql = whereSql + " ?, "
		}
		whereSql = whereSql + " )"
	}

	var (
		apiList []*define.ProjectAPI
		err     error
	)
	if err = dbClient.GormDB.Table(define.DBTableProjectAPI).Where(whereSql, bindDataList...).Limit(limit).Offset((page - 1) * limit).Find(&apiList).Error; nil != err {
		return make([]*define.ProjectAPI, 0), err
	}
	return apiList, nil
}
