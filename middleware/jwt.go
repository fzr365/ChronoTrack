package middleware

import (
    "ChronoTrack/models"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "os"
)

func JWTProtected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        tokenString := c.Get("Authorization")
        if tokenString == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
        }

        token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })
        if err != nil || !token.Valid {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
        }

        claims := token.Claims.(*jwt.StandardClaims)
        user := models.User{ID: uint(claims.Subject)}
        c.Locals("user", &user)

        return c.Next()
    }
}