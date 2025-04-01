package api

import (
	"luna/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, models.Response{
		Code: code,
		Msg:  msg,
	})
}

func ErrorWithData(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, models.Response{
		Code: code,
		Msg:  msg,
	})
}
