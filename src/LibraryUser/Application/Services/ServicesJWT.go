package services


import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v4"

	"log"
)

// Claims representa las reclamaciones que se almacenan en el JWT.
type Claims struct {
	Username string `json:"username"`
	Folio    int    `json:"folio"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT genera un JWT para el usuario de la biblioteca.
func GenerateJWT(username, role string, folio int) (string, error) {
	claims := Claims{
		Username: username,
		Role:     role,
		Folio:    folio,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // El token expira en 24 horas
			Issuer:    "LibraryApp", // El emisor del token
		},
	}

	// Usar la clave secreta del archivo .env
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error al firmar el token:", err)
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT valida un token JWT y extrae las reclamaciones.
func ValidateJWT(tokenStr string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
