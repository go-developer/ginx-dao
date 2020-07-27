package dao

import (
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

// BaseDao 基础dao
//
// Author : go_developer@163.com<张德满>
type BaseDao struct{}

// ModifyStatus 根据主键修改状态
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/27 23:51:42
func (bd *BaseDao) ModifyStatus(dc *godb.DBClient, table string, primaryID uint64, fromStatus uint, targetStatus uint) (affectRows int64, err error) {
	db := dc.GormDB.Table(table).Where("id = ? AND status = ?", primaryID, fromStatus).Update(map[string]interface{}{"status": targetStatus})
	return db.RowsAffected, db.Error
}

// Create 创建新数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/27 23:11:17
func (bd *BaseDao) Create(dc *godb.DBClient, table string, data interface{}) error {
	return dc.GormDB.Table(define.ProjectTableName).Create(data).Error
}
