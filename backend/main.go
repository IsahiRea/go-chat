package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	jwtSecret string
}

func main() {

	godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")

	apiConfig := ApiConfig{
		jwtSecret: jwtSecret,
	}

	// Initialize Redis and MongoDB
	initRedis()
	initMongoDB()

	// WebSocket route for real-time chat
	http.HandleFunc("/ws", apiConfig.handlerWebSocket)

	// HTTP routes for login and message history
	http.HandleFunc("/register", apiConfig.handlerRegister)
	http.HandleFunc("/login", apiConfig.handlerLogin)
	http.HandleFunc("/messages", apiConfig.getMessageHistoryHandler)
	http.HandleFunc("/ready", handlerReadiness)

	// Start the server
	log.Println("Server started at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
