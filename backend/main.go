package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize Redis and MongoDB
	initRedis()
	initMongoDB()

	// WebSocket route for real-time chat
	http.HandleFunc("/ws", handlerWebSocket)

	// HTTP routes for login and message history
	http.HandleFunc("/register", handlerRegister)
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/messages", getMessageHistoryHandler)
	http.HandleFunc("/ready", handlerReadiness)

	// Start the server
	log.Println("Server started at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
