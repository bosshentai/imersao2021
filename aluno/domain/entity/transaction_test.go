package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTRansaction_IsValid (t *testing.T){
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 900
	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_IsNotValidWithAmountGreaterThan1000 (t *testing.T){
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001
	err := transaction.IsValid()
	assert.Error(t,err)
	// assert.NotNil(t, transaction.IsValid())
	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransaction_IsNotValidWithAmountLessthan1 (t *testing.T){
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0
	err := transaction.IsValid()
	assert.Error(t,err)
	// assert.NotNil(t, transaction.IsValid())
	assert.Equal(t, "the amount must be greater than 1", err.Error())
}