package controllerSystem

import (
	"easy-go-admin/app/service"
	"github.com/gin-gonic/gin"
)

// DictData 获取所有字典及数据
// @Summary 获取所有字典及数据
// @Produce json
// @Description 获取所有字典及数据
// @Success 200 {object} message.Response
// @router /api/system/dict/dictData [get]
func DictData(c *gin.Context) {
	service.SystemDictService().GetDictData(c)
}
