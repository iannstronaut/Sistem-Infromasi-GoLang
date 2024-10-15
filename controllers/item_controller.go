package controllers

import (
    "restapi/models"
    "restapi/services"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func GetItems(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uint)
        items, err := services.GetItemsByUserID(db, userID)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(items)
    }
}

func AddItem(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uint)
        var item models.Item
        if err := c.BodyParser(&item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        item.UserID = userID
        if err := services.CreateItem(db, &item); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Item created successfully"})
    }
}

func UpdateItem(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uint)
        itemID := c.Params("id")
        var item models.Item
        if err := c.BodyParser(&item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        if err := services.UpdateItem(db, itemID, userID, &item); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item updated successfully"})
    }
}

func DeleteItem(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Locals("userID").(uint)
        itemID := c.Params("id")
        if err := services.DeleteItem(db, itemID, userID); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item deleted successfully"})
    }
}
