package factory

import "github.com/bosshentai/imersao-gateway/domain/repository"

type RepositoryFactory interface {

	CreateTransactionRepository() repository.TransactionRepository
}