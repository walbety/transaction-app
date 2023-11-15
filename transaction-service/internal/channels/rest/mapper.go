package rest

import (
	"fmt"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"math/big"
	"strings"
	"time"
)

func mapRequestToTransaction(req TransactionRequest) canonical.Transaction {

	// already validated in validation.go

	date, _ := time.Parse(config.Env.Validations.Rest.DateFormat, req.Date)

	return canonical.Transaction{
		Amount:      req.Amount,
		Date:        date,
		Description: req.Description,
	}

}

func mapTransactionToConvertedResponse(converted canonical.ConvertedTransaction) ConvertedTransactionResponse {
	transactionDateFmtd := converted.TransactionDate.Format(config.Env.Validations.Rest.DateFormat)

	originalAmountStr, _ := asStringFromFloat(2, &converted.OriginalAmount)
	convertedAmountStr, _ := asStringFromFloat(2, &converted.ConvertedAmount)
	result := ConvertedTransactionResponse{
		Id:              converted.Id,
		OriginalAmount:  originalAmountStr,
		ConvertedAmount: convertedAmountStr,
		Currency:        converted.Currency,
		TransactionDate: transactionDateFmtd,
		Description:     converted.Description,
		ExchangeRate:    converted.ExchangeRate,
	}

	return result
}

func asStringFromFloat(precision int, amount *big.Float) (string, error) {
	fmtString := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(strings.TrimRight(strings.TrimRight(fmt.Sprintf(fmtString, amount), ""), ".")), nil
}
