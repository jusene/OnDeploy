package controllers

import (
	"OnDeploy/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Git服务安装
// @Description Git服务安装
// @Tags Git服务
// @Accept json
// @Produce json
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/git/install [post]
func GitInstall(ctx *gin.Context) {
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
