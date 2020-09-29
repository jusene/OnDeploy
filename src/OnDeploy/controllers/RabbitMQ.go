package controllers

import (
	"OnDeploy/models"
	"OnDeploy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary RabbitMQ服务安装
// @Description RabbitMQ服务安装
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/install [post]
func RabbitMQInstall(ctx *gin.Context) {
	var server models.ServerDetail
	// 检查请求json
	if err := ctx.ShouldBindJSON(&server); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := utils.InstallRabbitMQ(server); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%s rabbitmq服务安装成功", server.Address),
	})
}

