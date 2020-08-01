package dao

import (
	"github.com/go-developer/ginx-dao/define"
	godb "github.com/go-developer/gorm-mysql"
	"github.com/jinzhu/gorm"
)

var (
	// User 用户操作实例
	//
	// Author : go_developer@163.com<张德满>
	User *UserDao
)

// NewUserDao 获取用户操作实例
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/07/28 22:58:26
func NewUserDao() *UserDao {
	if nil == User {
		User = &UserDao{}
	}
	return User
}

// UserDao 用户操作模型
//
// Author : go_developer@163.com<张德满>
type UserDao struct {
	BaseDao
}

// CreateUser 创建用户
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 20:06:00
func (ud *UserDao) CreateUser(dc *godb.DBClient, info *define.User) error {
	return ud.Create(dc, define.UserTableName, info)
}

// GetDetailByPhone 根据手机号获取用户详情
//
// Author : go_developer@163.com<张德满>
//
// Date : 2020/08/01 20:17:26
func (ud *UserDao) GetDetailByPhone(dc *godb.DBClient, phone string) (*define.User, bool, error) {
	var (
		detail define.User
		ormDB  *gorm.DB
	)
	ormDB = dc.GormDB.Table(define.UserTableName).Where("phone = ?", phone).First(&detail)
	return &detail, ormDB.RecordNotFound(), ormDB.Error
}
