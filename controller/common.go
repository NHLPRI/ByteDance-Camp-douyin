package controller

import "github.com/RaymondCode/simple-demo/dto"

//公共的响应结构体，每个响应结构体都包含该结构体
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64       `json:"id,omitempty"`
	Author        dto.UserDto `json:"author"`
	PlayUrl       string      `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string      `json:"cover_url,omitempty"`
	FavoriteCount int64       `json:"favorite_count,omitempty"`
	CommentCount  int64       `json:"comment_count,omitempty"`
	IsFavorite    bool        `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64       `json:"id,omitempty"`
	User       dto.UserDto `json:"user"`
	Content    string      `json:"content,omitempty"`
	CreateDate string      `json:"create_date,omitempty"`
}
