package config

import (
	"auth_api/internal/config/database"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database database.GoOrmDatabase
)

func InitializeDatabase() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: os.Getenv("DSN"),
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&database.User{})
	db.AutoMigrate(&database.Session{})
	Database = database.GoOrmDatabase{DB: db}
}
