package util

import (
	"github.com/RaymondCode/simple-demo/service"
	"log"
)

//判断是否互关
//user_id:关注者
//follow_id:被关注者
//如果关注了返回true,否则返回false

func JudgeIsFollow(user_id int64, follow_id int64) bool {

	log.Println("judgeIsFollow running...")
	//db:= common.InitDbConnection()
	//defer db.Close()
	//
	//followDao:=repository.InitFollowDao()
	//follow:=followDao.Find(user_id,follow_id)
	//if follow==nil{
	//	return false
	//}else {
	//	return true
	//}

	followService := service.InitFollowService()
	f := followService.FindFollowsExist(user_id, follow_id)
	if f == nil {
		return false
	} else {
		return true
	}

}
