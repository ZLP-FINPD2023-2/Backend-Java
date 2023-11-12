package models

type TransactionRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Reason   string  `json:"reason"`
}
