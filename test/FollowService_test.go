package test

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/service"
	"testing"
)

func Test_FollowService(t *testing.T) {
	fmt.Println("begin run")
	db := common.InitDbConnection()
	defer db.Close()

	followService := service.InitFollowService()
	//followService.FollowAction(1,2,1)
	follows := followService.Follows(0)

	if len(follows) == 0 {
		fmt.Println("该用户没有关注其他人")
	} else {
		for _, val := range follows {
			fmt.Println(val)
		}
	}

}
