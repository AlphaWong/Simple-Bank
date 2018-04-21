package models

import (
	"log"

	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"
)

var o *Operation

const Deposit = "DEPOSIT"
const Withdraw = "WITHDRAW"

const HongKongDollar = "HKD"

type Operation struct{}

func NewOperationModel() *Operation {
	if nil == o {
		o = new(Operation)
	}
	return o
}

func newOperation(account *types.Account, currency string, amount float64) *types.Operation {
	operation := new(types.Operation)
	operation.Account = account
	operation.Amount = amount
	operation.Currency = currency
	return operation
}

func NewDeposit(account *types.Account, currency string, amount float64) *types.Operation {
	operation := newOperation(account, currency, amount)
	operation.Type = Deposit
	return operation
}

func NewWithdraw(account *types.Account, currency string, amount float64) *types.Operation {
	operation := newOperation(account, currency, amount)
	operation.Type = Withdraw
	operation.Amount = operation.Amount * -1
	return operation
}

func (*Operation) Create(operation *types.Operation) (int64, error) {
	utils.OrmInstance.Begin()
	id, err := utils.OrmInstance.Insert(operation)
	if nil != err {
		utils.OrmInstance.Rollback()
		log.Printf("Rollback operation %v, error %v", operation, err)
		return 0, err
	}
	utils.OrmInstance.Commit()
	return id, nil
}

func (*Operation) Get(id int64) (*types.Operation, error) {
	operation := new(types.Operation)
	operation.Id = id

	err := utils.OrmInstance.Read(operation)
	if nil != err {
		return nil, err
	}
	return operation, nil
}
