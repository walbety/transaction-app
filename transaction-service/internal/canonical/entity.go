package canonical

import (
	"math/big"
	"time"
)

type Transaction struct {
	Id          string    `bson:"_id,omitempty"` // todo check this
	Amount      string    `bson:"amount,omitempty"`
	Date        time.Time `bson:"date,omitempty"`
	Description string    `bson:"description,omitempty"`
}

type ExchangeRate struct {
	Currency     string
	ExchangeRate string
	RecordDate   time.Time
}

type ConvertedTransaction struct {
	Id              string
	OriginalAmount  big.Float
	ConvertedAmount big.Float
	Currency        string
	TransactionDate time.Time
	Description     string
	ExchangeRate    string
}
