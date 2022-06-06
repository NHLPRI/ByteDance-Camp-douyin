package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"log"
)

type commentService struct{
	commentDao repository.CommentDao
	userDao   repository.UserDao
}



//初始化followService结构体

func InitCommentService() commentService {
	userDao := repository.InitUserDao()
	commentDao := repository.InitCommentDao()
	log.Println("[InitFollowService func] success !")
	return commentService{
		commentDao: commentDao,
		userDao:   userDao,
	}
}

// Create 新增评论操作
// 失败返回对应的错误信息
func (c *commentService) Create(comment *model.Comment) (status_code int64, status_msg string) {
	if err := c.commentDao.Create(comment); err != nil {
		log.Println(err)
		return 500, "新增评论失败"}
	return 0, "评论成功"
}

// Find 查询评论操作
// 失败返回对应的错误信息

func (c *commentService) Find(id int64) []model.Comment {
	comments := c.commentDao.Find(id)
	return comments
}

//通过用户ID查找用户

func (c *commentService) FindUserById(id int64) (*model.User, error) {
	log.Println("[userService FindUserById]")
	user, err := c.userDao.QueryById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}