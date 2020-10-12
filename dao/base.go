package dao

import (
	"errors"
	"strings"

	godb "github.com/go-developer/gorm-mysql"
	"github.com/jinzhu/gorm"
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
func (bd *BaseDao) GetDataList(dbClient *godb.DBClient, table string, result interface{}, optionList ...SetSearchOption) error {
	if queryObject, err := bd.buildQueryObject(dbClient, table, optionList...); nil != err {
		return err
	} else {
		return queryObject.Find(result).Error
	}
}

// buildQueryObject 构建查询对象
//
// Author : go_developer@163.com<张德满>
//
// Date : 12:00 下午 2020/10/12
func (bd *BaseDao) buildQueryObject(dbClient *godb.DBClient, table string, optionList ...SetSearchOption) (*gorm.DB, error) {
	option := &SearchOption{
		Page:    0,
		Size:    0,
		Where:   nil,
		WhereIn: nil,
	}
	// 设置查询条件
	for _, o := range optionList {
		if err := o.Func(option, o.Data); nil != err {
			return nil, err
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
			return nil, errors.New(field + " set where in condition，but value is empty")
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

	return dbClient.GormDB.Table(table).Where(whereSql, bindDataList...).Limit(option.Size).Offset((option.Page - 1) * option.Size), nil
}

// getBatchCreateSql 获取批量写入数据的sql
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:35 下午 2020/10/12
func (bd *BaseDao) getBatchCreateSql(table string, valueList []map[string]interface{}) (string, []interface{}, error) {
	if len(valueList) == 0 {
		return "", nil, errors.New(table + " batch create data with nothing")
	}

	// 提取字段列表
	fieldList := make([]string, 0)
	for field := range valueList[0] {
		fieldList = append(fieldList, field)
	}
	sql := "INSERT INTO " + table + "(`" + strings.Join(fieldList, "`,`") + "`) VALUES "
	// 记录绑定数据
	bindDataList := make([]interface{}, 0)
	// 组装写入的数据
	writeDataList := make([]string, 0)
	for _, item := range valueList {
		// 按组装字段的顺序组装数据
		tmpValList := make([]string, 0)
		for i := 0; i < len(fieldList); i++ {
			tmpValList = append(tmpValList, "?")
			bindDataList = append(bindDataList, item[fieldList[i]])
		}
		writeDataList = append(writeDataList, "("+strings.Join(tmpValList, ",")+")")
	}
	sql = sql + strings.Join(writeDataList, ",") + ";"
	return sql, bindDataList, nil
}

// BatchCreate 批量写入数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 4:50 下午 2020/10/12
func (bd *BaseDao) BatchCreate(dbClient *godb.DBClient, table string, valueList []map[string]interface{}) error {
	var (
		batchCreateSql string
		err            error
		bindDataList   []interface{}
	)
	if batchCreateSql, bindDataList, err = bd.getBatchCreateSql(table, valueList); nil != err {
		return err
	}
	if _, err = dbClient.GormDB.DB().Exec(batchCreateSql, bindDataList...); nil != err {
		return err
	}
	return nil
}
