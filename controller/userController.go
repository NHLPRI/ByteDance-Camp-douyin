package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

var codeMap = map[int32]string{
	0:   "success !",
	402: "密码或用户名字符长度不能超过32个字符",
	403: "用户名已存在",
	404: "用户不存在",
	405: "密码错误",
	406: "违规操作",
	500: "服务器错误",
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User dto.UserDto `json:"user"`
}

type UserController struct {
	userService service.UserService
}

//初始化UserController
func InitUserController() UserController {
	return UserController{userService: service.InitUserService()}
}

//用户登录路由
func (u *UserController) Login(ctx *gin.Context) {
	//客户端是将用户名和密码保存到URL参数中传递
	username := ctx.Query("username")
	password := ctx.Query("password")
	log.Println("[username] ", username)
	log.Println("[password] ", password)

	//生成token
	token := username + password

	user, code := u.userService.Login(username, password)
	msg := codeMap[code]

	if code == 0 {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
			UserId:   user.ID,
			Token:    token,
		})
		log.Println("[login method] login success !")
	} else {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
		})
		log.Println("[login method] login failed !")
	}

}

//用户注册路由
func (u *UserController) Register(ctx *gin.Context) {
	//客户端是将用户名和密码保存到URL参数中传递
	username := ctx.Query("username")
	password := ctx.Query("password")
	log.Println("[username] ", username)
	log.Println("[password] ", password)

	//生成token
	token := username + password

	//验证并创建用户
	user, code := u.userService.Register(username, password)
	msg := codeMap[code]

	//响应
	if code == 0 {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
			UserId:   user.ID,
			Token:    token,
		})
		log.Println("[register method] register success !")
	} else {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
		})
		log.Println("[register method] register failed !")
	}

}

func (u *UserController) UserInfo(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	//token := ctx.Query("token")
	log.Println("[user id]", id)

	//验证token是否合法？逻辑放这里还是放拦截器

	//获取用户对象并封装到DTO对象
	user, code := u.userService.UserInfo(id)
	msg := codeMap[code]
	userDto := dto.UserDto{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}

	if code == 0 {
		ctx.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: code},
			User:     userDto,
		})
		log.Println("[user info] success !")
	} else {
		ctx.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
		})
		log.Println("[user info] failed !")
	}
}
