//model层，存放实体类以及POJO
package model

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

type UserDTO struct {
}
