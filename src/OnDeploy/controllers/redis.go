package controllers

import (
	"OnDeploy/models"
	"OnDeploy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Redis服务安装
// @Description Redis服务安装
// @Tags Redis服务
// @Accept json
// @Produce json
// @Security basic
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/redis/install [post]
func RedisInstall(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var server models.ServerDetail
	// 检查请求json
	if err := ctx.ShouldBindJSON(&server); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}


}
