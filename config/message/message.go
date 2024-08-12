package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功返回消息
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	response := Response{}
	response.Code = int(Code.SUCCESS)
	response.Message = Code.GetMessage(Code.SUCCESS)
	response.Data = data
	c.JSON(http.StatusOK, response)
}

// Fail 失败返回消息
func Fail(c *gin.Context, code int, message string) {
	response := Response{}
	response.Code = code
	response.Message = message
	c.JSON(http.StatusOK, response)
}
