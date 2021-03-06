package process_transaction

import (
	"github.com/bosshentai/imersao-gateway/domain/entity"
	"github.com/bosshentai/imersao-gateway/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOuput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	cc, invalidCC := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)

	if invalidCC != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, invalidCC.Error())
		if err != nil {
			return TransactionDtoOuput{}, err
		}

		output := TransactionDtoOuput{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}

		return output, nil
	}

	transaction.SetCreditCard(*cc)
	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, invalidCC.Error())
		if err != nil {
			return TransactionDtoOuput{}, err
		}

		output := TransactionDtoOuput{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidTransaction.Error(),
		}

		return output, nil
	}

	return TransactionDtoOuput{}, nil

}
