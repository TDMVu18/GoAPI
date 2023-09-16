package initializer

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectEnv() {
	err := godotenv.Load(".env")
	//loi ket noi moi truong
	if err != nil {
		log.Fatalf("can't connect to environment file; error: %s", err)
	}
}

func ConnectDatabase() *gorm.DB {
	ConnectEnv()
	dsn := os.Getenv("DB_CONNECT")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//loi ket noi database
	if err != nil {
		log.Fatalf("can't connect to database, error: %s", err)
	}
	return db
}
