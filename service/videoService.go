package service

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"log"
	"time"
)

type VideoService struct {
	userDao  repository.UserDao
	videoDao repository.VideoDao
}

type VideoList struct {
	VideoList []dto.VideoDto `json:"video_list,omitempty"`
}

func NewVideoServiceInstance() VideoService {
	return VideoService{userDao: repository.InitUserDao(), videoDao: repository.NewVideoDaoInstance()}
}

//返回user发布的videolist
func (v *VideoService) PublishList(userId int64) ([]dto.VideoDto, error) {
	videos, err := v.videoDao.QueryByUserId(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	user, err := v.userDao.QueryById(userId)
	//pending
	author := dto.UserDto{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}

	videoList := make([]dto.VideoDto, len(videos))
	for i := 0; i < len(videos); i++ {
		videoList[i] = dto.VideoDto{
			Id:            videos[i].ID,
			PlayUrl:       videos[i].PlayURL,
			CoverUrl:      videos[i].CoverURL,
			FavoriteCount: videos[i].FavouriteCount,
			CommentCount:  videos[i].CommentCount,
			//
			IsFavorite: false,
			Author:     author,
		}
	}
	return videoList, nil

}

// 通过用户ID查找用户
func (v *VideoService) FindVideoById(id int64) (*model.Video, error) {
	video, err := v.videoDao.QueryById(id)
	if err != nil {
		return nil, err
	}
	return video, nil
}

// Video实体类转换为VideoDto数据传输对象
func (v *VideoService) ToVideoDto(video *model.Video, like bool) (res *dto.VideoDto, code int32) {
	if video.ID == 0 {
		return nil, 404
	}
	videoDto := dto.VideoDto{
		Id:            video.ID,
		Title:         video.Title,
		PlayUrl:       video.PlayURL,
		CoverUrl:      video.CoverURL,
		FavoriteCount: video.FavouriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    like,
	}
	return &videoDto, 0
}

func (v *VideoService) Feed(lasttime int64) ([]dto.VideoDto, int64, error) {
	//sql查询
	videos, err := v.videoDao.QueryBytime(lasttime)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, err
	}
	//没有视频更新了
	if len(videos) == 0 {
		return nil, 0, nil
	}
	videoList := make([]dto.VideoDto, len(videos))
	for i := 0; i < len(videos); i++ {
		user, err := v.userDao.QueryById(videos[i].UserID)
		if err != nil {
			log.Println(err.Error())
			return nil, 0, err
		}
		//找到视频作者
		author := dto.UserDto{
			Id:            user.ID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		//转化视频格式
		videoList[i] = dto.VideoDto{
			Id:            videos[i].ID,
			PlayUrl:       videos[i].PlayURL,
			CoverUrl:      videos[i].CoverURL,
			FavoriteCount: videos[i].FavouriteCount,
			CommentCount:  videos[i].CommentCount,
			//
			IsFavorite: false,
			Author:     author,
		}
	}
	return videoList, videos[len(videos)-1].Create_time, nil
}

//将发布的视频上传数据库
func (v *VideoService) Public_action(userId int64, fileName string, titleString string) error {
	var finalName string = "178.79.130.90:9000/" + common.BUCKETNAME + "/" + fileName
	var converName string = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
	log.Print(finalName)
	time := time.Now().Unix()
	err1 := v.videoDao.Create(&model.Video{
		UserID:         userId,
		PlayURL:        finalName,
		CoverURL:       converName,
		FavouriteCount: 0,
		CommentCount:   0,
		Title:          titleString,
		Create_time:    time})
	if err1 != nil {
		log.Printf("数据库上传失败")
		return err1
	}
	return nil
}
