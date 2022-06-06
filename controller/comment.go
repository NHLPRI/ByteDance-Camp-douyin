package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/RaymondCode/simple-demo/dto"
	"github.com/gin-gonic/gin"
)

//评论列表响应对象
type CommentListResponse struct {
	Response
	CommentList []dto.CommentDto `json:"comment_list,omitempty"`
}

//评论操作响应对象
type CommentActionResponse struct {
	Response
	Comment dto.CommentDto `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	user, _ := c.Get("user")
	if user == nil {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction][user]用户不存在")
		return
	}
	u, ok := user.(model.User)
	if !ok{
		log.Println("User断言失败")
	}
	log.Println(user)
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil{
		log.Println("获取videoID失败")
	}
	action_type, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if action_type != 1{
		log.Println("获取action_type失败 action_type 不是1")
	}
	comment_text := c.Query("comment_text")
	if len(comment_text) == 0{
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
		log.Println("[RelationAction][user]用户不存在")
		return
	}
	comment := &model.Comment{
		ID:        0,
		UserID:    u.ID,
		VideoID:   videoID,
		Content:   comment_text,
		Floor:     0,
		CreatedAt: time.Now(),
	}
	commentService := service.InitCommentService()
	status_code, status_msg := commentService.Create(comment)

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

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	//获取URL中的token和视频id
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	log.Println(videoID)
	if err != nil{
		c.JSON(http.StatusNotFound, CommentListResponse{
			Response:    Response{},
			CommentList: nil,
		})
	}
	commentService := service.InitCommentService()
	comments := commentService.Find(int64(videoID))
	log.Println(comments)
	// 没有评论提醒下
	if len(comments) == 0{
		c.JSON(http.StatusOK, CommentListResponse{CommentList: nil, Response: Response{
			StatusCode: 0,
			StatusMsg:  "该用户没有关注其他人",
		}})
	}else {
		commentsLen := len(comments)
		dtoComments := make([]dto.CommentDto, commentsLen)
		// 转换格式
		for i, val := range comments{
			user, err := commentService.FindUserById(val.UserID)
			log.Println(user)
			if err != nil {
				log.Println("[FollowList]err:")
				log.Println(err)
				c.JSON(http.StatusInternalServerError, UserListResponse{UserList: nil, Response: Response{
					StatusCode: 500,
					StatusMsg:  "服务器错误",
				}})
			}
			dtoComments[i] = dto.CommentDto{
				Id:         val.ID,
				User:       dto.UserDto{
					Id:            user.ID,
					Name:          user.Name,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
					IsFollow:      false,
				},
				Content:    val.Content,
				CreateDate: val.CreatedAt.String(),
			}
		}
		//发送响应
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: dtoComments,
		})
	}
}
