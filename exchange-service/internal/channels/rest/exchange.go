package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/walbety/transaction-app/exchange-service/internal/canonical"
	"net/http"
	"time"
)

const (
	ExchangeBasePath = "/exchange"

	ParamCurrency = "currency"
	ParamDate     ="date"
)

func getExchange(c *fiber.Ctx) error {

	currency := c.Query(ParamCurrency)
	dateParam := c.Query(ParamDate)

	if currency == ""{
		return c.Status(http.StatusBadRequest).JSON(canonical.ErrCurrencyIsRequired)
	}

	date,err := time.Parse("02-01-2006", dateParam)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(canonical.ErrDateWrongFormat)
	}

	exchange, err := svc.GetLatestExchangeRateFromCurrencyAndDate(c.UserContext(),currency, date)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(exchange)
}
