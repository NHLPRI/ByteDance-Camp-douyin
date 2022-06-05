package controller

import (
	"log"
	"net/http"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/model"
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
	407: "权限不足",
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

	user, code := u.userService.Login(username, password)
	msg := codeMap[code]

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		code = 500
		log.Println("[userController Login token err]", err)
	}

	if code == 0 {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
			UserId:   user.ID,
			Token:    token,
		})
		log.Println("[userController Login] login success !")
	} else {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
		})
		log.Println("[userController Login] login failed !")
	}

}

//用户注册路由
func (u *UserController) Register(ctx *gin.Context) {
	//客户端是将用户名和密码保存到URL参数中传递
	username := ctx.Query("username")
	password := ctx.Query("password")
	log.Println("[username] ", username)
	log.Println("[password] ", password)

	//验证并创建用户
	user, code := u.userService.Register(username, password)
	msg := codeMap[code]

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		code = 500
	}
	//响应
	if code == 0 {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
			UserId:   user.ID,
			Token:    token,
		})
		log.Println("[userController register] register success !")
	} else {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
		})
		log.Println("[userController register] register failed !")
	}

}

func (u *UserController) UserInfo(ctx *gin.Context) {
	//id, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	//拦截器已经验证token，并已将请求的user对象放入上下文中
	tempUser, _ := ctx.Get("user")
	if tempUser == nil {
		ctx.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[UserInfo] can not get context attribute \"user \" ")
	}
	id := tempUser.(model.User).ID

	//token := ctx.Query("token")
	log.Println("[userController UserInfo] user id =", id)

	//获取用户DTO对象
	userDto, code := u.userService.UserInfo(id)
	msg := codeMap[code]

	if code == 0 {
		ctx.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: code},
			User:     *userDto,
		})
		log.Println("[user info] success !")
	} else {
		ctx.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: code, StatusMsg: msg},
		})
		log.Println("[user info] failed !")
	}
}
