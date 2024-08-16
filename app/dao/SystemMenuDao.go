package dao

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/config/database"
	"easy-go-admin/config/util"
)

// SystemMenuAll 通过entity.SystemMenu获取所有菜单
func SystemMenuAll() (systemMenus []entity.SystemMenu, err error) {
	err = database.Db.Model(&entity.SystemMenu{}).Preload("RolesInfo").Find(&systemMenus).Error
	return systemMenus, err
}

// SystemMenuQuery 通过条件进行筛选获取菜单
func SystemMenuQuery(wheres interface{}, columns interface{}, orderBy interface{}, page, rows int, total *int64) (systemMenus []*entity.SystemMenu, err error) {
	db := database.Db

	var model []*entity.SystemMenu
	var mod entity.SystemMenu
	db, err = util.BuildQueryList(db, wheres, columns, orderBy, page, rows)
	if err != nil {
		return nil, err
	}
	err = db.Preload("RolesInfo").Find(&model).Error
	db, err = util.BuildWhere(db, wheres)

	db = database.Db
	db.Model(&mod).Count(total)

	return model, nil
}

// SystemMenuCreate 新增菜单
func SystemMenuCreate(menu entity.SystemMenu) (err error) {
	err = database.Db.Create(&menu).Error
	return err
}

// SystemMenuUpdate 更新菜单
func SystemMenuUpdate(menu entity.SystemMenu) (err error) {
	err = database.Db.Model(&menu).Updates(menu).Error
	return err
}

// SystemMenuDelete 删除菜单
func SystemMenuDelete(menu entity.SystemMenu) (err error) {
	err = database.Db.Delete(&menu).Error
	return err
}

// SystemMenuDeleteByIds 批量删除菜单
func SystemMenuDeleteByIds(ids []uint) (err error) {
	err = database.Db.Delete(&entity.SystemMenu{}, "id in (?)", ids).Error
	return err
}

// SystemMenuDetail 获取菜单详情
func SystemMenuDetail(id uint) (menu entity.SystemMenu, err error) {
	err = database.Db.Model(&entity.SystemMenu{}).Where("id = ?", id).First(&menu).Error
	return menu, err
}
