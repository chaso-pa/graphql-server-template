package middleware

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() *gorm.DB {
	//データベースへの接続
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
