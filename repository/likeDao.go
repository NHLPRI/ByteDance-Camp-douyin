// 点赞CRUD
package repository

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/jinzhu/gorm"
)

// Like的持久层对象
type LikeDao struct {
	db *gorm.DB //数据库对象属性
}

// 初始化LikeDao对象
func InitLikeDao() LikeDao {
	db := common.GetDB()
	return LikeDao{db: db}
}

/**
添加点赞记录
*/
func (l *LikeDao) Create(like *model.Like) (*model.Like, error) {
	err := l.db.Create(like).Error
	if err != nil {
		return nil, err
	}
	return like, nil
}

/**
删除点赞记录,取消赞
*/
func (l *LikeDao) Delete(id int64) error {
	err := l.db.Delete(id).Error
	return err
}

/**
返回用户点赞的视频列表，所有赞视频
*/
func (l *LikeDao) SelectLikeVideos(id int64) []model.Like {

	var likeVideos []model.Like

	l.db.Where("user_id = ?", id).Find(&likeVideos)

	return likeVideos
}

/**
根据video_id,user_id查询是否存在这条记录
*/
func (l *LikeDao) Find(user_id int64, video_id int64) *model.Like {
	var like = model.Like{}
	l.db.Where("user_id = ? AND video_id = ?", user_id, video_id).First(&like)

	if like.ID != 0 {
		return &like
	}
	return nil
}
