package models

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/Simple-Bank/mocks"
	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"
)

func TestNewDeposit(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	type args struct {
		account  *types.Account
		currency string
		amount   float64
	}
	tests := []struct {
		name string
		args args
		want *types.Transaction
	}{
		{
			"TestNewDeposit pass",
			args{
				account:  acc,
				currency: utils.HongKongDollar,
				amount:   float64(100),
			},
			&types.Transaction{
				int64(0),
				acc,
				utils.Deposit,
				utils.HongKongDollar,
				float64(100),
				"",
				time.Time{},
				time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeposit(tt.args.account, tt.args.currency, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeposit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithdraw(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	type args struct {
		account  *types.Account
		currency string
		amount   float64
	}
	tests := []struct {
		name string
		args args
		want *types.Transaction
	}{
		{
			"TestNewWithdraw pass",
			args{
				account:  acc,
				currency: utils.HongKongDollar,
				amount:   float64(100),
			},
			&types.Transaction{
				int64(0),
				acc,
				utils.Withdraw,
				utils.HongKongDollar,
				float64(100) * -1,
				"",
				time.Time{},
				time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithdraw(tt.args.account, tt.args.currency, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCrossCustomerServiceCharge(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	type args struct {
		account  *types.Account
		currency string
		amount   float64
	}
	tests := []struct {
		name string
		args args
		want *types.Transaction
	}{
		{
			"TestNewCrossCustomerServiceCharge pass",
			args{
				account:  acc,
				currency: utils.HongKongDollar,
				amount:   float64(100),
			},
			&types.Transaction{
				int64(0),
				acc,
				utils.CrossCustomerServiceCharge,
				utils.HongKongDollar,
				float64(100) * -1,
				"",
				time.Time{},
				time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCrossCustomerServiceCharge(tt.args.account, tt.args.currency, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCrossCustomerServiceCharge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSendTransaction(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	type args struct {
		account  *types.Account
		currency string
		amount   float64
	}
	tests := []struct {
		name string
		args args
		want *types.Transaction
	}{
		{
			"TestNewSendTransaction pass",
			args{
				account:  acc,
				currency: utils.HongKongDollar,
				amount:   float64(100),
			},
			&types.Transaction{
				int64(0),
				acc,
				utils.Transfer,
				utils.HongKongDollar,
				float64(100) * -1,
				"Account 1 send money HKD 100",
				time.Time{},
				time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSendTransaction(tt.args.account, tt.args.currency, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSendTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewReceiveTransaction(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	type args struct {
		account  *types.Account
		currency string
		amount   float64
	}
	tests := []struct {
		name string
		args args
		want *types.Transaction
	}{
		{
			"TestNewReceiveTransaction pass",
			args{
				account:  acc,
				currency: utils.HongKongDollar,
				amount:   float64(100),
			},
			&types.Transaction{
				int64(0),
				acc,
				utils.Transfer,
				utils.HongKongDollar,
				float64(100),
				"Account 1 receive money HKD 100",
				time.Time{},
				time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReceiveTransaction(tt.args.account, tt.args.currency, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReceiveTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_Create(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	transaction := new(Transaction)
	trans := &types.Transaction{
		Account:  acc,
		Type:     utils.Transfer,
		Currency: utils.HongKongDollar,
		Amount:   float64(100),
		Remark:   "Account 1 receive money HKD 100",
		Created:  time.Time{},
		Updated:  time.Time{},
	}
	mockOrmer.EXPECT().Begin().Return(nil).Times(1)
	mockOrmer.EXPECT().Commit().Return(nil).Times(1)
	mockOrmer.EXPECT().Insert(trans).Return(int64(0), nil).Times(1)
	id, err := transaction.Create(trans)
	assert.Equal(t, int64(0), id)
	assert.Nil(t, err)
}

func TestTransaction_CreateFailByDbGeneralError(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	transaction := new(Transaction)
	trans := &types.Transaction{
		Account:  acc,
		Type:     utils.Transfer,
		Currency: utils.HongKongDollar,
		Amount:   float64(100),
		Remark:   "Account 1 receive money HKD 100",
		Created:  time.Time{},
		Updated:  time.Time{},
	}
	mockOrmer.EXPECT().Begin().Return(nil).Times(1)
	mockOrmer.EXPECT().Rollback().Return(nil).Times(1)
	mockOrmer.EXPECT().Commit().Return(nil).Times(0)
	mockOrmer.EXPECT().Insert(trans).Return(int64(0), errors.New("General error")).Times(1)
	id, err := transaction.Create(trans)
	assert.Equal(t, int64(0), id)
	assert.EqualError(t, err, "General error")
}

func TestTransaction_SendSuccess(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	transaction := new(Transaction)
	transs := []*types.Transaction{
		{
			Account:  acc,
			Type:     utils.Transfer,
			Currency: utils.HongKongDollar,
			Amount:   float64(100),
			Remark:   "Account 1 send money HKD 100",
			Created:  time.Time{},
			Updated:  time.Time{},
		}, {
			Account:  acc,
			Type:     utils.Transfer,
			Currency: utils.HongKongDollar,
			Amount:   float64(100),
			Remark:   "Account 1 receive money HKD 100",
			Created:  time.Time{},
			Updated:  time.Time{},
		}, {
			Account:  acc,
			Type:     utils.Transfer,
			Currency: utils.HongKongDollar,
			Amount:   float64(100),
			Remark:   "",
			Created:  time.Time{},
			Updated:  time.Time{},
		},
	}
	mockOrmer.EXPECT().Begin().Return(nil).Times(1)
	// mockOrmer.EXPECT().Rollback().Return(nil).Times(1)
	mockOrmer.EXPECT().Commit().Return(nil).Times(1)
	mockOrmer.EXPECT().Insert(transs[0]).Return(int64(0), nil).Times(1)
	mockOrmer.EXPECT().Insert(transs[1]).Return(int64(0), nil).Times(1)
	mockOrmer.EXPECT().Insert(transs[2]).Return(int64(0), nil).Times(1)
	ts, err := transaction.Send(transs[0], transs[1], transs[2])
	assert.Equal(t, transs, ts)
	assert.Nil(t, err)
}

func TestTransaction_SendFailByDbGeneralError(t *testing.T) {
	acc := &types.Account{
		int64(1),
		true,
		&types.Customer{
			int64(1),
			"name",
			"password",
			time.Now(),
			time.Now(),
		},
		time.Now(),
		time.Now(),
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	transaction := new(Transaction)
	transs := []*types.Transaction{
		{
			Account:  acc,
			Type:     utils.Transfer,
			Currency: utils.HongKongDollar,
			Amount:   float64(100),
			Remark:   "Account 1 send money HKD 100",
			Created:  time.Time{},
			Updated:  time.Time{},
		}, {
			Account:  acc,
			Type:     utils.Transfer,
			Currency: utils.HongKongDollar,
			Amount:   float64(100),
			Remark:   "Account 1 receive money HKD 100",
			Created:  time.Time{},
			Updated:  time.Time{},
		}, {
			Account:  acc,
			Type:     utils.Transfer,
			Currency: utils.HongKongDollar,
			Amount:   float64(100),
			Remark:   "",
			Created:  time.Time{},
			Updated:  time.Time{},
		},
	}
	mockOrmer.EXPECT().Begin().Return(nil).Times(1)
	mockOrmer.EXPECT().Rollback().Return(nil).Times(1)
	mockOrmer.EXPECT().Commit().Return(nil).Times(0)
	mockOrmer.EXPECT().Insert(transs[0]).Return(int64(0), nil).Times(1)
	mockOrmer.EXPECT().Insert(transs[1]).Return(int64(0), errors.New("General error")).Times(1)
	mockOrmer.EXPECT().Insert(transs[2]).Return(int64(0), nil).Times(0)
	ts, err := transaction.Send(transs[0], transs[1], transs[2])
	assert.Equal(t, make([]*types.Transaction, 0), ts)
	assert.EqualError(t, err, "General error")
}

func TestTransaction_GetFailByNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOrmer := mocks.NewMockOrmer(mockCtrl)
	utils.OrmInstance = mockOrmer
	transaction := new(Transaction)
	trans := &types.Transaction{
		Id: int64(0),
	}
	mockOrmer.EXPECT().Read(trans).Return(orm.ErrNoRows).Times(1)
	cus, err := transaction.Get(int64(0))
	assert.Nil(t, cus)
	assert.EqualError(t, err, orm.ErrNoRows.Error())
}
