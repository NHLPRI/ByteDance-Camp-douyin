//数据传输对象，用于响应和业务层之间的数据传输
package dto

type UserDto struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type VideoDto struct {
	Id            int64   `json:"id,omitempty"`
	Author        UserDto `json:"author"`
	PlayUrl       string  `json:"play_url,omitempty"`
	CoverUrl      string  `json:"cover_url,omitempty"`
	FavoriteCount int64   `json:"favorite_count,omitempty"`
	CommentCount  int64   `json:"comment_count,omitempty"`
	IsFavorite    bool    `json:"is_favorite,omitempty"`
}

type CommentDto struct {
	Id         int64   `json:"id,omitempty"`
	User       UserDto `json:"user"`
	Content    string  `json:"content,omitempty"`
	CreateDate string  `json:"create_date,omitempty"`
}
