package service

import (
	"easy-go-admin/app/dao"
	"easy-go-admin/app/entity"
	"easy-go-admin/config/message"
	"github.com/gin-gonic/gin"
)

type ISystemDictService interface {
	GetDictData(c *gin.Context)
}

type SystemDictServiceImpl struct{}

func SystemDictService() ISystemDictService {
	return &SystemDictServiceImpl{}
}

// GetDictData 查询所有字典以及数据
func (s *SystemDictServiceImpl) GetDictData(c *gin.Context) {
	dictData, err := dao.SystemDictQueryAllData()
	if err != nil {
		message.Fail(c, int(message.Code.FAILURE), err.Error())
		return
	}
	// 提取字典
	uniqueDict := make(map[string][]entity.SystemDictData)
	for _, dict := range dictData {
		uniqueDict[dict.Code] = dict.DictData
	}
	message.Success(c, uniqueDict)
}
