package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"log"
)

type LikeService struct {
	userDao  repository.UserDao
	videoDao repository.VideoDao
	likeDao  repository.LikeDao
}

// 初始化LikeService结构体
func InitLikeService() LikeService {
	userDao := repository.InitUserDao()
	videoDao := repository.NewVideoDaoInstance()
	likeDao := repository.InitLikeDao()
	log.Println("[InitLikeService func] success !")
	return LikeService{
		userDao:  userDao,
		videoDao: videoDao,
		likeDao:  likeDao,
	}
}

/**
点赞操作
*/
func (l *LikeService) LikeAction(user_id int64, video_id int64, action_type int64) (status_code int32, status_msg string) {

	log.Println("service_user_id:", user_id)
	log.Println("service_video_id:", video_id)
	log.Println("service_action_type", action_type)
	// 点赞操作
	if action_type == 1 {
		// 往like表中插入一条数据
		_, err := l.likeDao.Create(&model.Like{
			UserID:  user_id,
			VideoID: video_id,
		})

		if err != nil {
			log.Println(err)
			return 500, "点赞记录添加到likes表失败！"
		} else {
			log.Println("点赞记录添加到likes表成功！")
		}

		// 更新视频表的的点赞数+1
		video, _ := l.videoDao.QueryById(video_id)
		log.Println("待更新的video：", video)
		err = l.videoDao.Update(&model.Video{
			FavouriteCount: video.FavouriteCount + 1,
			ID:             video.ID,
			UserID:         video.UserID,
			PlayURL:        video.PlayURL,
			CoverURL:       video.CoverURL,
			CommentCount:   video.CommentCount,
			Title:          video.Title,
		})
		if err != nil {
			log.Println("[add video_favorite_count error]", err)
		}
		if video.ID == 0 {
			log.Println("Video don't exit")
			return 404, "此视频不存在"
		}

		return 0, "点赞成功"

	} else if action_type == 2 { // 取消赞

		// 得到点赞记录
		like := l.likeDao.Find(user_id, video_id)

		if like == nil {
			return 500, "本条点赞记录已经不存在"
		}

		id := like.ID
		log.Println("取消赞id", id)
		err := l.likeDao.Delete(id)
		if err != nil {
			return 500, "取消赞数据库删除操作失败"
		}

		// 更新视频表的的点赞数-1
		video, _ := l.videoDao.QueryById(video_id)
		log.Println("待更新的video：", video)
		err = l.videoDao.Update(&model.Video{
			FavouriteCount: video.FavouriteCount - 1,
			ID:             video.ID,
			UserID:         video.UserID,
			PlayURL:        video.PlayURL,
			CoverURL:       video.CoverURL,
			CommentCount:   video.CommentCount,
			Title:          video.Title,
		})
		if err != nil {
			log.Println("[reduce video_favorite_count error]", err)
		}
		if video.ID == 0 {
			log.Println("Video don't exit")
			return 404, "此视频不存在"
		}

		return 0, "取消点赞成功"

	} else {
		return 300, "客户端存在改包逻辑漏洞！"
	}
}

/**
返回用户的所有点赞视频列表
*/
func (l *LikeService) Likes(id int64) []model.Like {
	likeVideos := l.likeDao.SelectLikeVideos(id)
	return likeVideos
}
