package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handlerWebSocket(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	username, err := validateToken(tokenStr)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	room := r.URL.Query().Get("room")
	if room == "" {
		http.Error(w, "Room not specified", http.StatusBadRequest)
		return
	}

	log.Printf("User '%s' connected to room '%s'", username, room)

	// Subscribe to the Redis channel for this room
	pubsub := subscribeRoom(room)
	defer pubsub.Close()

	// Channel for receiving messages from Redis
	redisChannel := pubsub.Channel()

	// Goroutine for sending messages from Redis to WebSocket
	go func() {
		for msg := range redisChannel {
			// Forward message to the WebSocket client
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Println("Error sending message to WebSocket:", err)
				break
			}
		}
	}()

	// WebSocket message loop
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received message: %#v", msg)

		// Include the username from the JWT in the message
		msg.User = username

		saveMessage(msg)

		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Println("Error marshaling message to JSON:", err)
			continue
		}

		publishMessage(msg.Room, jsonMsg)

	}
}

// Handle request to fetch message history for a room
func getMessageHistoryHandler(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	_, err := validateToken(tokenStr)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	room := r.URL.Query().Get("room")
	if room == "" {
		http.Error(w, "Room not specified", http.StatusBadRequest)
		return
	}

	messages, err := getMessages(room)
	if err != nil {
		log.Println("Error fetching messages:", err)
		http.Error(w, "Failed to retrieve message history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
