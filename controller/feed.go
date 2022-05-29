package controller

import (
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/dto"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []dto.VideoDto `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
