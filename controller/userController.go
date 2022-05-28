package controller

import (
	"log"
	"net/http"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

var codeMap = map[int32]string{
	0:   "success !",
	401: "密码或用户名不能为空",
	402: "密码或用户名字符长度不能超过32个字符",
	403: "用户名已存在",
	500: "服务器错误",
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
	//客户端将登录页面的username和password封装到URL参数中
	//获取参数
	//username := ctx.Query("username")
	//password := ctx.Query("password")

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
		log.Panicln("[register method] register failed !")
	}

}
