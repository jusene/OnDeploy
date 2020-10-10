package controllers

import (
	"OnDeploy/models"
	"OnDeploy/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary RabbitMQ服务安装
// @Description RabbitMQ服务安装
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Security basic
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/install [post]
func RabbitMQInstall(ctx *gin.Context) {
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

	if err := utils.InstallRabbitMQ(server, user, pass); err != nil {
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

// @Summary RabbitMQ虚拟机添加
// @Description RabbitMQ虚拟机添加
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param vhost body models.NewRabbitVhost true "vhost"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/vhost/add [put]
func RabbitMQVhostAdd(ctx *gin.Context) {
	var vhost models.NewRabbitVhost
	if err := ctx.ShouldBindJSON(&vhost);err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(vhost.Address, "guest", "guest", 15672)
	if err := rabbitapi.AddVhost(vhost.Vhost, vhost.Trace); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Err{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 虚拟机添加成功", vhost.Vhost),
	})


}

// @Summary RabbitMQ虚拟机删除
// @Description RabbitMQ虚拟机删除
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param vhost body models.NewRabbitVhost true "vhost"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/vhost/del [delete]
func RabbitMQVhostDel(ctx *gin.Context) {
	var vhost models.NewRabbitVhost
	if err := ctx.ShouldBindJSON(&vhost);err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(vhost.Address, "guest", "guest", 15672)
	if err := rabbitapi.DelVhost(vhost.Vhost); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Err{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 虚拟机删除成功", vhost.Vhost),
	})
}

// @Summary RabbitMQ虚拟机列表
// @Description RabbitMQ虚拟机列表
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param address path string true "address"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/vhost/lst/{address} [get]
func RabbitMQVhostLst(ctx *gin.Context) {
	address := ctx.Param("address")
	rabbitapi := utils.NewRabbitApi(address, "guest", "guest", 15672)
	ret, err := rabbitapi.LstVhost()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	var vhosts models.RabbitVhosts
	err = json.Unmarshal([]byte(ret), &vhosts)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, vhosts)
}

// @Summary RabbitMQ用户添加
// @Description RabbitMQ用户添加
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param user body models.NewRabbitUser true "user"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/user/add [put]
func RabbitMQUserAdd(ctx *gin.Context) {
	var user models.NewRabbitUser
	// 检查请求json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(user.Address, "guest", "guest", 15672)
	err := rabbitapi.AddUser(user.User, user.Pass)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Err{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 用户添加成功", user.User),
	})
}

// @Summary RabbitMQ用户删除
// @Description RabbitMQ用户删除
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param user body models.CommonRabbitUser true "user"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/user/del [delete]
func RabbitMQUserDel(ctx *gin.Context) {
	var user models.CommonRabbitUser
	// 检查请求json
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(user.Address, "guest", "guest", 15672)
	err := rabbitapi.DelUser(user.User)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Err{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 用户删除成功", user.User),
	})
}

// @Summary RabbitMQ用户列表
// @Description RabbitMQ用户列表
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param address path string true "address"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/user/lst/{address} [get]
func RabbitMQUserLst(ctx *gin.Context) {
	address := ctx.Param("address")
	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(address, "guest", "guest", 15672)
	ret, err := rabbitapi.LstUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	var users models.RabbitUsers
	err = json.Unmarshal([]byte(ret), &users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// @Summary RabbitMQ权限列表
// @Description RabbitMQ权限列表
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param address path string true "address"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/permission/lst/{address} [get]
func RabbitMQPermissionLst(ctx *gin.Context) {
	address := ctx.Param("address")
	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(address, "guest", "guest", 15672)
	ret, err := rabbitapi.LstPermission()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	var permissions models.RabbitPermissions
	err = json.Unmarshal([]byte(ret), &permissions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, permissions)
}

// @Summary RabbitMQ权限添加
// @Description RabbitMQ权限添加
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param permission body models.NewRabbitPermission true "permission"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/permission/add [put]
func RabbitMQPermissionAdd(ctx *gin.Context) {
	var permission models.NewRabbitPermission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(permission.Address, "guest", "guest", 15672)
	if err := rabbitapi.AddPermission(&permission); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Err{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 用户权限添加成功", permission.User),
	})
}

// @Summary RabbitMQ权限删除
// @Description RabbitMQ权限删除
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param permission body models.ComRabbitPermission true "permission"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/permission/del [delete]
func RabbitMQPermissionDel(ctx *gin.Context) {
	var permission models.ComRabbitPermission
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	var rabbitapi utils.Rabbit
	rabbitapi = utils.NewRabbitApi(permission.Address, "guest", "guest", 15672)
	if err := rabbitapi.DelPermission(&permission); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Err{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Err{
		Code: http.StatusOK,
		Message: fmt.Sprintf("%s 用户权限删除成功", permission.User),
	})
}