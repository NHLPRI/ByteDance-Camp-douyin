package test

import (
	"fmt"
	_ "github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"testing"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
)

func Test_FollowerCountUpdate(t *testing.T) {
	fmt.Println("begin run")
	db := common.InitDbConnection()
	defer db.Close()

	userService := service.InitUserService()
	_, code := userService.FollowerCountUpdate(3, true)
	fmt.Println("[test over] ", code)
}

func Test_FollowDao(t *testing.T) {

	//dsn := "root:Lgh1906200123.@tcp(119.29.184.229:3306)/lghTest?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println("test followDao")
	db := common.InitDbConnection()
	defer db.Close()

	//
	followDao := repository.InitFollowDao()

	//err:=followDao.Create(&model.Follow{
	//	UserID: 1,
	//	FollowID: 12,
	//})
	//
	//if err != nil{
	//	fmt.Println(err)
	//}

	//follow := followDao.Find(111, 2)
	////
	//if follow == nil {
	//	fmt.Println("不存在这条记录")
	//} else {
	//	fmt.Println(*follow)
	//}

	//
	//follow=followDao.Find(2,1)
	//
	//if &follow ==nil {
	//	fmt.Println("不存在这条记录")
	//}else{
	//	fmt.Println(follow)
	//}
	//
	//fmt.Println("关注列表")
	////var follows [] model.Follow
	//follows:=followDao.SelectFollows(1)
	//
	//for _,val:=range follows{
	//	fmt.Println(val)
	//}

	err := followDao.Delete(1, 2)
	if err != nil {
		fmt.Println(err)
	}

}
