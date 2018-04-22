package models

import (
	"fmt"
	"log"
	"math"

	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"
)

var t *Transaction

type Transaction struct{}

func NewTransactionModel() *Transaction {
	if nil == t {
		t = new(Transaction)
	}
	return t
}

func newTransaction(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := new(types.Transaction)
	transaction.Account = account
	transaction.Amount = amount
	transaction.Currency = currency
	return transaction
}

func NewDeposit(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := newTransaction(account, currency, amount)
	transaction.Type = utils.Deposit
	return transaction
}

func NewWithdraw(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := newTransaction(account, currency, amount)
	transaction.Type = utils.Withdraw
	transaction.Amount = transaction.Amount * -1
	return transaction
}

func NewCrossCustomerServiceCharge(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := newTransaction(account, currency, amount)
	transaction.Account = account
	transaction.Type = utils.CrossCustomerServiceCharge
	transaction.Amount = transaction.Amount * -1
	return transaction
}

func NewSendTransaction(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := newTransaction(account, currency, amount)
	transaction.Account = account
	transaction.Type = utils.Transfer
	transaction.Amount = transaction.Amount * -1
	transaction.Remark = fmt.Sprintf("Account %v send money %v %v", account.Id, currency, amount)
	return transaction
}

func NewReceiveTransaction(account *types.Account, currency string, amount float64) *types.Transaction {
	transaction := NewSendTransaction(account, currency, amount)
	transaction.Amount = math.Abs(transaction.Amount)
	transaction.Remark = fmt.Sprintf("Account %v receive money %v %v", account.Id, currency, amount)
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

func (*Transaction) Send(transactions ...*types.Transaction) ([]*types.Transaction, error) {
	utils.OrmInstance.Begin()
	for _, transaction := range transactions {
		_, err := utils.OrmInstance.Insert(transaction)
		if nil != err {
			utils.OrmInstance.Rollback()
			log.Printf("Rollback transaction %v, error %v", transaction, err)
			return make([]*types.Transaction, 0), err
		}
	}
	utils.OrmInstance.Commit()
	return transactions, nil
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
