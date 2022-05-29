package model

import (
	"time"
)

type User struct {
	ID            int64   `json:"id" gorm:"primary_key"`
	Name          string  `json:"name" gorm:"type:varchar(32);not null;unique;index:uniqueIndex"`
	Password      string  `json:"password" gorm:"type:varchar(32);not null"`
	FollowCount   int64   `json:"follow_count" gorm:"not null"`
	FollowerCount int64   `json:"follower_count" gorm:"not null"`
	Videos        []Video `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Video struct {
	ID             int64 `json:"id" gorm:"primary_key"`
	UserID         int64
	PlayURL        string
	CoverURL       string
	FavouriteCount int64
	CommentCount   int64
	Title          string
}

type Follow struct {
	ID     int64 `gorm:"primary_key;AUTO_INCREMENT"`
	UserID int64
	FansID int64
}
type Like struct {
	ID      int64 `gorm:"primary_key;AUTO_INCREMENT"`
	UserID  int64
	VideoID int64
}

type Comment struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	UserID    int64
	VideoID   int64
	Content   string `gorm:"type:longText"`
	Floor     uint
	CreatedAt time.Time
}
