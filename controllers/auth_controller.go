package controllers

import (
    "restapi/models"
    "restapi/services"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func Register(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var user models.User
        if err := c.BodyParser(&user); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }

        if err := services.CreateUser(db, &user); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }

        return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
    }
}

func Login(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var user models.User
        if err := c.BodyParser(&user); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }

        token, err := services.LoginUser(db, user.Username, user.Password)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
    }
}

func Logout(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if err := services.BlacklistToken(token); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
    }
}
