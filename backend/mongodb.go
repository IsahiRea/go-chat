package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var userCollection *mongo.Collection
var messageCollection *mongo.Collection
var mongoCtx = context.Background()

func initMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(mongoCtx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	mongoClient = client
	userCollection = mongoClient.Database("chat").Collection("users")
	messageCollection = mongoClient.Database("chat").Collection("messages")
}

// Save a message to MongoDB
func saveMessage(message Message) {
	_, err := messageCollection.InsertOne(mongoCtx, message)
	if err != nil {
		log.Println("Error saving message to MongoDB:", err)
	}
}

// Retrieve messages from MongoDB
func getMessages(room string) ([]Message, error) {
	var messages []Message
	cursor, err := messageCollection.Find(mongoCtx, bson.M{"room": room})
	if err != nil {
		return nil, err
	}
	err = cursor.All(mongoCtx, &messages)
	return messages, err
}

// Save users to MongoDB
func saveUser(user User) {
	_, err := userCollection.InsertOne(mongoCtx, user)
	if err != nil {
		log.Println("Error saving User to MongoDB:", err)
	}
}
