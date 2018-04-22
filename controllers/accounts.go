package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strconv"

	"gitlab.com/Simple-Bank/models"
	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"

	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}

// @router /v1/accounts/create [post]
func (this *AccountController) Create() {
	var customer types.Customer
	err := json.NewDecoder(bytes.NewReader(this.Ctx.Input.RequestBody)).Decode(&customer)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, utils.ErrorMessageInvalidJSON, http.StatusBadRequest)
		this.StopRun()
	}
	customerModel := models.NewCustomerModel()
	customerModel.Create(&customer)

	accountModel := models.NewAccountModel()
	accountId, _ := accountModel.Create(&customer)
	account, _ := accountModel.Get(accountId)
	account.Customer = nil

	this.Data["json"] = account
	this.ServeJSON()
}

// @router /v1/customers/:id/accounts/add [post]
func (this *AccountController) Add() {
	customerId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	customerModel := models.NewCustomerModel()
	customer, err := customerModel.Get(customerId)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, "Customer not found", http.StatusNotFound)
		this.StopRun()
	}
	accountModel := models.NewAccountModel()
	accountId, _ := accountModel.Create(customer)
	account, _ := accountModel.Get(accountId)
	account.Customer = nil

	this.Data["json"] = account
	this.ServeJSON()
}

// @router /v1/accounts/:id/close [put]
func (this *AccountController) Close() {
	accountModel := models.NewAccountModel()

	accountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	account, err := accountModel.Get(accountId)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}

	account.Active = false
	_, err = accountModel.Update(account)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}
	account.Customer = nil

	this.Data["json"] = account
	this.ServeJSON()
}

func transactionCommon(this *AccountController) (*types.Account, *types.TransactionRequest) {
	accountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var transactionRequest types.TransactionRequest
	err := json.NewDecoder(bytes.NewReader(this.Ctx.Input.RequestBody)).Decode(&transactionRequest)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, utils.ErrorMessageInvalidJSON, http.StatusBadRequest)
		this.StopRun()
	}

	accountModel := models.NewAccountModel()
	account, err := accountModel.Get(accountId)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusNotFound)
		this.StopRun()
	}
	return account, &transactionRequest
}

// @router /v1/accounts/:id/deposit [put]
func (this *AccountController) Deposit() {
	account, transactionRequest := transactionCommon(this)

	transactionModel := models.NewTransactionModel()
	transaction := models.NewDeposit(account, transactionRequest.Currency, transactionRequest.Amount)
	id, err := transactionModel.Create(transaction)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}

	transaction, _ = transactionModel.Get(id)
	this.Data["json"] = transaction
	this.ServeJSON()
}

// @router /v1/accounts/:id/withdraw [put]
func (this *AccountController) Withdraw() {
	account, transactionRequest := transactionCommon(this)
	transactionModel := models.NewTransactionModel()
	transaction := models.NewWithdraw(account, transactionRequest.Currency, transactionRequest.Amount)

	if ok, err := IsValidWithdrawTransaction(account.Id, transaction); !ok {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}

	id, err := transactionModel.Create(transaction)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}

	transaction, _ = transactionModel.Get(id)
	this.Data["json"] = transaction
	this.ServeJSON()
}

// @router /v1/accounts/:id/balance [get]
func (this *AccountController) Balance() {
	// TODO: Check account existence
	accountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	accountModel := models.NewAccountModel()
	balance, err := accountModel.GetCurrentBalance(accountId)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}
	this.Data["json"] = map[string]float64{
		"balance": balance,
	}
	this.ServeJSON()
}

func IsValidCurrency(transaction *types.Transaction) (bool, error) {
	if transaction.Currency != models.HongKongDollar {
		return false, errors.New("Only support HKD")
	}
	return true, nil
}

func IsPositiveNumberAfterTransaction(transaction *types.Transaction, balance float64) (bool, error) {
	if math.Abs(transaction.Amount) > balance {
		return false, errors.New("Account balance not enough")
	}
	return true, nil
}

func IsValidWithdrawTransaction(id int64, transaction *types.Transaction) (bool, error) {
	if ok, err := IsValidCurrency(transaction); nil != err {
		return ok, err
	}

	accountModel := models.NewAccountModel()
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

func WithCrossCustomerAccountServiceCharge(amount float64) float64 {
	return amount + utils.CrossCustomerSendServiceCharge
}

// func IsPositiveNumberAfterCrossCustomerSend(id int64, transaction *types.Transaction) (bool, error) {
// 	accountModel := models.NewAccountModel()
// 	balance, err := accountModel.GetCurrentBalance(id)

// 	if ok, err := IsPositiveNumberAfterTransaction(transaction, balance); nil != err {
// 		return ok, err
// 	}
// 	return true, nil
// }

// func IsValidSendOperation(senderAccount, receiverAccount *types.Account, transaction *types.Transaction) (bool, error) {
// 	if !IsOwnBySameCustomer(senderAccount, receiverAccount) {
// 		transaction.Amount = WithCrossCustomerAccountServiceCharge(transaction.Amount)
// 	}

// 	if ok, err := IsValidWithdrawOperation(senderAccount.Id, operation); nil != err {
// 		return ok, err
// 	}
// }

// // @router /v1/accounts/:from/send/:to [post]
// func (this *AccountController) Send() {
// 	senderAccountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":from"), 10, 64)
// 	receiverAccountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":to"), 10, 64)
// 	accountModel := models.NewAccountModel()
// 	senderAccount, err := accountModel.Get(senderAccountId)
// 	if nil != err {
// 		http.Error(this.Ctx.ResponseWriter, "Sender account not found", http.StatusNotFound)
// 		this.StopRun()
// 	}
// 	receiverAccount, err := accountModel.Get(receiverAccountId)
// 	if nil != err {
// 		http.Error(this.Ctx.ResponseWriter, "Receiver account not found", http.StatusNotFound)
// 		this.StopRun()
// 	}
// }
