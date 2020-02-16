package jwt

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

// Client ...
type Client struct {
	in ClientInput
}

// ClientInput ...
type ClientInput struct {
	Secret string
	TTL    int64
}

// NewClient ...
func NewClient(in ClientInput) *Client {
	if in.Secret == "" {
		log.Panic("JWT secret is required")
	}

	if in.TTL == 0 {
		log.Panic("TTL is required")
	}

	return &Client{
		in: in,
	}
}

// GenerateJWT ...
func (c Client) GenerateJWT(userID string) (string, error) {
	jwtInstance := jwt.New(jwt.SigningMethodHS256)
	claims := jwtInstance.Claims.(jwt.MapClaims)

	claims["userID"] = userID
	claims["exp"] = c.in.TTL

	token, tokenErr := jwtInstance.SignedString([]byte(c.in.Secret))
	if tokenErr != nil {
		fmt.Println(tokenErr)
		return "", fmt.Errorf("Failed to signin the JWT with the secret")
	}

	return token, nil
}
