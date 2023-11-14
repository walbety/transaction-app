package rest

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	TransactionBasePath = "/transaction"

	ParamCurrency = "currency"
	ParamDate     ="date"
)

func getPurchase(c *fiber.Ctx) error {


	log.WithContext(c.UserContext()).Info("CHEGUEI AQUIIIII")


	return c.Status(http.StatusOK).JSON("")
}

func savePurchase(c *fiber.Ctx) error {

	transaction := new(TransactionRequest)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	err := validateSaveTransactionRequest(*transaction)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	id, err := svc.SavePurchase(c.UserContext(), mapRequestToTransaction(*transaction))
	if err != nil { //todo create map of errors before return 500
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON("id:" + id)
}