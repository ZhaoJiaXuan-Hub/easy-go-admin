package controller

import (
	"easy-go-admin/config/message"
	"easy-go-admin/config/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	salt := util.GenerateRandomString(15)
	pass := util.EncryptionPassword("123456", salt)
	message.Success(c, []string{
		pass,
		salt,
	})
}
