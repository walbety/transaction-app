package rest

import (
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"net/http"
	"time"
)

var (
	ErrDescriptionMaxLen  = Error{Code: "2001", Message: "Description max length (50) exceeded.", HttpCode: http.StatusBadRequest}
	ErrZeroAmount  = Error{Code: "2002", Message: "Amount cannot be zero.", HttpCode: http.StatusBadRequest}
	ErrDateWrongFormat  = Error{Code: "2003", Message: "Field date should be in DD/MM/YYYY format.", HttpCode: http.StatusBadRequest}
	ErrRequiredDate  = Error{Code: "2004", Message: "Date field is required.", HttpCode: http.StatusBadRequest}
)

func validateSaveTransactionRequest(request TransactionRequest) error {

	if request.Description != "" && len(request.Description) > 50{
		return ErrDescriptionMaxLen
	}
	
	if request.Amount == 0{
		return ErrZeroAmount
	}

	if request.Date == "" {
		return ErrRequiredDate
	}

	if _, err := time.Parse(config.Env.Validations.Rest.DateFormat, request.Date); err != nil {
		return ErrDateWrongFormat
	}

	return nil
}