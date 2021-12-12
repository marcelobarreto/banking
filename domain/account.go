package domain

import (
	"github.com/marcelobarreto/banking/dto"
	"github.com/marcelobarreto/banking/errs"
)

const DBTSLayout = "2006-01-02 15:04:05"

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
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}
