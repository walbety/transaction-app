package rest


import (
	"github.com/gofiber/fiber/v2"
)

const (
	PRODUCT_BASE_PATH = "/product"
)

func listProduct(c *fiber.Ctx) error {
	users, err := svc.ListUsers(c.UserContext())
	if err != nil {
		return c.Status(500).JSON(err)
	}
	return c.Status(200).JSON(users)
}