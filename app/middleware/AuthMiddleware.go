package middleware

import (
	"easy-go-admin/app/constant"
	"easy-go-admin/app/dao"
	"easy-go-admin/app/entity"
	"easy-go-admin/config/jwt"
	"easy-go-admin/config/message"
	"fmt"
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
		parts := strings.SplitN(header, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			message.Fail(c, int(message.Code.AUTHFORMMATER), message.Code.GetMessage(message.Code.AUTHFORMMATER))
			c.Abort()
			return
		}
		// 验证token
		tokenStr := parts[1]
		userData, err := jwt.ValidateToken(tokenStr)
		if err != nil {
			message.Fail(c, int(message.Code.INVALIDTOKEN), err.Error())
			c.Abort()
			return
		}
		// 根据用户 ID 查询用户
		user, err := dao.SystemAccountDetailById(userData.Id)
		if err != nil {
			fmt.Println(err.Error())
			message.Fail(c, int(message.Code.NOTFOUND), "登录用户已经不存在")
			c.Abort()
			return
		}
		// 检查用户状态
		if user.Status != 1 { // 1 表示正常状态
			message.Fail(c, int(message.Code.STATUSISENABLE), message.Code.GetMessage(message.Code.STATUSISENABLE))
			c.Abort()
			return
		}
		userObj := entity.JwtSystemAccountData{
			UserName: user.UserName,
			NickName: user.NickName,
			Avatar:   user.Avatar,
			Email:    user.Email,
			Roles:    user.Roles,
			Dept:     user.Dept,
		}
		c.Set(constant.ContextKeyUserObj, &userObj)
		c.Next()
	}
}
