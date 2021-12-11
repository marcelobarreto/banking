package domain

import (
	"github.com/marcelobarreto/banking/dto"
	"github.com/marcelobarreto/banking/errs"
)

type Account struct {
	AccountID   string `db:"account_id"`
	CustomerID  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}
