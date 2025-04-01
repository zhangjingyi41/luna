package router

import (
	"luna/controller"
	_ "luna/docs"

	"github.com/gin-gonic/gin"
)

// 后台路由
func AdminRouter(server *gin.Engine) {
	admin := server.Group("/admin")

	admin.POST("/login", controller.Login)

	admin.POST("/logout", controller.Logout)
}
