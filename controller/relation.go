package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
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

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {

	User, _ := c.Get("user")
	if User == nil {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction]用户不存在")
	}

	//token := c.Query("token")
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}

	f_id, ok := c.GetPostForm("to_user_id")
	if !ok {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction]用户不存在")
	}

	a_type, ok := c.GetPostForm("action_type")
	if !ok {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction]用户不存在")
	}

	follow_id, err := strconv.ParseInt(f_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserListResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[RelationAction]string转换为int64出错")
	}

	action_type, err := strconv.ParseInt(a_type, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserListResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[RelationAction]string转换为int64出错")
	}

	followService := service.InitFollowService()
	status_code, status_msg := followService.FollowAction(User.(model.User).ID, follow_id, action_type)

	if status_code != 0 {
		c.JSON(http.StatusInternalServerError, UserListResponse{
			Response: Response{StatusCode: 500, StatusMsg: status_msg},
		})
		log.Println("[RelationAction]执行FollowAction方法出错")
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

	if len(follows) == 0 {
		c.JSON(http.StatusOK, UserListResponse{UserList: nil, Response: Response{
			StatusCode: 0,
			StatusMsg:  "该用户没有关注其他人",
		}})
	} else {

		userService := service.InitUserService()

		//将User实体类转换为UserDto数据传输对象
		var dtoUsers []dto.UserDto
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

			//转换
			dto, _ := service.ToUserDto(user, true)

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

	follows := followService.Fans(id)

	if len(follows) == 0 {
		c.JSON(http.StatusOK, UserListResponse{UserList: nil, Response: Response{
			StatusCode: 0,
			StatusMsg:  "该用户没有粉丝",
		}})
	} else {

		userService := service.InitUserService()

		//将User实体类转换为UserDto数据传输对象
		var dtoUsers []dto.UserDto
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

			//true为关注,false为未关注
			var to_follow bool

			//判断是否互关
			follow := followService.FindFollowsExist(id, val.ID)
			if follow == nil {
				to_follow = false
			} else {
				to_follow = true
			}

			//转换
			dto, _ := service.ToUserDto(user, to_follow)

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
