package model

import (
	"GoAPI/initializer"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var mongoClient *mongo.Client
var collection *mongo.Collection

type Person struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name" form:"name"`
	Major     string             `json:"major" bson:"major" form:"major"`
	CreatedAt *time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time         `json:"updated_at" bson:"updated_at"`
	Deleted   bool               `json:"deleted" bson:"deleted"`
}

func ModelList(search string) []bson.M {
	collection := initializer.ConnectDB()
	defer initializer.DisconnectDB()
	filter := bson.M{}
	if search != "" {
		filter["$or"] = []bson.M{
			bson.M{"name": bson.M{"$regex": search, "$options": "i"}},
			bson.M{"major": bson.M{"$regex": search, "$options": "i"}},
		}
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results
}

func ModelGet(id string) *Person {
	collection := initializer.ConnectDB()
	defer initializer.DisconnectDB()
	var person Person
	personId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": personId}
	err := collection.FindOne(context.TODO(), filter).Decode(&person)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}
	return &person
}

func ModelCreate(person Person) string {
	collection := initializer.ConnectDB()
	defer initializer.DisconnectDB()
	_, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		panic(err)
	}
	return "Created successfully"
}

func ModelDelete(id string) string {
	collection := initializer.ConnectDB()
	defer initializer.DisconnectDB()
	personId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": personId}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return "Deleted successfully"
}

func ModelUpdate(person Person) string {
	collection := initializer.ConnectDB()
	defer initializer.DisconnectDB()
	update := bson.M{"$set": bson.M{"name": person.Name, "major": person.Major}}
	filter := bson.M{"_id": person.ID}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return "Updated successfully"
}
