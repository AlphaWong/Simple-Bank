package models

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"

	"gitlab.com/Simple-Bank/mocks"
	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"

	"github.com/golang/mock/gomock"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestAccount_CreateSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	a := &Account{}
	c := &types.Customer{
		int64(1),
		"name",
		"password",
		time.Now(),
		time.Now(),
	}
	account := new(types.Account)
	account.Customer = c
	account.Active = true
	mockOrmer.EXPECT().Insert(account).Return(int64(0), nil).Times(1)
	id, err := a.Create(c)
	assert.Equal(t, int64(0), id)
	assert.Nil(t, err)
}

func TestAccount_GetFailByDeactivateAccount(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	a := &Account{}
	account := new(types.Account)
	account.Id = int64(1)
	account.Active = false
	mockOrmer.EXPECT().Read(account).Return(nil).Times(1)
	acc, err := a.Get(int64(1))
	assert.Nil(t, acc)
	assert.EqualError(t, err, utils.ErrorMessageAccountDeactivate)
}

func TestAccount_GetFailByAccountNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	a := &Account{}
	account := new(types.Account)
	account.Id = int64(1)
	account.Active = false
	mockOrmer.EXPECT().Read(account).Return(orm.ErrNoRows).Times(1)
	acc, err := a.Get(int64(1))
	assert.Nil(t, acc)
	assert.EqualError(t, err, utils.ErrorMessageAccountNotFound)
}
func TestAccount_GetFailByGeneralError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	a := &Account{}
	account := new(types.Account)
	account.Id = int64(1)
	account.Active = false
	mockOrmer.EXPECT().Read(account).Return(errors.New("General error")).Times(1)
	acc, err := a.Get(int64(1))
	assert.Nil(t, acc)
	assert.EqualError(t, err, "General error")
}

func TestAccount_UpdateSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	a := &Account{}
	account := new(types.Account)
	mockOrmer.EXPECT().Update(account).Return(int64(1), nil).Times(1)
	num, err := a.Update(account)
	assert.Equal(t, int64(1), num)
	assert.Nil(t, err)
}

func TestAccount_UpdateFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	a := &Account{}
	account := new(types.Account)
	mockOrmer.EXPECT().Update(account).Return(int64(0), errors.New("General error")).Times(1)
	num, err := a.Update(account)
	assert.Equal(t, int64(0), num)
	assert.EqualError(t, err, "General error")
}

func TestAccount_GetCurrentBalanceFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	mockRawSeter := mocks.NewMockRawSeter(mockCtrl)
	expectedBalance := float64(0)
	mockRawSeter.EXPECT().QueryRow(&expectedBalance).Return(errors.New("General error")).Times(1)
	a := &Account{}
	mockOrmer.EXPECT().Raw("SELECT SUM(amount) FROM transaction WHERE account_id = ? ORDER BY id ASC", int64(1)).Return(mockRawSeter).Times(1)
	balance, err := a.GetCurrentBalance(int64(1))
	assert.Equal(t, float64(0), balance)
	assert.EqualError(t, err, "General error")
}

func TestGetPaymentApprovalSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		utils.PaymentApprovalURI,
		httpmock.NewStringResponder(
			200,
			`{"status":"success"}`,
		),
	)

	ok, err := GetPaymentApproval()
	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestGetPaymentApprovalFailByNonSuccessHttpStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		utils.PaymentApprovalURI,
		httpmock.NewStringResponder(
			http.StatusBadRequest,
			`{"status":"I am nonsuccess"}`,
		),
	)

	ok, err := GetPaymentApproval()
	assert.False(t, ok)
	assert.EqualError(t, err, utils.ErrorMessagePaymentApprovalGatewayDeclined)
}
