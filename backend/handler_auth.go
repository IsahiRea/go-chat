package main

import (
	"encoding/json"
	"net/http"
)

// Dummy user authentication for simplicity
// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

// LoginHandler - Authenticate the user and return a JWT token
func (cfg *ApiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//storedPassword, exists := users[loginReq.Username]

	// Check if the user exists and password matches
	user, exists := getUser(loginReq.Username)
	if !exists {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err = CheckPasswordHash(loginReq.Password, user.Password)
	if err != nil {
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
	json.NewEncoder(w).Encode(Token{
		Token: token,
	})
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

	w.WriteHeader(http.StatusOK)
}
