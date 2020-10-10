package controllers

import (
	"OnDeploy/models"
	"OnDeploy/templates"
	"OnDeploy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Git服务安装
// @Description Git服务安装
// @Tags Git服务
// @Accept json
// @Produce json
// @Security basic
// @Param server body models.GITDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/git/install [post]
func GitInstall(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var server models.GITDetail
	// 检查请求json
	if err := ctx.ShouldBindJSON(&server); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	gitConf, err := utils.RendTmp(templates.GITConf, server)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if err := utils.InstallGit(server, user, pass, gitConf); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%s git服务安装成功", server.Address),
	})
}

// @Summary Git仓库创建
// @Description Git仓库创建
// @Tags Git服务
// @Accept json
// @Produce json
// @Security basic
// @Param repo body models.GITRepo true "repo"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/git/repo/add [put]
func GitRepoAdd(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var repo models.GITRepo
	if err := ctx.ShouldBindJSON(&repo); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var git utils.GIT
	gitobj := new(utils.GITObj)
	gitobj.Server = models.GITDetail{
		ServerDetail: models.ServerDetail{Address: repo.Address, Port: repo.Port},
	}
	gitobj.Pass = pass
	gitobj.User = user
	git = gitobj
	if err := git.AddRepo(repo.Repo); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 仓库创建成功", repo.Repo),
	})
}

// @Summary Git仓库删除
// @Description Git仓库删除
// @Tags Git服务
// @Accept json
// @Produce json
// @Security basic
// @Param repo body models.GITRepo true "repo"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/git/repo/del [delete]
func GitRepoDel(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var repo models.GITRepo
	if err := ctx.ShouldBindJSON(&repo); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var git utils.GIT
	gitobj := new(utils.GITObj)
	gitobj.Server = models.GITDetail{
		ServerDetail: models.ServerDetail{Address: repo.Address, Port: repo.Port},
	}
	gitobj.Pass = pass
	gitobj.User = user
	git = gitobj
	if err := git.DelRepo(repo.Repo); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 仓库删除成功", repo.Repo),
	})
}


// @Summary Git用户添加
// @Description Git用户添加
// @Tags Git服务
// @Accept json
// @Produce json
// @Security basic
// @Param user body models.GITUserPass true "user"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/git/user/add [post]
func GitUserAdd(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var userPass models.GITUserPass
	if err := ctx.ShouldBindJSON(&userPass); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var git utils.GIT
	gitobj := new(utils.GITObj)
	gitobj.Server = models.GITDetail{
		ServerDetail: models.ServerDetail{Address: userPass.Address, Port: userPass.Port},
	}
	gitobj.Pass = pass
	gitobj.User = user
	git = gitobj

	if err := git.AddUser(userPass.User, userPass.Pass); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 用户添加成功", userPass.User),
	})
}

// @Summary Git用户密码修改
// @Description Git用户密码修改
// @Tags Git服务
// @Accept json
// @Produce json
// @Security basic
// @Param user body models.GITUserPass true "user"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/git/user/update [put]
func GitUserUpdate(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var userPass models.GITUserPass
	if err := ctx.ShouldBindJSON(&userPass); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var git utils.GIT
	gitobj := new(utils.GITObj)
	gitobj.Server = models.GITDetail{
		ServerDetail: models.ServerDetail{Address: userPass.Address, Port: userPass.Port},
	}
	gitobj.Pass = pass
	gitobj.User = user
	git = gitobj

	if err := git.UpdateUser(userPass.User, userPass.Pass); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 用户密码修改成功", userPass.User),
	})

}

// @Summary Git用户删除
// @Description Git用户删除
// @Tags Git服务
// @Accept json
// @Produce json
// @Security basic
// @Param user body models.GITUser true "user"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/git/user/del [delete]
func GitUserDel(ctx *gin.Context) {
	user, pass, err := utils.AuthRequired(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	var userPass models.GITUser
	if err := ctx.ShouldBindJSON(&userPass); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var git utils.GIT
	gitobj := new(utils.GITObj)
	gitobj.Server = models.GITDetail{
		ServerDetail: models.ServerDetail{Address: userPass.Address, Port: userPass.Port},
	}
	gitobj.Pass = pass
	gitobj.User = user
	git = gitobj

	if err := git.DelUser(userPass.User); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Res{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 用户删除成功", userPass.User),
	})
}