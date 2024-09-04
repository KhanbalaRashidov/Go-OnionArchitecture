package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&User{}, &Post{}, &Comment{}, &Tag{})
}
