package common

import (
	"errors"
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("dousheng_secret_key")

//定义payload部分
type Claims struct {
	model.User         //私有字段
	jwt.StandardClaims //官方字段
}

func ReleaseToken(user *model.User) (token string, err error) {
	if user == nil || user.ID == 0 {
		return "", errors.New("user model is null or user is not exists")
	}
	//设置过期时间: one week
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	//设置payload部分
	claims := &Claims{
		User: model.User{
			ID:            user.ID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //exp，过期时间
			IssuedAt:  time.Now().Unix(),     //iat，签发时间
			Issuer:    "dousheng",            //iss，jwt签发者
			Subject:   "user token",          //sub，主题
		},
	}

	//指定签名算法，生成token对象
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//根据签名算法和密钥，对header和payload部分进行签名，并通过Base64URL生成token字符串
	//得到的token字符串是未加密的，header和payload部分可见
	tokenString, err := tokenObj.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析token,从token中解析出claims然后返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	//得到claims并返回
	return token, claims, err
}
