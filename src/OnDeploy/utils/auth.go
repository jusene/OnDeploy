package utils

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthRequired(ctx *gin.Context) (user, pass string, err error) {
	authToken := ctx.GetHeader("Authorization")
	if authToken == "" {
		return "", "", errors.New("请进行服务器认证")
	}
	user, pass = AuthDecode(authToken)
	return
}

func AuthDecode(token string) (user, pass string){
	baseToken := strings.Split(token, " ")[1]
	decodeBytes, _ := base64.StdEncoding.DecodeString(baseToken)
	tokenSplit := strings.Split(string(decodeBytes), ":")
	user, pass = tokenSplit[0], tokenSplit[1]
	return
}
