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

// // @router /v1/accounts/:id [get]
// func (this *AccountController) Get() {
// 	accountID, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
// 	account := utils.NewAccountInstance()
// 	account.ID = accountID
// 	this.Data["json"] = account
// 	this.ServeJSON()
// }

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
	customerID, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	customerModel := models.NewCustomerModel()
	customer, err := customerModel.Get(customerID)
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

	accountID, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	account, err := accountModel.Get(accountID)
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

func operationCommon(this *AccountController) (*types.Account, *types.OperationRequest) {
	accountID, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var operationRequest types.OperationRequest
	err := json.NewDecoder(bytes.NewReader(this.Ctx.Input.RequestBody)).Decode(&operationRequest)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, utils.ErrorMessageInvalidJSON, http.StatusBadRequest)
		this.StopRun()
	}

	accountModel := models.NewAccountModel()
	account, err := accountModel.Get(accountID)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusNotFound)
		this.StopRun()
	}
	return account, &operationRequest
}

// @router /v1/accounts/:id/deposit [put]
func (this *AccountController) Deposit() {
	account, operationRequest := operationCommon(this)

	operationModel := models.NewOperationModel()
	operation := models.NewDeposit(account, operationRequest.Currency, operationRequest.Amount)
	id, err := operationModel.Create(operation)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}

	operation, _ = operationModel.Get(id)
	this.Data["json"] = operation
	this.ServeJSON()
}

// @router /v1/accounts/:id/withdraw [put]
func (this *AccountController) Withdraw() {
	account, operationRequest := operationCommon(this)
	operationModel := models.NewOperationModel()
	operation := models.NewWithdraw(account, operationRequest.Currency, operationRequest.Amount)

	if ok, err := IsValidAccountOperation(account.Id, operation); !ok {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}

	id, err := operationModel.Create(operation)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}

	operation, _ = operationModel.Get(id)
	this.Data["json"] = operation
	this.ServeJSON()
}

// @router /v1/accounts/:id/balance [get]
func (this *AccountController) Balance() {
	accountID, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	accountModel := models.NewAccountModel()
	balance, err := accountModel.GetCurrentBalance(accountID)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusBadRequest)
		this.StopRun()
	}
	this.Data["json"] = map[string]float64{
		"balance": balance,
	}
	this.ServeJSON()
}

func IsValidAccountOperation(id int64, operation *types.Operation) (bool, error) {
	accountModel := models.NewAccountModel()
	balance, err := accountModel.GetCurrentBalance(id)
	if nil != err {
		return false, err
	}

	if operation.Currency != models.HongKongDollar {
		return false, errors.New("Only support HKD")
	}
	if operation.Type == models.Withdraw && math.Abs(operation.Amount) > balance {
		return false, errors.New("Account balance not enough")
	}
	return true, nil
}

// @router /v1/accounts/send [post]
// func (this *AccountController) Send() {
// 	transaction := utils.NewTransactionInstance()
// 	err := json.NewDecoder(bytes.NewReader(this.Ctx.Input.RequestBody)).Decode(&transaction)
// 	if nil != err {
// 		http.Error(this.Ctx.ResponseWriter, utils.ErrorMessageInvalidJSON, http.StatusBadRequest)
// 		this.StopRun()
// 	}
// 	this.Data["json"] = map[string]string{
// 		"status": "success",
// 	}
// 	this.ServeJSON()
// }
