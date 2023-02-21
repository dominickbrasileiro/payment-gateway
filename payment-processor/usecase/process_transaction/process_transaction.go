package usecase

import (
	"github.com/dominickbrasileiro/payment-gateway/payment-processor/domain/entity"
	"github.com/dominickbrasileiro/payment-gateway/payment-processor/domain/repository"
)

type ProcessTransactionUsecase struct {
	repository repository.TransactionRepository
}

func NewProcessTransactionUsecase(repository repository.TransactionRepository) *ProcessTransactionUsecase {
	return &ProcessTransactionUsecase{
		repository: repository,
	}
}

func (u *ProcessTransactionUsecase) Execute(input *TransactionDtoInput) (*TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	_, ccErr := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpMonth, input.CreditCardExpYear, input.CreditCardCvv)
	if ccErr != nil {
		err := u.repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", ccErr.Error())
		if err != nil {
			return nil, err
		}
		output := &TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       "rejected",
			ErrorMessage: ccErr.Error(),
		}
		return output, nil
	}

	return &TransactionDtoOutput{}, nil
}
