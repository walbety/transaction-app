package rest

import (
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"github.com/walbety/transaction-app/transaction-service/internal/config"
	"math/big"
	"time"
)

func mapRequestToTransaction(req TransactionRequest) canonical.Transaction {

	// already validated in validation.go
	date, _ := time.Parse(config.Env.Validations.Rest.DateFormat, req.Date)
	br := big.NewRat(1,1) // todo: look for better way...
	amount, _  := br.SetString(req.Amount)

	return canonical.Transaction{
		Amount:      *amount,
		Date:        date,
		Description: req.Description,
	}

}
