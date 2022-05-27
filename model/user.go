//model层，存放实体类，一实体类一个表
package model

type User struct {
	ID            int64  `json:"id" gorm:"primary_key"`
	Name          string `json:"name" gorm:"type:varchar(50);not null;unique"`
	Password      string `json:"password" gorm:"type:varchar(500);not null"`
	FollowCount   int64  `json:"follow_count" gorm:"not null"`
	FollowerCount int64  `json:"follower_count" gorm:"not null"`
}
