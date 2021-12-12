package dto

import (
	"github.com/marcelobarreto/banking/errs"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountID       string  `json:"account_id"`
	CustomerID      string  `json:"customer_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

type TransactionResponse struct {
	TransactionID   string  `json:"transaction_id"`
	AccountID       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit and withdrawal")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount can't be less than 0")
	}

	return nil
}
