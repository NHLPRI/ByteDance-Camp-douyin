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
	log.Println("User:", User)
	if User == nil {
		c.JSON(http.StatusNotFound, LikeVideoResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
	}

	// 获取url数据
	v_id := c.Query("video_id")
	a_type := c.Query("action_type")

	// 字符串转化为数字，传输类型为字符串类型
	video_id, err := strconv.ParseInt(v_id, 10, 64)
	log.Println("video_id", video_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LikeVideoResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[FavoriteAction]string转换为int64出错")
	}

	action_type, err := strconv.ParseInt(a_type, 10, 64)
	log.Println("action_type", action_type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LikeVideoResponse{
			Response: Response{StatusCode: 500, StatusMsg: "服务器错误"},
		})
		log.Println("[FavoriteAction]string转换为int64出错")
	}
	log.Println("Uid :", User.(model.User).ID)

	likeService := service.InitLikeService()
	// 点赞操作执行
	status_code, status_msg := likeService.LikeAction(User.(model.User).ID, video_id, action_type)
	log.Println("controller_code:", status_code)
	log.Println("controller_msg:", status_msg)

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
	log.Println("登录用户：", User)
	if User == nil {
		c.JSON(http.StatusNotFound, LikeVideoResponse{
			Response: Response{StatusCode: 404, StatusMsg: "用户不存在"},
		})
	}

	id := User.(model.User).ID
	likeService := service.InitLikeService()
	likes := likeService.Likes(id)

	if len(likes) == 0 {
		log.Println("没有喜欢的视频！")
		c.JSON(http.StatusOK, LikeVideoResponse{VideoList: nil, Response: Response{
			StatusCode: 0,
			StatusMsg:  "用户没有点赞任何的视频哦",
		}})
	} else {
		log.Println("有喜欢的视频！", len(likes))
		videoService := service.NewVideoServiceInstance()
		var dtoVideos []dto.VideoDto
		for i, val := range likes {
			// 根据id查询视频
			log.Println("当前循环的视频id：", val.VideoID)
			video, err := videoService.FindVideoById(val.VideoID)
			log.Println("当前循环的视频具体信息：", video)
			if err != nil {
				log.Println("[VideoList]err:", err)
				c.JSON(http.StatusInternalServerError, LikeVideoResponse{VideoList: nil, Response: Response{
					StatusCode: 500,
					StatusMsg:  "服务器错误",
				}})
			}

			//转换
			dto, _ := videoService.ToVideoDto(video, true)
			log.Println("转换的video dto：", dto)

			dtoVideos = append(dtoVideos, *dto)
			log.Println("当前存储的视频流对象", dtoVideos[i])

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
