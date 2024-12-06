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
func (cfg *ApiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//TODO: Change check for future database implementation
	// Check if the user exists and password matches
	storedPassword, exists := users[loginReq.Username]
	if !exists || storedPassword != loginReq.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token
	token, err := generateToken(loginReq.Username, cfg.jwtSecret)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Send the token as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (cfg *ApiConfig) handlerRegister(w http.ResponseWriter, r *http.Request) {

	params := RegisterParams{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		return
	}

	hashed, err := HashPassword(params.Password)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user := User{
		Username: params.Username,
		Password: hashed,
	}

	saveUser(user)

	//FIXME: JSON
	//RespondwithJSON()
	w.WriteHeader(200)
}
