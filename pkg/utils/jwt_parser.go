package utils

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID      uuid.UUID
	Role        string
	Credentials map[string]bool
	Expires     int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		role := string(claims["role"].(string))
		// Expires time.
		expires := int64(claims["expires"].(float64))

		// User credentials.
		credentials := map[string]bool{
			"User:create_user":    claims["User:create_user"].(bool),
			"User:create_banker":  claims["User:create_banker"].(bool),
			"User:update_user":    claims["User:update_user"].(bool),
			"User:update_banker":  claims["User:update_banker"].(bool),
			"User:disable_user":   claims["User:disable_user"].(bool),
			"User:Disable_banker": claims["User:create_user"].(bool),
		}

		return &TokenMetadata{
			UserID:      userID,
			Role:        role,
			Credentials: credentials,
			Expires:     expires,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
