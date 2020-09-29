package main

import (
	"OnDeploy/controllers"
	_ "OnDeploy/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title 应用交付部署服务
// @Version 1.0
// @Description 应用交付部署服务API
// @Schemes http https
// @BasePath /api/v1
func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(Cors())
	v1 := router.Group("/api/v1")
	{
		// 服务器初始化
		v1.POST("/server/init", controllers.ServerInit)
		v1.POST("/servers/init", controllers.ServersInit)

		// 应用部署
		v1.POST("/app/rabbitmq/install", controllers.RabbitMQInstall)
		v1.PUT("/app/rabbitmq/vhost/add", controllers.RabbitMQVhostAdd)
		v1.DELETE("/app/rabbitmq/vhost/del", controllers.RabbitMQVhostDel)
		v1.GET("/app/rabbitmq/vhost/lst/:address", controllers.RabbitMQVhostLst)
		v1.PUT("/app/rabbitmq/user/add", controllers.RabbitMQUserAdd)
		v1.DELETE("/app/rabbitmq/user/del", controllers.RabbitMQUserDel)
		v1.GET("/app/rabbitmq/user/lst/:address", controllers.RabbitMQUserLst)
		// v1.PUT("/app/rabbitmq/perssion/add", controllers.RabbitMQPerssionAdd)
		// v1.DELETE("/app/rabbitmq/perssion/del", controllers.RabbitMQPerssionDel)
		// v1.GET("/app/rabbitmq/perssion/lst", controllers.RabbitMQPerssionLst)


		v1.POST("/app/git/install", controllers.GitInstall)
		v1.POST("/app/nfs/install", controllers.NFSInstall)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":9000")
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Server", "GIN-GO")
		context.Next()
	}
}