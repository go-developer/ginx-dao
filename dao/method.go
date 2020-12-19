// Package dao...
//
// Description : method 请求方法表操作
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020-08-30 11:12 下午
package dao

import (
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
)

var (
	// Method method 表操作实例
	//
	// Author : go_developer@163.com<张德满>
	Method *methodDao
)

func init() {
	Method = &methodDao{}
}

type methodDao struct {
	BaseDao
}

// GetMethodList : GetMethodList 获取请求方法的列表
//
// Author : go_developer@163.com<张德满>
func (md *methodDao) GetMethodList(dbClient *godb.DBClient, optionList ...SetSearchOption) ([]define.Method, error) {
	var (
		methodList []define.Method
		err        error
	)
	if err = md.GetDataList(dbClient, define.DBTableMethod, &methodList, optionList...); nil != err {
		return make([]define.Method, 0), err
	}
	return methodList, nil
}

// GetMethodCount  获取请求方法的数量
//
// Author : go_developer@163.com<张德满>
func (md *methodDao) GetMethodCount(dbClient *godb.DBClient, optionList ...SetSearchOption) (int64, error) {
	return md.GetTotalDataCount(dbClient, define.DBTableMethod, optionList...)
}
