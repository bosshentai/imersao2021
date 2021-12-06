package factory

type RepositoryFactory interface {

	CreateTransactionRepository() repository.TransactionRepository
}