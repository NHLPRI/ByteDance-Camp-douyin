package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/service"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/RaymondCode/simple-demo/dto"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []dto.VideoDto `json:"video_list"`
}

type VideoController struct {
	videoService service.VideoService
}

//var codeMap = map[int32]string{
//	0:   "success !",
//	402: "密码或用户名字符长度不能超过32个字符",
//	403: "用户名已存在",
//	404: "用户不存在",
//	405: "密码错误",
//	406: "违规操作",
//	500: "服务器错误",
//}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// get all publish video list by the login user
func PublishList(c *gin.Context) {
	videoController := VideoController{
		videoService: service.NewVideoServiceInstance(),
	}
	userIdStr := c.Query("user_id")
	if userIdStr == "" {
		log.Println("user_id doesn't exist")
		c.JSON(http.StatusBadRequest, VideoListResponse{
			Response: Response{
				StatusCode: -1, StatusMsg: "user_id doesn't exist",
			},
		})
		return
	}

	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	videos, err := videoController.videoService.PublishList(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, VideoListResponse{
			Response: Response{
				StatusCode: -1, StatusMsg: "server error",
			},
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0, StatusMsg: "success",
			},
			VideoList: videos,
		})
	}

}
