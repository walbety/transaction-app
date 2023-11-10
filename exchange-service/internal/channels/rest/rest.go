package rest

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/exchange-service/internal/config"
	"github.com/walbety/transaction-app/exchange-service/internal/service"
	"net/http"
)

var (
	svc service.Service
	app *fiber.App
)

func Start(service service.Service) error {
	svc = service

	app = fiber.New()

	public := app.Group("/")
	public.Get("/", welcome)
	// todo add health-check endpoint
	private := app.Group("/v1")
	private.Get(USER_BASE_PATH, listUser)
	private.Get(PRODUCT_BASE_PATH, listProduct)

	app.Add(http.MethodGet, "/", welcome)

	return app.Listen(fmt.Sprintf(":%s", config.Env.RestPort))
}

func Stop(ctx context.Context) {
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.WithContext(ctx).WithError(err).Fatal("Error shutting down rest server")
	}
}

func welcome(c *fiber.Ctx) error {
	return c.Status(200).JSON("WELCOME!!")
}
