package config

import (
	"backend_relawanku/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Admin{})
}