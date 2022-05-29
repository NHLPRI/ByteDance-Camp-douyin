//model层，实体类，一表对应一个结构体
package model

import (
	"time"
)

/*
* Name:用户名，上限32个字符
* Password：密码，上限32个字符
 */
type User struct {
	ID            int64  `json:"id" gorm:"primary_key"`
	Name          string `json:"name" gorm:"type:varchar(32);not null;unique;index:uniqueIndex"`
	Password      string `json:"password" gorm:"type:varchar(32);not null"`
	FollowCount   int64  `json:"follow_count" gorm:"not null"`
	FollowerCount int64  `json:"follower_count" gorm:"not null"`
}

type Video struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	UserID         int64  `json:"user_id" gorm:"not null"`
	PlayURL        string `json:"play_url" gorm:"type:varchar(500);not null"`
	CoverURL       string `json:"cover_url" gorm:"type:varchar(500);not null"`
	FavouriteCount int64  `json:"favourite_count" gorm:"not null"`
	CommentCount   int64  `json:"common_count" gorm:"not null"`
	Title          string `json:"title" gorm:"type:varchar(255);not null"`
}

//关注表
type Follow struct {
	ID     int64 `gorm:"primary_key"`
	UserID int64 `gorm:"not null"`
	FansID int64 `gorm:"not null"`
}

//点赞表
type Like struct {
	ID      int64 `gorm:"primary_key"`
	UserID  int64 `gorm:"not null"`
	VideoID int64 `gorm:"not null"`
}

type Comment struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	UserID    int64  `json:"user_id" gorm:"not null"`
	VideoID   int64  `json:"video_id" gorm:"nor null"`
	Content   string `gorm:"type:longText"`
	Floor     uint
	CreatedAt time.Time
}
