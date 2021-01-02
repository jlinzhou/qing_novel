package dto

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespSucess(c *gin.Context, data interface{}) {
	ret := Response{Code: 0, Msg: "", Data: data}
	c.JSON(http.StatusOK, ret)
	c.Abort()
}

func RespFail(c *gin.Context, code int, msg string, data interface{}) {

	ret := Response{Code: code, Msg: msg, Data: data}
	c.JSON(http.StatusOK, ret)
	c.Abort()
}
