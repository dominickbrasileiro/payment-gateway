package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       int
	CreditCard   *CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	return t
}

func (t *Transaction) SetCreditCard(cc *CreditCard) {
	t.CreditCard = cc
}

func (t *Transaction) IsValid() error {
	if t.Amount > 100000 {
		return errors.New("insufficient funds")
	}
	if t.Amount < 100 {
		return errors.New("amount must be greater than 100")
	}
	return nil
}
