package database

import (
	"easy-go-admin/app/entity"
	"easy-go-admin/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

// InitDB 初始化数据库
func InitDB() error {
	var err error
	var dbConfig = config.Config.Database
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset,
	)
	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}
	sqlDB, err := Db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
	// 执行迁移
	err = Db.AutoMigrate(&entity.SystemRole{}, &entity.SystemAccount{}, &entity.SystemDictData{}, &entity.SystemDict{}, &entity.SystemMenu{}, &entity.SystemDept{})
	if err != nil {
		panic(err)
	}
	return nil
}
