package main

import (
	"context"
	"errors"
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

	defer func() {
		if err := client.Disconnect(mongoCtx); err != nil {
			panic(err)
		}
	}()

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

// TODO: Test the Implementation
func getUser(username string) (User, bool) {

	var result User
	filter := bson.M{"username": username}

	err := userCollection.FindOne(mongoCtx, filter).Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return User{}, false
	} else if err != nil {
		return User{}, false
	}

	return result, true

}
