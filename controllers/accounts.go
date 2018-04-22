package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)
	var customer types.Customer
	err := json.NewDecoder(bytes.NewReader(this.Ctx.Input.RequestBody)).Decode(&customer)
	if nil != err {
		http.Error(this.Ctx.ResponseWriter, fmt.Sprintf("message: %v, footPrint: %v", utils.ErrorMessageInvalidJSON, footPrint), http.StatusBadRequest)
		this.StopRun()
	}
	customerModel := models.NewCustomerModel()
	customerModel.Create(&customer)

	accountModel := models.NewAccountModel()
	accountId, _ := accountModel.Create(&customer)
	account, _ := accountModel.Get(accountId)
	account.Customer = nil

	response := map[string]interface{}{
		"account":   account,
		"footPrint": footPrint,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

// @router /v1/customers/:id/accounts/add [post]
func (this *AccountController) Add() {
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)

	customerId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	customerModel := models.NewCustomerModel()
	customer, err := customerModel.Get(customerId)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusNotFound)
		this.StopRun()
	}
	accountModel := models.NewAccountModel()
	accountId, _ := accountModel.Create(customer)
	account, _ := accountModel.Get(accountId)
	account.Customer = nil

	response := map[string]interface{}{
		"account":   account,
		"footPrint": footPrint,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

// @router /v1/accounts/:id/close [put]
func (this *AccountController) Close() {
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)
	accountModel := models.NewAccountModel()

	accountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	account, err := accountModel.Get(accountId)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	account.Active = false
	_, err = accountModel.Update(account)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}
	account.Customer = nil

	response := map[string]interface{}{
		"account":   account,
		"footPrint": footPrint,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

func transactionCommon(this *AccountController) (*types.Account, *types.TransactionRequest) {
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)

	accountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)
	var transactionRequest types.TransactionRequest
	err := json.NewDecoder(bytes.NewReader(this.Ctx.Input.RequestBody)).Decode(&transactionRequest)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, utils.ErrorMessageInvalidJSON, footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	accountModel := models.NewAccountModel()
	account, err := accountModel.Get(accountId)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusNotFound)
		this.StopRun()
	}
	return account, &transactionRequest
}

// @router /v1/accounts/:id/deposit [put]
func (this *AccountController) Deposit() {
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)
	account, transactionRequest := transactionCommon(this)

	transactionModel := models.NewTransactionModel()
	transaction := models.NewDeposit(account, transactionRequest.Currency, transactionRequest.Amount)
	id, err := transactionModel.Create(transaction)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	transaction, _ = transactionModel.Get(id)

	response := map[string]interface{}{
		"transaction": transaction,
		"footPrint":   footPrint,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

// @router /v1/accounts/:id/withdraw [put]
func (this *AccountController) Withdraw() {
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)
	account, transactionRequest := transactionCommon(this)
	transactionModel := models.NewTransactionModel()
	transaction := models.NewWithdraw(account, transactionRequest.Currency, transactionRequest.Amount)

	if ok, err := models.IsValidTransaction(account.Id, transaction); !ok {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	id, err := transactionModel.Create(transaction)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	transaction, _ = transactionModel.Get(id)

	response := map[string]interface{}{
		"transaction": transaction,
		"footPrint":   footPrint,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

// @router /v1/accounts/:id/balance [get]
func (this *AccountController) Balance() {
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)
	accountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":id"), 10, 64)

	accountModel := models.NewAccountModel()
	_, err := accountModel.Get(accountId)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusNotFound)
		this.StopRun()
	}

	balance, err := accountModel.GetCurrentBalance(accountId)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	response := map[string]interface{}{
		"balance":   balance,
		"footPrint": footPrint,
	}
	this.Data["json"] = response
	this.ServeJSON()
}

// @router /v1/accounts/:from/send/:to [post]
func (this *AccountController) Send() {
	footPrint := this.Ctx.Request.Context().Value(utils.ContextFootPrintKey).(string)

	var transactionRequest types.TransactionRequest
	err := json.NewDecoder(bytes.NewReader(this.Ctx.Input.RequestBody)).Decode(&transactionRequest)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, utils.ErrorMessageInvalidJSON, footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	senderAccountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":from"), 10, 64)
	receiverAccountId, _ := strconv.ParseInt(this.Ctx.Input.Param(":to"), 10, 64)

	accountModel := models.NewAccountModel()
	senderAccount, err := accountModel.Get(senderAccountId)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, utils.ErrorMessageSenderAccountNotFound, footPrint, http.StatusNotFound)
		this.StopRun()
	}
	receiverAccount, err := accountModel.Get(receiverAccountId)
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, utils.ErrorMessageReceiverAccountNotFound, footPrint, http.StatusNotFound)
		this.StopRun()
	}

	var crossCustomerServiceCharge *types.Transaction
	if !models.IsOwnBySameCustomer(senderAccount, receiverAccount) {
		crossCustomerServiceCharge = models.NewCrossCustomerServiceCharge(senderAccount, transactionRequest.Currency, utils.CrossCustomerSendServiceCharge)
		if ok, err := models.IsValidTransaction(senderAccount.Id, crossCustomerServiceCharge); !ok {
			utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
			this.StopRun()
		}
		if ok, err := models.GetPaymentApproval(); !ok {
			utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
			this.StopRun()
		}
	}

	sendTransaction := models.NewSendTransaction(senderAccount, transactionRequest.Currency, transactionRequest.Amount)
	if ok, err := models.IsValidTransaction(senderAccount.Id, sendTransaction); !ok {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	receiveTransaction := models.NewReceiveTransaction(receiverAccount, transactionRequest.Currency, transactionRequest.Amount)

	transactionModel := models.NewTransactionModel()

	var transactions []*types.Transaction
	if nil != crossCustomerServiceCharge {
		transactions, err = transactionModel.Send(crossCustomerServiceCharge, sendTransaction, receiveTransaction)
	} else {
		transactions, err = transactionModel.Send(sendTransaction, receiveTransaction)
	}
	if nil != err {
		utils.SendHttpError(this.Ctx.ResponseWriter, err.Error(), footPrint, http.StatusBadRequest)
		this.StopRun()
	}

	response := map[string]interface{}{
		"transactions": transactions,
		"footPrint":    footPrint,
	}
	this.Data["json"] = response
	this.ServeJSON()
}
