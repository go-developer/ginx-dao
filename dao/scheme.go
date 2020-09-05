// Package dao...
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
	return sd.Create(dbClient, define.SchemeTableName, schemeData)
}

// GetAllScheme 获取全部的scheme
//
// Author : go_developer@163.com<张德满>
func (sd *schemeDao) GetAllScheme(dbClient *godb.DBClient) ([]*define.Scheme, error) {
	var (
		err    error
		result []*define.Scheme
	)

	err = dbClient.GormDB.Table(define.SchemeTableName).Find(&result).Error
	return result, err
}

// ChangeStatus 修改 scheme 状态
//
// Author : go_developer@163.com<张德满>
func (sd *schemeDao) ChangeStatus(dbClient *godb.DBClient, schemeID uint64, currentStatus uint, targetStatus uint) (int64, error) {
	return sd.ModifyStatus(dbClient, define.SchemeTableName, schemeID, currentStatus, targetStatus)
}

// Update 更新scheme信息
//
// Author : go_developer@163.com<张德满>
func (sd *schemeDao) Update(dbClient *godb.DBClient, schemeID uint64, schemeName string) (int64, error) {
	db := dbClient.GormDB.Table(define.SchemeTableName).Where("id = ?", schemeID).Limit(1).Update(map[string]interface{}{"scheme": schemeName})
	return db.RowsAffected, db.Error
}
