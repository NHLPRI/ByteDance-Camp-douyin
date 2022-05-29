package test

import (
	"fmt"
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
