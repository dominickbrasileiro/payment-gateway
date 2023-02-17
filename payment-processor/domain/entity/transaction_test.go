package entity

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestValidTransaction(t *testing.T) {
	tr := NewTransaction()
	tr.ID = "1"
	tr.AccountID = "1"
	tr.Amount = 50000

	err := tr.IsValid()

	assert.NilError(t, err)
}

func TestTransactionWithAmountLessThan100(t *testing.T) {
	tr := NewTransaction()
	tr.ID = "1"
	tr.AccountID = "1"
	tr.Amount = 10

	err := tr.IsValid()

	assert.Error(t, err, "amount must be greater than 100")
}

func TestTransactionWithAmountGreaterThan100000(t *testing.T) {
	tr := NewTransaction()
	tr.ID = "1"
	tr.AccountID = "1"
	tr.Amount = 100001

	err := tr.IsValid()

	assert.Error(t, err, "insufficient funds")
}
