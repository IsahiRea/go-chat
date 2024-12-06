package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Generate JWT token for a user
func generateToken(username, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Parse and validate JWT token
func validateToken(tokenStr, jwtSecret string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	} else {
		return "", err
	}
}

func HashPassword(password string) (string, error) {

	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		return "", fmt.Errorf("error hashing")
	}

	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {

	bytePassword := []byte(password)
	byteHash := []byte(hash)

	if err := bcrypt.CompareHashAndPassword(byteHash, bytePassword); err != nil {
		return fmt.Errorf("error password mismatch")
	}

	return nil
}
