package repository

type TransactionRepository interface {
	Insert(id sring,account string, amount float64, status string, errorMessage string) error
}