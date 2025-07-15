package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey = os.Getenv("SECRET_KEY")

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	token, err := createJWTToken(username, password)
	if err != nil {
		log.Fatalf("error to create token: %v", err)
	}

	log.Println("token created!")
	log.Println(token)
}

func createJWTToken(username, password string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, //header
		jwt.MapClaims{ //payload
			"username": username,
			"password": password,
			"iat":      time.Now().Unix(),
			"exp":      time.Now().Add(time.Hour).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}
