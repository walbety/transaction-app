package rest

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"github.com/walbety/transaction-app/transaction-service/internal/service"
)

var (
	svc service.Transaction
	app *fiber.App
)

func Start(service service.Transaction) error {
	svc = service

	app = fiber.New()

	private := app.Group("/v1")

	private.Get(TransactionBasePath, getPurchase)
	private.Post(TransactionBasePath, savePurchase)

	return app.Listen(fmt.Sprintf(":%s", config.Env.RestPort))
}

func Stop(ctx context.Context) {
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.WithContext(ctx).WithError(err).Fatal("Error shutting down rest server")
	}
}
