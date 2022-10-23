package middleware

import (
	"github.com/darkykek/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func AdminProtection() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.GetKey("SECRET")),
		ErrorHandler: JWTError,
	})
}
