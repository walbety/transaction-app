package canonical

import (
	"math/big"
	"time"
)

type Transaction struct {
	Amount      big.Rat   `bson:"amount,omitempty"`
	Date        time.Time `bson:"date,omitempty"`
	Description string    `bson:"description,omitempty"`
}


type ExchangeRate struct {
	Currency     string
	ExchangeRate string
	RecordDate   time.Time
}
