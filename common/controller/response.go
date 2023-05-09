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

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    200,
		Msg:     "",
		Success: true,
		Result:  data,
	})
}

func SuccessCreate(c *gin.Context) {
	c.JSON(http.StatusCreated, ResponseData{
		Code:    201,
		Msg:     "",
		Success: true,
		Result:  "",
	})
}

func Failure(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    http.StatusBadRequest,
		Msg:     err.Error(),
		Success: false,
		Result:  "",
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
