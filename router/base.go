package router

import (
	"luna/controller"
	"luna/logger"

	_ "luna/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func RunServer(mode string) *gin.Engine {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.New()
	// 跨域配置
	corsConfig := cors.DefaultConfig()
	// 加载中间件
	// 中间件的执行顺序是按照注册顺序排列
	server.Use(
		logger.GinLogger(),
		logger.GinRecovery(true),
		cors.New(corsConfig),
	)
	// 暂时不管静态文件和html文件

	// swagger路由
	server.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))

	// 健康检查
	server.GET("/health", controller.Health)
	// 加载其他模块的路由
	AdminRouter(server)

	return server
}
