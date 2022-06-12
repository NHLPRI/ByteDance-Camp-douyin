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
	//如果视频已经播放完了，返回null的视频数组和null的nexttime
	if nextTime == 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: nil,
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  nextTime,
	})
}
