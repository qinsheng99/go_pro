package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespOption func(m map[string]interface{})

func Success(c *gin.Context, code int, data interface{}, options ...RespOption) {
	c.JSON(code, successReturn(data, options...))
}

func Failure(c *gin.Context, err error) {
	c.JSON(http.StatusOK, handleBadReturn(err))
}

func QueryFailure(c *gin.Context) {
	c.JSON(http.StatusOK, queryHandleBadReturn(nil))
}

func successReturn(data interface{}, options ...RespOption) map[string]interface{} {
	var info = make(map[string]interface{})
	info["code"] = 0
	info["msg"] = ""
	info["success"] = true
	info["result"] = data
	for _, option := range options {
		option(info)
	}
	return info
}
func handleBadReturn(err error) map[string]interface{} {
	var info = make(map[string]interface{})
	info["code"] = 1
	info["msg"] = ""
	info["success"] = false
	info["result"] = err.Error()
	return info
}

func queryHandleBadReturn(data interface{}) map[string]interface{} {
	var info = make(map[string]interface{})
	info["code"] = 1
	info["msg"] = "数据参数不正确"
	info["success"] = false
	info["result"] = data
	return info
}
