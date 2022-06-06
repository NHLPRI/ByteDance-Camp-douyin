package controller

import (
	"net/http"

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
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: dto.CommentDto{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	//获取URL中的token和视频id

	//查询对应的评论数据并封装

	//发送响应
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
