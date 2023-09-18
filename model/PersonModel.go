package model

import (
	"GoAPI/initializer"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var mongoClient *mongo.Client
var collection *mongo.Collection

type Person struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Major     string             `json:"major" bson:"major"`
	CreatedAt *time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at"`
	Deleted   bool               `json:"deleted" bson:"deleted"`
}

func ConnectDB() *mongo.Collection {
	mongoClient = initializer.ConnectMongo()
	collection = mongoClient.Database("personlist").Collection("person_list")
	return collection

}

func DisconnectDB() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

type ManageDB interface {
	Insert(interface{}) error
	GetAll() ([]Person, error)
	DeleteData(primitive.ObjectID) error
	UpdateData(Person) error
}
