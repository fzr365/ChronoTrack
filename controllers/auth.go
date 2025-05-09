package controllers

import (
    "ChronoTrack/database"
    "ChronoTrack/models"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "os"
)

func Register(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

    user := models.User{
        Username: data["username"],
        Password: string(password),
    }

    if err := database.DB.Create(&user).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
    }

    return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    var user models.User
    if err := database.DB.Where("username = ?", data["username"]).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    claims := jwt.StandardClaims{
        Subject:   string(rune(user.ID)),
        ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
    }

    return c.JSON(fiber.Map{"token": tokenString})
}