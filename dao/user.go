package dao

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
