package handler

// import (
// 	"context"
// 	"encoding/json"
// 	"keycloak-demo/config"
// 	keycloak "keycloak-demo/keyCloak"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// type LoginReq struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// func LoginHandler(cfg config.Config) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		var req LoginReq
// 		if err := c.BodyParser(&req); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
// 		}

// 		ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
// 		defer cancel()

// 		resp, data, err := keycloak.Login(ctx, cfg, req.Username, req.Password)
// 		if err != nil {
// 			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
// 		}

// 		c.Status(resp.StatusCode)
// 		var jsonResp map[string]interface{}
// 		if err := json.Unmarshal(data, &jsonResp); err != nil {
// 			return c.Send(data)
// 		}
// 		return c.JSON(jsonResp)
// 	}
// }
