package middleware

import (
	"easy-go-admin/app/constant"
	"easy-go-admin/config/message"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthMiddleware 权限验证中间件
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" {
			message.Fail(c, int(message.Code.NOAUTH), message.Code.GetMessage(message.Code.NOAUTH))
			c.Abort()
			return
		}
		parts := strings.SplitN(header, "", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			message.Fail(c, int(message.Code.AUTHFORMMATER), message.Code.GetMessage(message.Code.AUTHFORMMATER))
			c.Abort()
			return
		}
		// todo 检验token
		var token = "token"
		c.Set(constant.ContextKeyUserObj, token)
		c.Next()
	}
}
