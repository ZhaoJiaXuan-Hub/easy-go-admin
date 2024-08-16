package dao

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/config/database"
)

// SystemDictQueryAllData 查询所有字典及其数据
func SystemDictQueryAllData() ([]entity.SystemDict, error) {
	var dict []entity.SystemDict
	if err := database.Db.Preload("DictData").Find(&dict).Error; err != nil {
		return nil, err
	}
	return dict, nil
}
