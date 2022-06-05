package controller

import "github.com/RaymondCode/simple-demo/dto"

//本文件用于存储测试用的临时数据

var DemoVideos = []dto.VideoDto{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 666,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "http://10.30.1.205:8080/static/meeting.mp4", //访问服务器上的资源
		CoverUrl:      "http://10.30.1.205:8080/static/meeting.jpg",
		FavoriteCount: 101,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []dto.CommentDto{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "国家防脱发研究院yyds",
		CreateDate: "05-01",
	},
}

var DemoUser = dto.UserDto{
	Id:            1,
	Name:          "jiayuan01",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
