package models

import (
	"errors"
	"math"

	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"
)

func IsValidCurrency(transaction *types.Transaction) (bool, error) {
	if transaction.Currency != utils.HongKongDollar {
		return false, errors.New(utils.ErrorMessageCurrencyNotSupport)
	}
	return true, nil
}

func IsPositiveNumberAfterTransaction(transaction *types.Transaction, balance float64) (bool, error) {
	if math.Abs(transaction.Amount) > balance {
		return false, errors.New(utils.ErrorMessageAccountBalanceNotEnough)
	}
	return true, nil
}

func IsValidTransaction(id int64, transaction *types.Transaction) (bool, error) {
	if ok, err := IsValidCurrency(transaction); nil != err {
		return ok, err
	}

	accountModel := NewAccountModel()
	balance, err := accountModel.GetCurrentBalance(id)
	if nil != err {
		return false, err
	}
	if ok, err := IsPositiveNumberAfterTransaction(transaction, balance); nil != err {
		return ok, err
	}

	return true, nil
}

func IsOwnBySameCustomer(senderAccount, receiverAccount *types.Account) bool {
	return senderAccount.Customer.Id == receiverAccount.Customer.Id
}
