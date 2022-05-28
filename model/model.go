package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name          string
	Password      string
	FollowCount   int64
	FollowerCount int64
	Videos        []Video `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Video struct {
	gorm.Model
	UserID         uint
	PlayURL        string
	CoverURL       string
	FavouriteCount int64
	CommentCount   int64
	IsFavourite    bool
	Title          string
}

type Follow struct {
	ID     uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserID uint
	FansID uint
}
type Like struct {
	ID      uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserID  uint
	VideoID uint
}

type Comment struct {
	gorm.Model
	UserID  uint
	VideoID uint
	Content string `gorm:"type:longText"`
	Floor   uint
}
