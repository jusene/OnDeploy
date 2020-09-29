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

// @Summary RabbitMQ虚拟机添加
// @Description RabbitMQ虚拟机添加
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param user body models.NewRabbitUser true "user"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/vhost/add [put]
func RabbitMQVhostAdd(ctx *gin.Context) {


}

// @Summary RabbitMQ虚拟机删除
// @Description RabbitMQ虚拟机删除
// @Tags RabbitMQ服务
// @Accept json
// @Produce json
// @Param server body models.ServerDetail true "server"
// @Success 200 {object} models.Res
// @Failure 400 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /app/rabbitmq/vhost/del [delete]
func RabbitMQVhostDel(ctx *gin.Context) {

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