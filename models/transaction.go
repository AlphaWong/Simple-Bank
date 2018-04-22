package models

import (
	"log"

	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"
)

var t *Transaction

const Deposit = "DEPOSIT"
const Withdraw = "WITHDRAW"

const HongKongDollar = "HKD"

type Transaction struct{}

func NewTransactionModel() *Transaction {
	if nil == t {
		t = new(Transaction)
	}
	return t
}

func newTransaction(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := new(types.Transaction)
	transaction.Sender = account
	transaction.Amount = amount
	transaction.Currency = currency
	return transaction
}

func NewDeposit(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := newTransaction(account, currency, amount)
	transaction.Type = Deposit
	return transaction
}

func NewWithdraw(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := newTransaction(account, currency, amount)
	transaction.Type = Withdraw
	transaction.Amount = transaction.Amount * -1
	return transaction
}

func (*Transaction) Create(transaction *types.Transaction) (int64, error) {
	utils.OrmInstance.Begin()
	id, err := utils.OrmInstance.Insert(transaction)
	if nil != err {
		utils.OrmInstance.Rollback()
		log.Printf("Rollback transaction %v, error %v", transaction, err)
		return 0, err
	}
	utils.OrmInstance.Commit()
	return id, nil
}

func (*Transaction) Get(id int64) (*types.Transaction, error) {
	transaction := new(types.Transaction)
	transaction.Id = id

	err := utils.OrmInstance.Read(transaction)
	if nil != err {
		return nil, err
	}
	return transaction, nil
}
