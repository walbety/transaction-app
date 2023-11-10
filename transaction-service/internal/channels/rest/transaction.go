package rest

import (
	"github.com/gofiber/fiber/v2"
)

const (
	TRANSACTION_BASE_PATH = "/transaction"
)

func save(c *fiber.Ctx) error {
	users, err := svc.ListUsers(c.UserContext()) // todo create transaction-service
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(users)
}
