package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number   string
	name     string
	expMonth int
	expYear  int
	cvv      string
}

func NewCreditCard(number string, name string, expMonth int, expYear int, cvv string) (*CreditCard, error) {
	cc := &CreditCard{
		number:   number,
		name:     name,
		expMonth: expMonth,
		expYear:  expYear,
		cvv:      cvv,
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

	if cc.expYear < now.Year() {
		return errors.New("card is expired")
	}

	if cc.expYear == now.Year() && cc.expMonth < int(now.Month()) {
		return errors.New("card is expired")
	}

	return nil
}

func (cc *CreditCard) ValidateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	if !re.MatchString(cc.number) {
		return errors.New("invalid number")
	}

	return nil
}

func (cc *CreditCard) ValidateMonth() error {
	if cc.expMonth < 1 || cc.expMonth > 12 {
		return errors.New("invalid expiration month")
	}

	return nil
}

func (cc *CreditCard) ValidateYear() error {
	if cc.expYear <= 0 {
		return errors.New("invalid expiration year")
	}

	return nil
}
