package service

import (
	"github.com/RaymondCode/simple-demo/dto"
	"github.com/RaymondCode/simple-demo/repository"
	"log"
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
