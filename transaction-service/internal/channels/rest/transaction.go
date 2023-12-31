package rest

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	TransactionBasePath = "/transaction"

	ParamId       = "id"
	ParamCurrency = "currency"
)

func getPurchase(c *fiber.Ctx) error {

	id := c.Query(ParamId)
	currency := c.Query(ParamCurrency)

	err := validateGetPurchaseRequest(id, currency)
	if err != nil {
		return c.Status(err.(Error).HttpCode).JSON(err.(Error).Message)
	}

	transaction, err := svc.GetLatestExchangeRateFromCurrencyAndDate(c.UserContext(), id, currency)
	if err != nil { //todo create map of errors before return 500
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(mapTransactionToConvertedResponse(transaction))
}

func savePurchase(c *fiber.Ctx) error {

	transaction := new(TransactionRequest)
	log.Info(string(c.Body()))
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}


	err := validateSaveTransactionRequest(*transaction)
	if err != nil {
		return c.Status(err.(Error).HttpCode).JSON(err.(Error).Message)
	}

	id, err := svc.SavePurchase(c.UserContext(), mapRequestToTransaction(*transaction))
	if err != nil { //todo create map of errors before return 500
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(TransactionResponse{id})
}
