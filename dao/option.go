// Package dao...
//
// Author: go_developer@163.com<张德满>
//
// File:  option.go
//
// Description: option 处理各种功能性的option
//
// Date: 2020/10/11 10:23 下午
package dao

import "errors"

// SearchOption 检索条件
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:02 下午 2020/10/11
type SearchOption struct {
	Page       int                      // 页码
	Size       int                      // 每页数量
	Where      map[string]interface{}   // where 条件
	WhereIn    map[string][]interface{} // whereIn 条件
	WhereNotIn map[string][]interface{} // where not in 条件
	Expression []ExpressionOperate      // 表达式操作 如： create_time >= time.Now.Unix()
	QuerySql   []QuerySql               // 设置原始的sql
}

// QuerySql 手动设置sql
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:48 下午 2020/10/12
type QuerySql struct {
	Sql      string        // sql段
	BindData []interface{} // 绑定的数据
}

// ExpressionOperate ...
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:45 下午 2020/10/12
type ExpressionOperate struct {
	Field    string      // 字段
	Operate  string      // 操作符 > < >= <= != =
	BindData interface{} // 绑定的数据
}

// SearchOptionFunc 设置检索的参数
//
// Author : zhangdeman001@ke.com<张德满>
//
// Date : 10:15 下午 2020/10/11
type SearchOptionFunc func(so *SearchOption, data interface{}) error

// SetSearchOption 设置检索参数
//
// Author : go_developer@163.com<张德满>
//
// Date : 11:07 下午 2020/10/11
type SetSearchOption struct {
	Func SearchOptionFunc // 设置检索条件的函数
	Data interface{}      // 检索条件绑定的值
}

// SetSearchOptionPage 设置检索页码
//
// Author : zhangdeman001@ke.com<张德满>
//
// Date : 10:17 下午 2020/10/11
func SetSearchOptionPage(so *SearchOption, data interface{}) error {
	var (
		success bool
	)
	if so.Page, success = data.(int); !success {
		return errors.New("page is not int")
	}
	return nil
}

// SetSearchOptionSize 设置检索的size
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:22 下午 2020/10/11
func SetSearchOptionSize(so *SearchOption, data interface{}) error {
	var (
		success bool
	)
	if so.Size, success = data.(int); !success {
		return errors.New("size is not int")
	}
	return nil
}

// SetSearchOptionWhere 设置where条件
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:25 下午 2020/10/11
func SetSearchOptionWhere(so *SearchOption, data interface{}) error {
	var (
		success bool
		where   map[string]interface{}
	)
	if nil == so.Where {
		so.Where = make(map[string]interface{})
	}
	if where, success = data.(map[string]interface{}); !success {
		return errors.New("where is not map[string]interface{}")
	}

	if nil == so.Where {
		so.Where = make(map[string]interface{})
	}
	for k := range where {
		so.Where[k] = where[k]
	}
	return nil
}

// SetSearchOptionWhereIn 设置where in条件
//
// Author : go_developer@163.com<张德满>
//
// Date : 10:51 下午 2020/10/11
func SetSearchOptionWhereIn(so *SearchOption, data interface{}) error {
	var (
		success bool
		whereIn map[string][]interface{}
	)
	if whereIn, success = data.(map[string][]interface{}); !success {
		return errors.New("whereIn is not map[string][]interface{}")
	}
	if nil == so.WhereIn {
		so.WhereIn = make(map[string][]interface{})
	}
	for k := range whereIn {
		so.WhereIn[k] = whereIn[k]
	}
	return nil
}

// SetSearchOptionWhereNotIn 设置where not in条件
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:54 下午 2020/10/12
func SetSearchOptionWhereNotIn(so *SearchOption, data interface{}) error {
	var (
		success bool
	)
	if so.WhereNotIn, success = data.(map[string][]interface{}); !success {
		return errors.New("whereIn is not map[string][]interface{}")
	}
	return nil
}

// SetSearchOptionExpress 设置表达式
//
// Author : go_developer@163.com<张德满>
//
// Date : 2:55 下午 2020/10/12
func SetSearchOptionExpress(so *SearchOption, data interface{}) error {
	var (
		success    bool
		expression []ExpressionOperate
	)
	if expression, success = data.([]ExpressionOperate); !success {
		return errors.New("exception format is not []ExpressionOperate")
	}
	if nil == so.Expression {
		so.Expression = make([]ExpressionOperate, 0)
	}
	so.Expression = append(so.Expression, expression...)
	return nil
}

// SetSearchOptionQuerySql 设置原生的query sql
//
// Author : go_developer@163.com<张德满>
//
// Date : 3:13 下午 2020/10/12
func SetSearchOptionQuerySql(so *SearchOption, data interface{}) error {
	var (
		success  bool
		querySql []QuerySql
	)
	if querySql, success = data.([]QuerySql); !success {
		return errors.New("query sql format is not []QuerySql")
	}
	if nil == so.QuerySql {
		so.QuerySql = make([]QuerySql, 0)
	}
	so.QuerySql = append(so.QuerySql, querySql...)
	return nil
}
