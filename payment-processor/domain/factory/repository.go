package factory

import "github.com/dominickbrasileiro/payment-gateway/payment-processor/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
