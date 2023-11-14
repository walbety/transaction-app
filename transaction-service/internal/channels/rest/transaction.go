package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"net/http"
	"time"
)

const (
	TransactionBasePath = "/transaction"

	ParamId = "id"
	ParamDate = "date"
	ParamCurrency = "currency"
)

func getPurchase(c *fiber.Ctx) error {

	id := c.Query(ParamId)
	date := c.Query(ParamDate)
	currency := c.Query(ParamCurrency)

	err := validateGetPurchaseRequest(id,date,currency)
	if err != nil {
		return c.Status(err.(Error).HttpCode).JSON(err.(Error).Message)
	}

	// already validated
 	dateFormtd,_ := time.Parse(config.Env.Validations.Rest.DateFormat, date)

	transaction, err := svc.GetLatestExchangeRateFromCurrencyAndDate(c.UserContext(), id, currency, dateFormtd)
	if err != nil { //todo create map of errors before return 500
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(transaction)
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
		return c.Status(err.(Error).HttpCode).JSON(err.(Error).Message)
	}

	id, err := svc.SavePurchase(c.UserContext(), mapRequestToTransaction(*transaction))
	if err != nil { //todo create map of errors before return 500
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON("id:" + id)
}