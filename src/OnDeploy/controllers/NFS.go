package controllers

import (
	"OnDeploy/models"
	"OnDeploy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary NFS服务安装
// @Description NFS服务安装
// @Tags NFS服务
// @Accept json
// @Produce json
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/nfs/install [post]
func NFSInstall(ctx *gin.Context) {
	var server models.ServerDetail
	// 检查请求json
	if err := ctx.ShouldBindJSON(&server); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := utils.InstallNFS(server); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

}
