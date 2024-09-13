package authService

import (
	"errors"
	"time"
	"unified-atlassian-platform/db"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

// Credentials represents the login credentials.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims represents the JWT claims.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login authenticates a user and returns a JWT token.
func Login(creds Credentials) (string, error) {
	var storedPassword string
	query := "SELECT password FROM users WHERE username = ?"
	err := db.DB.QueryRow(query, creds.Username).Scan(&storedPassword)
	if err != nil || storedPassword != creds.Password {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT Token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
