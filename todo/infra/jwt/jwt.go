package jwt

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucasmls/todo/infra"
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

// ValidateJWT ...
func (c Client) ValidateJWT(tokenToValidate string) (*infra.DecodedJWT, error) {
	token, parseErr := jwt.Parse(tokenToValidate, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(c.in.Secret), nil
	})

	if parseErr != nil {
		return nil, fmt.Errorf("Failed to parse the JWT")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	decodedJWT := infra.DecodedJWT{
		UserID: claims["userID"].(string),
		Exp:    int64(claims["exp"].(float64)),
	}

	return &decodedJWT, nil
}
