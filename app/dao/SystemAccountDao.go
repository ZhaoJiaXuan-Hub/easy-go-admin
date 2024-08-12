package dao

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/config/database"
)

func SystemAccountDetail(username string) (systemAccount entity.SystemAccount) {
	database.Db.Where("username = ?", username).First(&systemAccount)
	return systemAccount
}

func SystemAccountAdd(account entity.SystemAccount) (err error) {
	if err = database.Db.Create(&account).Error; err != nil {
		return err
	}
	return nil
}
