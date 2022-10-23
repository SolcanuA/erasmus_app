package middleware

import (
	"github.com/darkykek/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func JWTError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed token" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func TokenProtection() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.GetKey("SECRET")),
		ErrorHandler: JWTError,
	})
}
