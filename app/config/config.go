package config

import (
	"mvc-golang-2/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBinit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:12345@/digitalent_bank?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database" + err.Error())
	}
	db.AutoMigrate(new(model.Account), new(model.Transaction))

	DB = db

	return db

}
