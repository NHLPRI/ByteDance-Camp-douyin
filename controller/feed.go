package controller

import (
	//"log"
	"net/http"
	"strconv"
	"time"

	//"time"

	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []dto.VideoDto `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
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

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	videoController := VideoController{
		videoService: service.NewVideoServiceInstance(),
	}
	last_timeStr := c.Query("latest_time")
	//log.Printf(last_timeStr)
	last_time, _ := strconv.ParseInt(last_timeStr, 10, 64)
	DemoVideos, nextTime, err := videoController.videoService.Feed(last_time)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 500, StatusMsg: "server error"},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  nextTime,
	})
}
