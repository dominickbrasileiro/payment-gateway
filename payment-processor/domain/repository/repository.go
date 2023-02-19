package repository

type TransactionRepository interface {
	Insert(id string, accountID string, amount int, status string, errMessage string) error
}
