package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/model"
)

var (
	db *gorm.DB
)

func Dbfrom() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/data?charset=utf8mb4&parseTime=True&loc=Local" //数据库登入
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{}, &model.Group{}, &model.Message{}, &model.Apply{}, &model.Session{})
	return db
}
