package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 健康检查接口
// @Description 健康检查接口
// @Tags 公共接口
// @Accept json
// @Produce json
// @Success 200 {string} string "ok"
// @Router /health [get]
func Health(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
