package controllers

import (
    "restapi/models"
    "restapi/services"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func GetUser(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uint)
        user, err := services.GetUserByID(db, userID)
        if err != nil {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
        }
        return c.Status(fiber.StatusOK).JSON(user)
    }
}

func UpdateUser(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uint)
        var user models.User
        if err := c.BodyParser(&user); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        if err := services.UpdateUser(db, userID, &user); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
    }
}

func DeleteUser(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uint)
        if err := services.DeleteUser(db, userID); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
    }
}
