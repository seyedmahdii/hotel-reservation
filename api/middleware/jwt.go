package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("-- JWTAuthentication --")

	token := c.Get("X-Api-Token")
	if token == "" {
		return fmt.Errorf("unauthorized")
	}

	if err := parseToken(token); err != nil {
		return err
	}

	return nil
}

func parseToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse token", err)
		return fmt.Errorf("unauthorized")

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims)
	}

	return fmt.Errorf("unauthorized")
}
