package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	Number   string
	Name     string
	ExpMonth int
	ExpYear  int
	Cvv      string
}

func NewCreditCard(number string, name string, expMonth int, expYear int, cvv string) (*CreditCard, error) {
	cc := &CreditCard{
		Number:   number,
		Name:     name,
		ExpMonth: expMonth,
		ExpYear:  expYear,
		Cvv:      cvv,
	}

	err := cc.IsValid()
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (cc *CreditCard) IsValid() error {
	err := cc.ValidateNumber()
	if err != nil {
		return err
	}

	err = cc.ValidateMonth()
	if err != nil {
		return err
	}

	err = cc.ValidateYear()
	if err != nil {
		return err
	}

	err = cc.IsExpired()
	if err != nil {
		return err
	}

	return nil
}

func (cc *CreditCard) IsExpired() error {
	now := time.Now()

	if cc.ExpYear < now.Year() {
		return errors.New("credit card is expired")
	}

	if cc.ExpYear == now.Year() && cc.ExpMonth < int(now.Month()) {
		return errors.New("credit card is expired")
	}

	return nil
}

func (cc *CreditCard) ValidateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	if !re.MatchString(cc.Number) {
		return errors.New("invalid credit card number")
	}

	return nil
}

func (cc *CreditCard) ValidateMonth() error {
	if cc.ExpMonth < 1 || cc.ExpMonth > 12 {
		return errors.New("invalid credit card expiration month")
	}

	return nil
}

func (cc *CreditCard) ValidateYear() error {
	if cc.ExpYear <= 0 {
		return errors.New("invalid credit card expiration year")
	}

	return nil
}
