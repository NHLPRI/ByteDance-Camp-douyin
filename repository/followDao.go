package repository

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/jinzhu/gorm"
	"log"
)

// follow的持久层对象

type FollowDao struct {
	db *gorm.DB
}

// 初始化followDB对象

func InitFollowDao() FollowDao {
	//自动创建follow表
	db := common.GetDB()
	db.AutoMigrate(&model.Follow{})

	log.Println("[InitFollowDao func] success")
	return FollowDao{db: db}
}

/**
插入follow字段,返回值为error
*/

func (u *FollowDao) Create(follow *model.Follow) error {
	err := u.db.Create(follow).Error
	return err
}

/**
根据user_id,to_user_id查询是否存在这条记录
*/

func (u *FollowDao) Find(user_id int64, to_user_id int64) *model.Follow {
	var follow = model.Follow{}
	u.db.Where("user_id = ? AND follow_id = ?", user_id, to_user_id).First(&follow)
	//fmt.Println(r.Error.Error())

	if follow.ID != 0 {
		return &follow
	}

	return nil
}

/**
删除follow字段,返回值为error
*/

func (u *FollowDao) Delete(user_id int64, follow_id int64) error {

	//err:=u.db.Delete(follow).Error
	////return err
	//err := u.db.Table("follows").Delete(user_id)
	//return err
	err := u.db.Table("follows").Where("user_id = ? and follow_id = ?", user_id, follow_id).Delete(&model.Follow{}).Error
	return err
}

/**
返回所有关注的用户
*/

func (u *FollowDao) SelectFollows(id int64) []model.Follow {

	var follows []model.Follow

	u.db.Where("user_id = ?", id).Find(&follows)

	return follows
}

/**
返回所有粉丝
*/

func (u *FollowDao) SelectFans(id int64) []model.Follow {

	var follows []model.Follow

	u.db.Where("follow_id = ?", id).Find(&follows)

	return follows
}
