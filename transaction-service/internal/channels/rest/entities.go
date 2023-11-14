package rest

import (
	"fmt"
)

type Error struct {
	Code     string
	Message  string
	HttpCode int
}

func (r Error) Error() string {
	return fmt.Sprintf("code: %s message: %s", r.Code, r.Message)
}

type TransactionRequest struct {
	Amount      float64 `json:"amount,omitempty"`
	Date        string  `json:"date,omitempty"`
	Description string  `json:"description,omitempty"`
}
