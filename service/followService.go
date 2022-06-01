package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"log"
	"sync/atomic"
)

//构建followService结构体

type followService struct {
	followDao repository.FollowDao
	userDao   repository.UserDao
}

//初始化followService结构体

func InitFollowService() followService {

	followDao := repository.InitFollowDao()
	userDao := repository.InitUserDao()
	log.Println("[InitFollowService func] success !")
	return followService{
		followDao: followDao,
		userDao:   userDao,
	}

}

//关注操作

func (f *followService) FollowAction(user_id int64, follow_id int64, action_type int64) (status_code int64, status_msg string) {

	if action_type == 1 { //关注
		//往follows表中插入一条数据
		err := f.followDao.Create(&model.Follow{
			UserID:   user_id,
			FollowID: follow_id,
		})
		if err != nil {
			log.Println(err)
			return 500, "关注失败"
		}

		//更新user_id的关注数

		user, err := f.userDao.QueryById(user_id)
		if err != nil {
			log.Println("[FollowAction bug]", err)
		}
		if user.ID == 0 {
			log.Println("User don't exit")
			return 404, "用户不存在"
		}
		atomic.AddInt64(&user.FollowCount, 1)

		//更新follow_id的粉丝

		user, err = f.userDao.QueryById(follow_id)
		if err != nil {
			log.Println("[FollowAction bug]", err)
		}
		if user.ID == 0 {
			log.Println("User don't exit")
			return 404, "用户不存在"
		}
		atomic.AddInt64(&user.FollowerCount, 1)

		return 0, "关注成功"

	} else if action_type == 2 { //取消关注

		//首先根据user_id以及follow_id得到一个记录
		follow := f.followDao.Find(user_id, follow_id)

		if follow == nil {
			return 500, "记录不存在"
		}

		id := follow.ID
		err := f.followDao.Delete(id)
		if err != nil {
			return 500, "操作失败"
		}

		//将user_id的关注数量减1

		user, err := f.userDao.QueryById(user_id)
		if err != nil {
			log.Println("[FollowAction bug]", err)
			return 500, "操作失败"
		}
		if user.ID == 0 {
			log.Println("User don't exit")
			return 404, "用户不存在"
		}
		atomic.AddInt64(&user.FollowCount, -1)

		//将follow_id的粉丝数量减1

		user, err = f.userDao.QueryById(follow_id)
		if err != nil {
			log.Println("[FollowAction bug]", err)
			return 500, "操作失败"
		}
		if user.ID == 0 {
			log.Println("User don't exit")
			return 404, "用户不存在"
		}
		atomic.AddInt64(&user.FollowerCount, 1)

		return 0, "取消关注成功"

	} else {
		return 300, "客户端非法操作"
	}

}

//查询关注列表

func (f *followService) Follows(id int64) []model.Follow {

	follows := f.followDao.SelectFollows(id)
	return follows

}

//查询粉丝列表

func (f *followService) Fans(id int64) []model.Follow {

	fans := f.followDao.SelectFans(id)
	return fans

}

//根据user_id以及follow_id查询是否存在这一条记录

func (f *followService) FindFollowsExit(user_id int64, to_user_id int64) *model.Follow {
	follow := f.followDao.Find(user_id, to_user_id)
	return follow
}
