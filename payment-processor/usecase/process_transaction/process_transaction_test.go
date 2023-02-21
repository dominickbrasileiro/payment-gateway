package usecase

import (
	"testing"
	"time"

	mock_repository "github.com/dominickbrasileiro/payment-gateway/payment-processor/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := &TransactionDtoInput{
		ID:                 "1",
		AccountID:          "1",
		Amount:             50000,
		CreditCardNumber:   "40000000000000000",
		CreditCardName:     "John Doe",
		CreditCardExpMonth: int(time.Now().Month()),
		CreditCardExpYear:  time.Now().Year(),
		CreditCardCvv:      "123",
	}

	expectedOutput := &TransactionDtoOutput{
		ID:           "1",
		Status:       "rejected",
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.
		EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	sut := NewProcessTransactionUsecase(repositoryMock)
	output, err := sut.Execute(input)

	assert.NilError(t, err)
	assert.Equal(t, *output, *expectedOutput)
}
