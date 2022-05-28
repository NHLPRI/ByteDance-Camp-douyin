//repository层就是dao层，负责对数据库直接的CURD
package repository

import (
	"log"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/jinzhu/gorm"
)

//user的持久层对象
type UserDao struct {
	db *gorm.DB //数据库对象属性
}

//初始化userDao对象
func InitUserDao() UserDao {
	//自动创建user表
	db := common.GetDB()
	db.AutoMigrate(&model.User{})

	log.Println("[InitUserDao func] success !")
	return UserDao{db: db}
}

//添加表记录方法
func (u *UserDao) Create(user *model.User) (*model.User, error) {
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

//用户ID查询记录
func (u *UserDao) QueryById(id int64) (*model.User, error) {
	var user model.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//用户name查询记录
func (u *UserDao) QueryByName(name string) *model.User {
	var user model.User
	u.db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return &user
	}
	return nil
}
