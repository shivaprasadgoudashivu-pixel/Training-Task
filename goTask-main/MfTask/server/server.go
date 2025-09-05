package server

import (
	"keycloak-demo/config"
	"keycloak-demo/handler"

	"github.com/gofiber/fiber/v2"
)

func New(cfg config.Config) *fiber.App {
	app := fiber.New()

	app.Post("/login", handler.LoginHandler(cfg))

	return app
}
