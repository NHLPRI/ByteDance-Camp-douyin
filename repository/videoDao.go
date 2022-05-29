package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/jinzhu/gorm"
	"log"
)

type VideoDao struct {
	db *gorm.DB //数据库对象属性
}

func (v *VideoDao) Create(video *model.Video) error {
	err := v.db.Create(video).Error
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (v *VideoDao) QueryById(id int64) (*model.Video, error) {
	//log.Println("[userDao queryById]")
	var video model.Video
	err := v.db.First(&video, id).Error
	if err != nil {
		//log.Println(err.Error())
		return nil, err
	}
	return &video, nil
}

func (v *VideoDao) Update(newVideo *model.Video) error {
	err := v.db.Save(newVideo).Error
	if err != nil {
		//log.Println(err.Error())
		return err
	}
	return nil
}

func (v *VideoDao) QueryByUserId(userId int64) ([]model.Video, error) {
	var videos []model.Video
	err := v.db.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {

		return nil, err
	}
	return videos, nil

}
