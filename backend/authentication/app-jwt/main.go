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
	tokenToVerify := os.Getenv("TOKEN_TO_VERIFY")

	token, err := createJWTToken(username, password)
	if err != nil {
		log.Fatalf("error to create token: %v", err)
	}

	log.Println("token created!")
	log.Println(token)

	valid, err := verifyToken(tokenToVerify)
	if err != nil {
		log.Fatalf("error to verify token: %v", err)
	}

	if valid {
		log.Println("token is valid!")
	} else {
		log.Println("token is not valid")
	}

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

func verifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}

	claims := token.Claims
	expDate, err := claims.GetExpirationTime()
	if err != nil {
		return false, err
	}

	log.Printf("exp: %s", expDate)

	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return false, err
	}

	log.Printf("iat: %s", issuedAt)

	fmt.Printf("claims: %v\n", claims)

	return token.Valid, nil
}
