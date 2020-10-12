// Package dao...
//
// Author: go_developer@163.com<张德满>
//
// File:  base_test.go
//
// Description: base_test 单元测试
//
// Date: 2020/10/12 6:34 下午
package dao

import (
	"fmt"
	"testing"
)

// TestGetBatchCreateSql 测试生成批量插入数据的sql
//
// Author : go_developer@163.com<张德满>
//
// Date : 6:40 下午 2020/10/12
func TestGetBatchCreateSql(t *testing.T) {
	table := "test"
	valueList := []map[string]interface{}{
		{"name": "zhangdeman", "age": 18},
		{"name": "zhangdeman1", "age": 181},
		{"name": "zhangdeman2", "age": 182},
		{"name": "zhangdeman3", "age": 183},
	}
	bd := new(BaseDao)
	sql, data, err := bd.getBatchCreateSql(table, valueList)
	fmt.Println(sql, data, err)
}
