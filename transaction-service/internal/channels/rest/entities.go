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
	Amount      string `json:"amount,omitempty"`
	Date        string `json:"date,omitempty"`
	Description string `json:"description,omitempty"`
}

type TransactionResponse struct {
	Id      string `json:"id,omitempty"`
}

type ConvertedTransactionResponse struct {
	Id              string `json:"id"`
	OriginalAmount  string `json:"originalAmount"`
	ConvertedAmount string `json:"convertedAmount"`
	Currency        string `json:"currency"`
	TransactionDate string `json:"transactionDate"`
	Description     string `json:"description"`
	ExchangeRate    string `json:"exchangeRate"`
}
