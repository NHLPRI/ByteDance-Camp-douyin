package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"log"
	"sync/atomic"
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

	// 点赞操作
	if action_type == 1 {
		// 往like表中插入一条数据
		err, _ := l.likeDao.Create(&model.Like{
			UserID:  user_id,
			VideoID: video_id,
		})

		if err != nil {
			log.Println(err)
			return 500, "点赞失败"
		}

		// 更新视频表的的点赞数
		video, error := l.videoDao.QueryById(video_id)
		if error != nil {
			log.Println("[add video_favorite_count error]", error)
		}
		if video.ID == 0 {
			log.Println("Video don't exit")
			return 404, "此视频不存在"
		}
		atomic.AddInt64(&video.FavouriteCount, 1)

		return 0, "点赞成功"

	} else if action_type == 2 { // 取消赞

		// 得到记录
		like := l.likeDao.Find(user_id, video_id)

		if like == nil {
			return 500, "本条点赞记录已经不存在"
		}

		id := like.ID
		err := l.likeDao.Delete(id)
		if err != nil {
			return 500, "取消赞数据库删除操作失败"
		}

		// 将视频表的关点赞数量减1
		video, err := l.videoDao.QueryById(video_id)
		if err != nil {
			log.Println("[delete video_favorite_count error]", err)
			return 500, "视频表减少点赞数失败"
		}
		if video.ID == 0 {
			log.Println("Video don't exit")
			return 404, "视频已经不存在不存在"
		}
		atomic.AddInt64(&video.FavouriteCount, -1)

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
