// Package dao ...
//
// Description : scheme 项目协议表(http/https/tcp等)
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020-08-30 11:06 下午
package dao

import (
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

var (
	// Scheme 表实例
	//
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

// Create 创建一个新的 scheme
//
// Author : go_developer@163.com<张德满>
func (sd *schemeDao) CreateScheme(dbClient *godb.DBClient, schemeData *define.Scheme) error {
	return sd.Create(dbClient, define.DBTableScheme, schemeData)
}

// GetAllScheme 获取全部的scheme
//
// Author : go_developer@163.com<张德满>
func (sd *schemeDao) GetAllScheme(dbClient *godb.DBClient) ([]*define.Scheme, error) {
	var (
		err    error
		result []*define.Scheme
	)

	err = dbClient.GormDB.Table(define.DBTableScheme).Find(&result).Error
	return result, err
}

// ChangeStatus 修改 scheme 状态
//
// Author : go_developer@163.com<张德满>
func (sd *schemeDao) ChangeStatus(dbClient *godb.DBClient, schemeID uint64, currentStatus uint, targetStatus uint) (int64, error) {
	return sd.ModifyStatus(dbClient, define.DBTableScheme, schemeID, currentStatus, targetStatus)
}

// Update 更新scheme信息
//
// Author : go_developer@163.com<张德满>
func (sd *schemeDao) Update(dbClient *godb.DBClient, schemeID uint64, schemeName string) (int64, error) {
	db := dbClient.GormDB.Table(define.DBTableScheme).Where("id = ?", schemeID).Limit(1).Update(map[string]interface{}{"scheme": schemeName})
	return db.RowsAffected, db.Error
}

// GetSchemeList 获取列表
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/11/08 01:38:46
func (sd *schemeDao) GetSchemeList(dbClient *godb.DBClient, optionList ...SetSearchOption) ([]define.Scheme, error) {
	var (
		schemeList []define.Scheme
		err        error
	)
	if err = sd.GetDataList(dbClient, define.DBTableScheme, &schemeList, optionList...); nil != err {
		return make([]define.Scheme, 0), err
	}
	return schemeList, nil
}
