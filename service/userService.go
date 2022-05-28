package service

import (
	"log"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
)

type UserService struct {
	userDao repository.UserDao
}

func InitUserService() UserService {
	//初始化userDao对象
	userDao := repository.InitUserDao()
	log.Println("[InitUserService func] success !")

	return UserService{userDao: userDao}
}

//用户注册，参数name为用户名，password为密码
func (u *UserService) Register(name string, password string) (*model.User, int32) {
	//验证数据长度合法性
	//客户端已帮忙进行验证长度必须大于5
	if len(name) > 32 || len(password) > 32 {
		return nil, 402
	}
	//查询是否已存在该用户
	temp := u.userDao.QueryByName(name)
	if temp != nil {
		return nil, 403 //用户名已存在
	}

	newUser := model.User{
		Name:          name,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	res, err := u.userDao.Create(&newUser)
	if err != nil || res == nil {
		return nil, 500
	}
	return res, 0
}

func (u *UserService) Login(name string, password string) (*model.User, error) {
	return nil, nil
}
