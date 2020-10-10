package controllers

import (
	"OnDeploy/models"
	"OnDeploy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// @Summary NFS服务安装
// @Description NFS服务安装
// @Tags NFS服务
// @Accept json
// @Produce json
// @Security basic
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/nfs/install [post]
func NFSInstall(ctx *gin.Context) {
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

	if err := utils.InstallNFS(server, user, pass); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%s nfs服务安装成功", server.Address),
	})
}

// @Summary NFS服务创建
// @Description NFS服务创建
// @Tags NFS服务
// @Accept json
// @Produce json
// @Security basic
// @Param nfs body models.NFSPath true "nfs"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/nfs/path/add [put]
func NFSPathAdd(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}
	var nfs models.NFSPath
	if err := ctx.ShouldBindJSON(&nfs); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var nfspath utils.NFS
	n := &utils.NFSObj{nfs, user, pass}
	nfspath = n
	if err := nfspath.AddPath(nfs); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%s NFS路径添加完成", nfs.PATH),
	})
}

// @Summary NFS服务查看
// @Description NFS服务查看
// @Tags NFS服务
// @Accept json
// @Produce json
// @Security basic
// @Param address path string true "address"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/nfs/path/lst/{address} [get]
func NFSPathLst(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	address := ctx.Param("address")
	var nfs models.NFSPath
	nfs.Address = address
	nfs.Port = 22
	var nfspath utils.NFS
	n := &utils.NFSObj{nfs, user, pass}
	nfspath = n
	ret, err := nfspath.LstPath()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	var nfsobj models.NFSPath
	var nfsSlice []models.NFSPath
	for _, path := range strings.Split(ret, "\n") {
		pathSplit := strings.Split(path, " ")
		if len(pathSplit) == 1 {
			continue
		}
		nfsobj.PATH = pathSplit[0]
		nfsobj.ACL = strings.Split(pathSplit[1], "(")[0]
		nfsobj.Param = strings.Fields(strings.Replace(strings.TrimRight(strings.TrimLeft(pathSplit[1], ".*("), ").*"), ",", " ", -1))
		nfsSlice = append(nfsSlice, nfsobj)
	}

	ctx.JSON(http.StatusOK, nfsSlice)
}

// @Summary NFS服务删除
// @Description NFS服务删除
// @Tags NFS服务
// @Accept json
// @Produce json
// @Security basic
// @Param nfs body models.NFSInfo true "nfs"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/nfs/path/del [delete]
func NFSPathDel(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var nfsinfo models.NFSInfo
	if err := ctx.ShouldBindJSON(&nfsinfo); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var nfs models.NFSPath
	nfs.Address = nfsinfo.Address
	nfs.PATH = nfsinfo.PATH
	nfs.Port = nfsinfo.Port
	var nfspath utils.NFS
	n := &utils.NFSObj{nfs, user, pass}
	nfspath = n
	if err := nfspath.DelPath(nfs); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%s NFS路径删除完成", nfs.PATH),
	})
}