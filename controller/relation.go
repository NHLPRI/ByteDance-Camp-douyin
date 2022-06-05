package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/util"
	"log"
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/dto"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []dto.UserDto `json:"user_list"`
}

//关注和取消关注
// RelationAction no practical effect, just check if token is valid

func RelationAction(c *gin.Context) {

	User, _ := c.Get("user")
	if User == nil {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction][user]用户不存在")
		return
	}

	//token := c.Query("token")
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}

	f_id := c.Query("to_user_id")
	//if !ok {
	//	c.JSON(http.StatusNotFound, UserListResponse{
	//		Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
	//	})
	//	log.Println("[RelationAction]用户[to_user_id]不存在")
	//	return
	//}
	if f_id == "" {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction]用户[to_user_id]不存在")
		return
	} else {
		log.Println("f_id:")
		log.Println(f_id)
	}

	a_type := c.Query("action_type")
	//if !ok {
	//	c.JSON(http.StatusNotFound, UserListResponse{
	//		Response: Response{StatusCode: 404, StatusMsg: "action_type不存在"},
	//	})
	//	log.Println("[RelationAction]action_type不存在")
	//	return
	//}
	if a_type == "" {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction]用户[to_user_id]不存在")
		return
	} else {
		log.Println("a_type:")
		log.Println(a_type)
	}

	follow_id, err := strconv.ParseInt(f_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserListResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[RelationAction][follow_id]string转换为int64出错")
		return
	} else {
		fmt.Println(follow_id)
	}

	action_type, err := strconv.ParseInt(a_type, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserListResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[RelationAction][action_type]string转换为int64出错")
		return
	} else {
		fmt.Println(action_type)
	}

	followService := service.InitFollowService()
	status_code, status_msg := followService.FollowAction(User.(model.User).ID, follow_id, action_type)

	if status_code != 0 {
		c.JSON(http.StatusInternalServerError, UserListResponse{
			Response: Response{StatusCode: 500, StatusMsg: status_msg},
		})
		log.Println("[RelationAction]执行FollowAction方法出错")
		return
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: status_msg})
	}

}

//获取用户关注列表
// FollowList all users have same follow list

func FollowList(c *gin.Context) {

	User, _ := c.Get("user")
	if User == nil {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[FollowList]用户不存在")
	}

	id := User.(model.User).ID
	followService := service.InitFollowService()
	follows := followService.Follows(id)

	log.Println("follows:")
	log.Println(follows)

	if len(follows) == 0 {
		c.JSON(http.StatusOK, UserListResponse{UserList: nil, Response: Response{
			StatusCode: 0,
			StatusMsg:  "该用户没有关注其他人",
		}})
	} else {

		l := len(follows)

		userService := service.InitUserService()

		//将User实体类转换为UserDto数据传输对象
		//var dtoUsers [l]dto.UserDto
		dtoUsers := make([]dto.UserDto, l)

		for i, val := range follows {

			//根据id查询用户
			user, err := userService.FindUserById(val.FollowID)
			if err != nil {
				log.Println("[FollowList]err:")
				log.Println(err)
				c.JSON(http.StatusInternalServerError, UserListResponse{UserList: nil, Response: Response{
					StatusCode: 500,
					StatusMsg:  "服务器错误",
				}})
			}

			log.Println("user:")
			log.Println(user)

			//转换
			dto, _ := service.ToUserDto(user, true)

			log.Println("dto:")
			log.Println(dto)

			dtoUsers[i] = *dto

			//返回成功
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "返回关注列表成功",
				},
				UserList: dtoUsers,
			})

		}
	}
}

//获取粉丝列表
// FollowerList all users have same follower list

func FollowerList(c *gin.Context) {

	User, _ := c.Get("user")
	if User == nil {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[FollowList]用户不存在")
	}

	id := User.(model.User).ID
	followService := service.InitFollowService()

	//查询粉丝列表

	fans := followService.Fans(id)

	//打印日志
	log.Println("fans:")
	log.Println(fans)

	if len(fans) == 0 {
		c.JSON(http.StatusOK, UserListResponse{UserList: nil, Response: Response{
			StatusCode: 0,
			StatusMsg:  "该用户没有粉丝",
		}})
	} else {

		userService := service.InitUserService()

		l := len(fans)

		//将User实体类转换为UserDto数据传输对象
		//var dtoUsers []dto.UserDto
		dtoUsers := make([]dto.UserDto, l)

		for i, val := range fans {
			//根据id查询粉丝用户
			user, err := userService.FindUserById(val.UserID)

			log.Println("user:", user)

			if err != nil {
				log.Println("[FollowList]err:")
				log.Println(err)
				c.JSON(http.StatusInternalServerError, UserListResponse{UserList: nil, Response: Response{
					StatusCode: 500,
					StatusMsg:  "服务器错误",
				}})
			}

			//true为关注,false为未关注
			//var to_follow bool

			////判断是否互关
			//follow := followService.FindFollowsExit(id, user.ID)
			//if follow == nil {
			//	to_follow = false
			//} else {
			//	to_follow = true
			//}

			to_follow := util.JudgeIsFollow(id, user.ID)

			log.Println("to_follow:", to_follow)

			//转换
			dto, _ := service.ToUserDto(user, to_follow)
			log.Println("dto:", dto)

			dtoUsers[i] = *dto

			//返回成功
			c.JSON(http.StatusOK, UserListResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "返回粉丝列表成功",
				},
				UserList: dtoUsers,
			})

		}
	}
}
