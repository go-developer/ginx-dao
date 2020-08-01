// Package dao 项目dao层
//
// Author: go_developer@163.com<张德满>
//
// Description:
//
// File: base.go
//
// Version: 1.0.0
//
// Date: 2020/07/19 19:26:34
package dao

import (
	godb "github.com/go-developer/gorm-mysql"
)

// InitDatabase 初始化数据库连接
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 19:29:27
func InitDatabase(dbConfig *godb.DBConfig) {
	godb.InitDatabaseClient(dbConfig)
}

// commonUpdateByID 根据主键ID更新数据
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 20:05:49
func commonUpdateByID(dc *godb.DBClient, table string, ID uint64, updateData interface{}) (affectRows int, err error) {
	return 0, nil
}

// commonCreateData 公共的写数据方法
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/19 20:07:58
func commonCreateData(dc *godb.DBClient, table string, saveData interface{}) (uint64, error) {
	return 0, nil
}
