package canonical

import (
	"fmt"
	"net/http"
)

var (
	ErrInvalidCurrency  = Error{Code: "1001", Message: "Invalid currency.", HttpCode: http.StatusBadRequest}
	ErrCurrencyIsRequired  = Error{Code: "1002", Message: "Field currency is required.", HttpCode: http.StatusBadRequest}
	ErrDateWrongFormat  = Error{Code: "1003", Message: "Field date should be in DD/MM/YYYY format.", HttpCode: http.StatusBadRequest}
	ErrTreasuryServiceError  = Error{Code: "1009", Message: "Unexpected error with treasury service integration.", HttpCode: http.StatusInternalServerError}
)

type Error struct {
	Code     string
	Message  string
	HttpCode int
}

func (r Error) Error() string {
	return fmt.Sprintf("code: %s message: %s", r.Code, r.Message)
}
