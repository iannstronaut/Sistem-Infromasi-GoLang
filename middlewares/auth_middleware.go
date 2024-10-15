package middlewares

import (
	"restapi/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header required"})
    }

    token := strings.TrimPrefix(authHeader, "Bearer ")

    claims, err := utils.ValidateToken(token)

    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
    }
    c.Locals("userID", claims.UserID)
    return c.Next()
}

