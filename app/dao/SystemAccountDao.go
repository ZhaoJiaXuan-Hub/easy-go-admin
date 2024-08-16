package dao

import (
	"easy-go-admin/app/constant"
	"easy-go-admin/app/entity"
	"easy-go-admin/config"
	"easy-go-admin/config/database"
)

// SystemAccountDetailByUsername 通过用户名获取用户详细
func SystemAccountDetailByUsername(username string) (systemAccount entity.SystemAccount, err error) {
	if err := database.Db.Where("user_name = ?", username).First(&systemAccount).Error; err != nil {
		return systemAccount, err // 或者根据具体情况返回错误或默认值
	}
	return systemAccount, nil
}

// SystemAccountDetailById 通过ID获取用户详细
func SystemAccountDetailById(id uint) (systemAccount entity.SystemAccount, err error) {
	if err := database.Db.Preload("Roles.Menus.RolesInfo").Preload("Dept").Where("id = ?", id).First(&systemAccount).Error; err != nil {
		return systemAccount, err // 或者根据具体情况返回错误或默认值
	}
	// 检查并设置默认头像
	if systemAccount.Avatar == "" {
		systemAccount.Avatar = config.Config.FileSettings.Host + constant.DefaultAvatarUrl // 设置默认头像 URL
	}
	return systemAccount, nil
}

// SystemAccountCreate 新增用户
func SystemAccountCreate(account entity.SystemAccount) (err error) {
	err = database.Db.Create(&account).Error
	return err
}
