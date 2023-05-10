package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

func SendRespPost(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Success: true,
		Result:  data,
	})
}

func SendRespGet(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Success: true,
		Result:  data,
	})
}

func SuccessCreate(c *gin.Context) {
	c.JSON(http.StatusCreated, ResponseData{
		Success: true,
	})
}

func Failure(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    http.StatusBadRequest,
		Msg:     err.Error(),
		Success: false,
	})
}

func QueryFailure(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    http.StatusBadRequest,
		Msg:     "query failed",
		Success: false,
		Result:  err.Error(),
	})
}
