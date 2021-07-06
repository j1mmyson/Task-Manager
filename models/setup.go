package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("User"), os.Getenv("Password"), os.Getenv("Host"), os.Getenv("DBNAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect db")
	}

	db.AutoMigrate(&List{}, &User{}, &Session{})

	DB = db
}
