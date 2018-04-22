package models

import (
	"errors"
	"net/http"

	"github.com/astaxie/beego/orm"

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
	if orm.ErrNoRows == err {
		return nil, errors.New(utils.ErrorMessageAccountNotFound)
	}
	if nil != err {
		return nil, err
	}

	if !account.Active {
		return nil, errors.New(utils.ErrorMessageAccountDeactivate)
	}
	return account, nil
}

func (*Account) Update(account *types.Account) (int64, error) {
	return utils.OrmInstance.Update(account)
}

func (*Account) GetCurrentBalance(id int64) (float64, error) {
	var balance float64
	err := utils.OrmInstance.Raw("SELECT SUM(amount) FROM transaction WHERE account_id = ? ORDER BY id ASC", id).QueryRow(&balance)
	if nil != err {
		return 0, err
	}
	return balance, nil
}

func GetPaymentApproval() (bool, error) {
	resp, err := http.Get(utils.PaymentApprovalURI)
	if nil != err {
		return false, err
	}
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, errors.New("Payment is declined by approval gateway")
}
