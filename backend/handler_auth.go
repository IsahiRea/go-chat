package main

import (
	"encoding/json"
	"net/http"
)

// Dummy user authentication for simplicity
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Request body for login
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler - Authenticate the user and return a JWT token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//TODO: Create database for User Management

	// Check if the user exists and password matches
	storedPassword, exists := users[loginReq.Username]
	if !exists || storedPassword != loginReq.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token
	token, err := generateToken(loginReq.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Send the token as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
