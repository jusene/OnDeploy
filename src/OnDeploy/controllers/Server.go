package controllers

import (
	"OnDeploy/models"
	"OnDeploy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// @Summary 单台服务器初始化
// @Description 单台服务器初始化
// @Tags 服务器
// @Accept json
// @Produce json
// @Security basic
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /server/init [post]
func ServerInit(ctx *gin.Context) {
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

	// 初始化服务器
	if err := utils.InitServer(server, user, pass); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%s 服务器初始化成功", server.Address),
	})
	return
}

// @Summary 批量服务器初始化
// @Description 批量服务器初始化
// @Tags 服务器
// @Accept json
// @Produce json
// @Security basic
// @Param server body models.ServersDetail true "servers"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /servers/init [post]
func ServersInit(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var servers models.ServersDetail
	// 检查请求json
	if err := ctx.ShouldBindJSON(&servers); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// 并发初始化服务器
	var wg sync.WaitGroup
	errChan := make(chan string, 5)
	for _, server := range servers {
		wg.Add(1)
		go func(server models.ServerDetail, errChan chan string) {
			defer wg.Done()
			utils.InitServers(server, user, pass, errChan)
		}(server, errChan)
	}
	wg.Wait()

	for e := range errChan {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: e,
		})
		close(errChan)
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("服务器初始化成功"),
	})
	return
}