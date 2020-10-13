// Package dao...
//
// Author: go_developer@163.com<张德满>
//
// File:  project_api_param.go
//
// Description: project_api_param api参数配置
//
// Date: 2020/10/12 4:07 下午
package dao

import (
	"encoding/json"

	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

var (
	// ProjectAPIParam 参数配置表操作实例
	ProjectAPIParam *projectAPIParamDao
)

func init() {
	ProjectAPIParam = &projectAPIParamDao{}
}

type projectAPIParamDao struct {
	BaseDao
}

// SetAPIParam 设置API的参数
//
// Author : zhangdeman001@ke.com<张德满>
//
// Date : 4:17 下午 2020/10/12
func (pap *projectAPIParamDao) BatchSetAPIParam(dbClient *godb.DBClient, paramList []*define.ProjectAPIParam) error {
	var (
		dataTable []map[string]interface{}
		err       error
		byteData  []byte
	)
	if byteData, err = json.Marshal(paramList); nil != err {
		return err
	}
	if err = json.Unmarshal(byteData, &dataTable); nil != err {
		return err
	}
	return pap.BatchCreate(dbClient, define.DBTableProjectAPIParam, dataTable)
}

// projectAPIParamDao 通过 API ID 获取参数列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:42 上午 2020/10/13
func (pap *projectAPIParamDao) GetParamListByAPIID(client godb.DBClient, apiID int64) ([]*define.ProjectAPIParam, error) {
	var (
		err       error
		paramList []*define.ProjectAPIParam
	)
	if err = client.GormDB.Table(define.DBTableProjectAPIParam).Where("api_id = ?", apiID).Find(&paramList).Error; nil != err {
		return make([]*define.ProjectAPIParam, 0), err
	}
	return paramList, nil
}
