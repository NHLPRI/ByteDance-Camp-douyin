package test

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/util"
	"testing"
)

func TestJudgeIsFollow(t *testing.T) {

	b := util.JudgeIsFollow(1, 2)
	fmt.Println(b)
}
