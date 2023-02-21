package entity

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

var validNumber = "4193523830170205"
var validMonth = int(time.Now().Month())
var validYear = time.Now().Year()

var invalidNumber = "40000000000000000"
var invalidMonth = 13
var invalidYear = -50

func TestValidCreditCard(t *testing.T) {
	_, err := NewCreditCard(validNumber, "John Doe", validMonth, validYear, "123")
	assert.NilError(t, err)
}

func TestExpiredCreditCard(t *testing.T) {
	_, err := NewCreditCard(validNumber, "John Doe", 12, 2010, "123")
	assert.Error(t, err, "credit card is expired")
}

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard(invalidNumber, "John Doe", validMonth, validYear, "123")
	assert.Error(t, err, "invalid credit card number")
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard(validNumber, "John Doe", invalidMonth, validYear, "123")
	assert.Error(t, err, "invalid credit card expiration month")
}

func TestCreditCardExpirationYear(t *testing.T) {
	_, err := NewCreditCard(validNumber, "John Doe", validMonth, invalidYear, "123")
	assert.Error(t, err, "invalid credit card expiration year")
}
