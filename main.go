package main

import (
	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	//建立数据库连接
	db := common.InitDbConnection()
	defer db.Close()

	//获取gin路由引擎
	r := gin.Default()
	//初始化路由
	initRouter(r)
	//获取服务器监听端口
	port := viper.GetString("server.port")
	//运行服务器
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
