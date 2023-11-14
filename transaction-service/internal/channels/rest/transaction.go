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
	//
	//currency := c.Query(ParamCurrency)
	//dateParam := c.Query(ParamDate)
	//
	//if currency == ""{
	//	return c.Status(http.StatusBadRequest).JSON(canonical.ErrCurrencyIsRequired)
	//}

	log.WithContext(c.UserContext()).Info("CHEGUEI AQUIIIII")

	//date,err := time.Parse("02-01-2006", dateParam)
	//if err != nil {
	//	return c.Status(http.StatusUnprocessableEntity).JSON(canonical.ErrDateWrongFormat)
	//}

	//exchange, err := svc.GetLatestExchangeRateFromCurrencyAndDate(c.UserContext(),currency, date)
	//if err != nil {
	//	return c.Status(http.StatusInternalServerError).JSON(err)
	//}
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

	log.WithContext(c.UserContext()).Info("CHEGUEI AQUIIIII")

	//date,err := time.Parse("02-01-2006", dateParam)
	//if err != nil {
	//	return c.Status(http.StatusUnprocessableEntity).JSON(canonical.ErrDateWrongFormat)
	//}

	//exchange, err := svc.GetLatestExchangeRateFromCurrencyAndDate(c.UserContext(),currency, date)
	//if err != nil {
	//	return c.Status(http.StatusInternalServerError).JSON(err)
	//}
	return c.Status(http.StatusOK).JSON("")
}