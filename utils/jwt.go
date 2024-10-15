package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GetENV() string{
    err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    
    apisecret := os.Getenv("API_SECRET")

    return apisecret
}

type Claims struct {
    UserID uint `json:"user_id"`
    jwt.RegisteredClaims
}

var jwtKey = []byte(GetENV())

func GenerateToken(userID uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        // Debugging: Cetak error dari jwt.ParseWithClaims
        fmt.Printf("Token parsing error: %v\n", err)
        return nil, err
    }
    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    return claims, nil
}

