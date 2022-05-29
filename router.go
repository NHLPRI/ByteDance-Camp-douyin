package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/test"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	//路由分组
	apiRouter := r.Group("/douyin")

	// basic apis
	//user apis
	userController := controller.InitUserController()
	//apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.GET("/user/", userController.UserInfo)
	//apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/register/", userController.Register)
	//apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/user/login/", userController.Login)

	//video apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)

	//测试路由
	testRouter := r.Group("/test")
	testRouter.POST("/testIAQ/", test.TestInsertAndQuery)
}
