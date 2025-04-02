package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type AuthenticationService struct {
	SecretKey string
}

func NewAuthenticationService() (*AuthenticationService, error) {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error al cargar el archivo .env")
	}

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT_SECRET no está definido en el archivo .env")
	}

	return &AuthenticationService{SecretKey: secretKey}, nil
}

func (a *AuthenticationService) GenerateToken(userID int64, username string) (string, string, error) {
	claims := jwt.MapClaims{
		"userID":   userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}

	// Generar el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", "", err
	}

	// Crear un hash SHA256 del token
	hash := sha256.New()
	hash.Write([]byte(tokenString))
	hashedToken := hex.EncodeToString(hash.Sum(nil))

	return tokenString, hashedToken, nil
}

func (a *AuthenticationService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return []byte(a.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return token, nil
}

func (a *AuthenticationService) HashToken(token string) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(token))
	return hex.EncodeToString(hash.Sum(nil)), nil
}
