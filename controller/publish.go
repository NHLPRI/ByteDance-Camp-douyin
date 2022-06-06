package controller

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	tokenString := c.PostForm("token")
	token, claims, err := common.ParseToken(tokenString)

	if err != nil || !token.Valid {
		c.JSON(http.StatusOK, Response{StatusCode: 500, StatusMsg: "server error"})
		return
	}
	//token验证通过
	userId := claims.ID
	db := common.GetDB()
	var user model.User
	db.First(&user, userId) //通过userId查询用户记录并封装
	//用户不存在
	if user.ID == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 404, StatusMsg: "用户不存在"})
		return
	}
	//用户存在
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 500,
			StatusMsg:  "server error",
		})
		return
	}
	//视频通过
	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%d_%d_%s", user.ID, time.Now().Unix(), filename)
	saveFile := filepath.Join("./public/", finalName)
	//original save to local
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 500,
			StatusMsg:  "server error",
		})
		return
	}
	// Upload the video/mp4 file with FPutObject
	minioClient := common.InitMinioClient()
	info, err := minioClient.FPutObject(c, common.BUCKETNAME, finalName, saveFile, minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 500,
			StatusMsg:  "server error",
		})
		return
	}
	log.Printf("Successfully uploaded %s of size %d\n", finalName, info.Size)
	//删除本地的文件
	fileerr := os.Remove(saveFile)
	if fileerr != nil {
		log.Println("file remove Error!")
		log.Printf("%s", err)
	} else {
		log.Print("file remove OK!")
	}
	//获取标题
	titleString := c.PostForm("title")
	//上传到数据库
	videoController := VideoController{
		videoService: service.NewVideoServiceInstance(),
	}
	err1 := videoController.videoService.Public_action(userId, finalName, titleString)
	if err1 != nil {
		log.Printf("[上传数据库失败]")
		c.JSON(http.StatusOK, Response{
			StatusCode: 500,
			StatusMsg:  "server error",
		})
		return
	} else {
		log.Printf("[上传数据库成功]")
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
	User, _ := c.Get("user")
	if User == nil {
		c.JSON(http.StatusNotFound, UserListResponse{
			Response: Response{StatusCode: 400, StatusMsg: "用户不存在"},
		})
		log.Println("[PublishList]用户不存在")
	}
	//token 鉴权
	userId := User.(model.User).ID
	//
	//userIdStr := c.Query("user_id")
	//if userIdStr == "" {
	//	log.Println("user_id doesn't exist")
	//	c.JSON(http.StatusBadRequest, VideoListResponse{
	//		Response: Response{
	//			StatusCode: -1, StatusMsg: "user_id doesn't exist",
	//		},
	//	})
	//	return
	//}
	//
	//userId, _ := strconv.ParseInt(userIdStr, 10, 64)
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
