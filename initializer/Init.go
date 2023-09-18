package initializer

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func ConnectMYSQL() *gorm.DB {
	ConnectEnv()
	dsn := os.Getenv("DB_CONNECT")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//loi ket noi database
	if err != nil {
		log.Fatalf("can't connect to database, error: %s", err)
	}
	return db
}

func ConnectMongo() *mongo.Client {
	ConnectEnv()
	uri := os.Getenv("MG_CONNECT")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
