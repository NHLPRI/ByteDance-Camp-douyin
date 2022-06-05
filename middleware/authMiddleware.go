//gin中间件，个人理解就是过滤器或者拦截器
package middleware

import (
	"log"
	"net/http"

	"github.com/RaymondCode/simple-demo/common"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
)

//拦截验证token
func TokenValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//从URL参数中获取token
		tokenString := ctx.Query("token")

		//检查是否是Bearer Token
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 407,
				"status_msg":  "token不合法",
			})
			log.Println("[TokenValidate] tokenString is not Bearer token")
			ctx.Abort()
			return
		}

		//解析token
		token, claims, err := common.ParseToken(tokenString)
		//错误或解析后的token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 407,
				"status_msg":  "用户权限不足",
			})
			log.Println("[AuthMiddleware] err or token is not valid")
			ctx.Abort()
			return
		}

		//token验证通过
		userId := claims.ID
		db := common.GetDB()
		var user model.User
		db.First(&user, userId) //通过userId查询用户记录并封装

		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"status_code": 404,
				"status_msg":  "用户不存在",
			})
			//请求中止
			ctx.Abort()
			return
		}

		//若用户存在，则将user信息写入上下文
		ctx.Set("user", user)
		//放行
		log.Println("[Middleware tokenValidate] Middleware Opteration ")
		ctx.Next()
	}
}
