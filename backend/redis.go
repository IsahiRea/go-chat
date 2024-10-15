package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}

// Publish a message to a Redis channel (chat room)
func publishMessage(room string, message []byte) {
	err := redisClient.Publish(ctx, room, message).Err()
	if err != nil {
		log.Println("Error publishing message:", err)
	}
}

// Subscribe to a Redis channel (chat room)
func subscribeRoom(room string) *redis.PubSub {
	return redisClient.Subscribe(ctx, room)
}
