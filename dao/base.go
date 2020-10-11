package dao

import (
	"errors"

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
	return dc.GormDB.Table(table).Create(data).Error
}

// GetDataList 查询数据列表
//
// Author : zhangdeman001@ke.com<张德满>
//
// Date : 5:17 下午 2020/10/10
// ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 9:57 下午 2020/10/11
func (pad *projectAPIDao) GetDataList(dbClient *godb.DBClient, table string, result interface{}, optionList ...SetSearchOption) error {
	option := &SearchOption{
		Page:    0,
		Size:    0,
		Where:   nil,
		WhereIn: nil,
	}
	// 设置查询条件
	for _, o := range optionList {
		if err := o.Func(option, o.Data); nil != err {
			return err
		}
	}
	whereSql := ""
	bindDataList := make([]interface{}, 0)
	for field, value := range option.Where {
		bindDataList = append(bindDataList, value)
		if len(whereSql) > 0 {
			whereSql = whereSql + " AND " + field + " = ? "
		} else {
			whereSql = field + " = ?"
		}
	}

	for field, valueList := range option.WhereIn {
		if len(valueList) == 0 {
			return errors.New(field + " set where in condition，but value is empty")
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

	if option.Page <= 0 {
		option.Page = 1
	}

	if option.Size <= 0 {
		option.Size = 20
	}

	if option.Size > 100 {
		option.Size = 100
	}

	return dbClient.GormDB.Table(table).Where(whereSql, bindDataList...).Limit(option.Size).Offset((option.Page - 1) * option.Size).Find(result).Error
}
