package database

import (
	"bankku/config"
	usermodel "bankku/domains/user/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MYSQL_USER,
		cfg.MYSQL_PASSWORD,
		cfg.MYSQL_HOST,
		cfg.MYSQL_PORT,
		cfg.MYSQL_DBNAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&usermodel.User{},
	)
}
