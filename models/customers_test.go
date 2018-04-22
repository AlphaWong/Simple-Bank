package models

import (
	"errors"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/Simple-Bank/mocks"
	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"
)

func TestCustomer_CreateSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	customer := new(Customer)
	c := &types.Customer{
		Name:     "name",
		Password: "password",
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	mockOrmer.EXPECT().Insert(c).Return(int64(0), nil).Times(1)
	id, err := customer.Create(c)
	assert.Equal(t, int64(0), id)
	assert.Nil(t, err)
}

func TestCustomer_GetFailByCustomerNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	customer := new(Customer)
	c := &types.Customer{
		Id: int64(0),
	}
	mockOrmer.EXPECT().Read(c).Return(orm.ErrNoRows).Times(1)
	cus, err := customer.Get(int64(0))
	assert.Nil(t, cus)
	assert.EqualError(t, err, utils.ErrorMessageCustomerNotFound)
}

func TestCustomer_GetFailByDbError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	customer := new(Customer)
	c := &types.Customer{
		Id: int64(0),
	}
	mockOrmer.EXPECT().Read(c).Return(errors.New("General error")).Times(1)
	cus, err := customer.Get(int64(0))
	assert.Nil(t, cus)
	assert.EqualError(t, err, "General error")
}

func TestCustomer_Update(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	customer := new(Customer)
	c := &types.Customer{
		Id: int64(0),
	}
	mockOrmer.EXPECT().Update(c).Return(int64(1), nil).Times(1)
	num, err := customer.Update(c)
	assert.Equal(t, int64(1), num)
	assert.Nil(t, err)
}

func TestCustomer_UpdateFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	customer := new(Customer)
	c := &types.Customer{
		Id: int64(0),
	}
	mockOrmer.EXPECT().Update(c).Return(int64(0), errors.New("General error")).Times(1)
	num, err := customer.Update(c)
	assert.Equal(t, int64(0), num)
	assert.EqualError(t, err, "General error")
}
