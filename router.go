package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/middleware"
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
	apiRouter.GET("/user/", middleware.TokenValidate(), userController.UserInfo)
	//apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/register/", userController.Register)
	//apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/user/login/", userController.Login)

	//video apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/publish/action/", middleware.TokenValidate(), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.TokenValidate(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.TokenValidate(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.TokenValidate(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.TokenValidate(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.TokenValidate(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.TokenValidate(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.TokenValidate(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.TokenValidate(), controller.FollowerList)

	//测试路由
	testRouter := r.Group("/test")
	testRouter.POST("/testIAQ/", test.TestInsertAndQuery)
}
