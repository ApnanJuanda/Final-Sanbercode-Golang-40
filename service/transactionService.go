package service

import "project/model"

type TransactionService interface {
	AddTransaction(AddTransaction *model.AddTransaction, customerId int) (*model.ResponseTransaction, error)
	GetTransaction(txId int) (*model.ResponseTransaction, error)
}