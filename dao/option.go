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
	Page    int                      // 页码
	Size    int                      // 每页数量
	Where   map[string]interface{}   // where 条件
	WhereIn map[string][]interface{} // whereIn 条件
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
	)
	if so.Where, success = data.(map[string]interface{}); !success {
		return errors.New("where is not map[string]interface{}")
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
	)
	if so.WhereIn, success = data.(map[string][]interface{}); !success {
		return errors.New("whereIn is not map[string][]interface{}")
	}
	return nil
}
