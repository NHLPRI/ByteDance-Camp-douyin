package controller

import (
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type LikeVideoResponse struct {
	Response
	VideoList []dto.VideoDto `json:"video_list"`
}

// 赞操作
func FavoriteAction(c *gin.Context) {

	User, _ := c.Get("user")
	if User == nil {
		c.JSON(http.StatusNotFound, LikeVideoResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
	}

	// 获取post数据
	v_id, ok := c.GetPostForm("video_id")
	if !ok {
		c.JSON(http.StatusNotFound, LikeVideoResponse{
			Response: Response{StatusCode: 404, StatusMsg: "视频不存在"},
		})
	}

	a_type, ok := c.GetPostForm("action_type")
	if !ok {
		c.JSON(http.StatusNotFound, LikeVideoResponse{
			Response: Response{StatusCode: 404, StatusMsg: "action_type不存在"},
		})
	}

	// 字符串转化为数字，传输类型为字符串类型
	video_id, err := strconv.ParseInt(v_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LikeVideoResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[FavoriteAction]string转换为int64出错")
	}

	action_type, err := strconv.ParseInt(a_type, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LikeVideoResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[FavoriteAction]string转换为int64出错")
	}

	likeService := service.InitLikeService()
	// 点赞操作执行
	status_code, status_msg := likeService.LikeAction(User.(model.User).ID, video_id, action_type)

	if status_code != 0 {
		c.JSON(http.StatusInternalServerError, LikeVideoResponse{
			Response: Response{StatusCode: 500, StatusMsg: status_msg},
		})
		log.Println("[FavoriteAction]执行FavoriteAction方法出错")
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: status_msg})
	}
}

// 获取赞列表
func FavoriteList(c *gin.Context) {
	User, _ := c.Get("user")
	if User == nil {
		c.JSON(http.StatusNotFound, LikeVideoResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
	}

	id := User.(model.User).ID
	likeService := service.InitLikeService()
	likes := likeService.Likes(id)

	if len(likes) == 0 {
		c.JSON(http.StatusOK, LikeVideoResponse{VideoList: nil, Response: Response{
			StatusCode: 0,
			StatusMsg:  "用户没有点赞任何的视频哦",
		}})
	} else {
		videoService := service.NewVideoServiceInstance()
		var dtoVideos []dto.VideoDto
		for i, val := range likes {
			// 根据id查询视频
			video, err := videoService.FindVideoById(val.VideoID)
			if err != nil {
				log.Println("[VideoList]err:", err)
				c.JSON(http.StatusInternalServerError, LikeVideoResponse{VideoList: nil, Response: Response{
					StatusCode: 500,
					StatusMsg:  "服务器错误",
				}})
			}

			//转换
			dto, _ := videoService.ToVideoDto(video, true)

			dtoVideos[i] = *dto

			//返回成功
			c.JSON(http.StatusOK, LikeVideoResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "返回点赞列表成功",
				},
				VideoList: dtoVideos,
			})
		}
	}
}
