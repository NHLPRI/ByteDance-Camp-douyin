package service

//存放公用，各个service都可能会用到的函数

import (
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/model"
)

//User实体类转换为UserDto数据传输对象
func ToUserDto(user *model.User, follow bool) (res *dto.UserDto, code int32) {
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
