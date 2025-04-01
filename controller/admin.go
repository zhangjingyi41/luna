package controller

import (
	"luna/pkg/api"

	"github.com/gin-gonic/gin"
)

// @Summary 登录接口
// @Description 登录接口
// @Tags 后台管理
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /admin/login [post]
func Login(c *gin.Context) {
	api.SuccessWithMsg(c, "登录成功", nil)
}

// @Summary 退出接口
// @Description 退出接口
// @Tags 后台管理
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Router /admin/logout [post]
func Logout(c *gin.Context) {
	api.SuccessWithMsg(c, "退出成功", nil)
}
