package models

import (
	"errors"

	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"
)

var a *Account

type Account struct{}

func NewAccountModel() *Account {
	if nil == a {
		a = new(Account)
	}
	return a
}

func (*Account) Create(customer *types.Customer) (int64, error) {
	account := new(types.Account)
	account.Customer = customer
	account.Active = true
	return utils.OrmInstance.Insert(account)
}

func (*Account) Get(id int64) (*types.Account, error) {
	account := new(types.Account)
	account.Id = id

	err := utils.OrmInstance.Read(account)
	if nil != err {
		return nil, err
	}

	if !account.Active {
		return nil, errors.New("Account is inactve")
	}
	return account, nil
}

func (*Account) Update(account *types.Account) (int64, error) {
	return utils.OrmInstance.Update(account)
}

func (*Account) GetCurrentBalance(id int64) (float64, error) {
	var balance float64
	err := utils.OrmInstance.Raw("SELECT SUM(amount) FROM transaction WHERE sender_id = ? ORDER BY id ASC", id).QueryRow(&balance)
	if nil != err {
		return 0, err
	}
	return balance, nil
}
