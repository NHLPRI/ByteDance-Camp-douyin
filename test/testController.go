package test

import (
	"log"
	"net/http"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TestUser struct {
	gorm.Model
	Name      string `json:"name" gorm:"varchar(20);not null"` //绑定字段
	Telephone string `json:"telephone" gorm:"varchar(110);not null;unique"`
}

//测试数据库连接、插入、查询  (测试成功)
func TestInsertAndQuery(ctx *gin.Context) {
	//获取对数据库的操作对象
	db := common.GetDB()
	//自动创建表,不会重复创建
	db.AutoMigrate(&TestUser{})

	//接收Apifox传来的请求，绑定body中的JSON到requestUser当中
	var requestUser TestUser
	ctx.Bind(&requestUser)

	log.Println("[test] request.telephone = ", requestUser.Telephone)

	//查询是否已存在该手机号，测试查询方法
	var user TestUser
	db.Where("telephone = ?", requestUser.Telephone).First(&user)
	log.Println("[test] user.Telephone = ", user.Telephone)
	if user.ID != 0 {
		ctx.JSON(http.StatusForbidden, gin.H{
			"code": 403,
			"msg":  "手机号已存在",
		})
		return
	}

	//插入
	db.Create(&requestUser)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

//测试git
