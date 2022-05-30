package service

import (
	"log"
	"sync/atomic"

	"github.com/RaymondCode/simple-demo/dto"
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
func (u *UserService) Register(name string, password string) (res *model.User, code int32) {
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

//用户登录
func (u *UserService) Login(name string, password string) (res *model.User, code int32) {
	//验证数据长度，客户端已帮忙进行验证长度必须大于5
	if len(name) > 32 || len(password) > 32 {
		return nil, 402
	}

	user := u.userDao.QueryByName(name)
	//查询是否已存在该用户
	if user == nil {
		return nil, 404
	}
	//密码验证
	if password != user.Password {
		return nil, 405
	}
	return user, 0
}

//用户信息，返回user的model实体类
func (u *UserService) UserInfo(id int64) (res *dto.UserDto, code int32) {
	user, err := u.userDao.QueryById(id)
	if err != nil || user == nil {
		return nil, 500
	}
	if user.ID == 0 {
		return nil, 404
	}
	//封装到DTO对象
	userDto, _ := u.ToUserDto(user, false)
	return userDto, 0
}

//用户粉丝数增1或减1,id为用户id，isAdd为True为粉丝数+1
func (u *UserService) FollowerCountUpdate(id int64, isAdd bool) (res *model.User, code int32) {
	if id == 0 {
		return nil, 404
	}
	//获取原来的对象
	oldUser, err := u.userDao.QueryById(id)
	if err != nil {
		log.Println("[userService FollowerCountAddOne bug]", err)
		return nil, 500
	}
	if oldUser.ID == 0 {
		return nil, 404
	}
	//修改原本的对象的值
	if isAdd {
		atomic.AddInt64(&oldUser.FollowerCount, 1)
	} else if oldUser.FollowerCount > 0 {
		atomic.AddInt64(&oldUser.FollowerCount, -1)
	} else {
		return nil, 406
	}
	newUser := oldUser
	//传入新对象
	res, err = u.userDao.Update(newUser)
	if err != nil {
		log.Println("[userService FollowerCountAddOne bug]", err)
		return nil, 500
	}
	return res, 0
}

//User实体类转换为UserDto数据传输对象
func (u *UserService) ToUserDto(user *model.User, follow bool) (res *dto.UserDto, code int32) {
	if user.ID == 0 {
		return nil, 404
	}
	userDto := dto.UserDto{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      follow,
	}
	return &userDto, 0
}
