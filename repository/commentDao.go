package repository

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/jinzhu/gorm"
	"log"
)

type CommentDao struct {
	db *gorm.DB
}

func InitCommentDao() CommentDao {
	//自动创建follow表
	db := common.GetDB()
	db.AutoMigrate(&model.Comment{})

	log.Println("[InitCommentDao func] success")
	return CommentDao{db: db}
}

// 增加评论
func (u *CommentDao) Create(comment *model.Comment) error {
	err := u.db.Create(comment).Error
	return err
}

// 查询评论
func (u *CommentDao) Find(id int64) []model.Comment {
	var comments []model.Comment
	u.db.Where("video_id = ?", id).Find(&comments)
	//fmt.Println(r.Error.Error())
	return comments
}