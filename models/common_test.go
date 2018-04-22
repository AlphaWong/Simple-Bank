package models

import (
	"testing"

	"gitlab.com/Simple-Bank/types"
)

func TestIsValidCurrency(t *testing.T) {
	type args struct {
		transaction *types.Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"IsValidCurrency pass in HKD",
			args{
				&types.Transaction{
					Currency: "HKD",
				},
			},
			true,
			false,
		}, {
			"IsValidCurrency fail in BTC",
			args{
				&types.Transaction{
					Currency: "BTC",
				},
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsValidCurrency(tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValidCurrency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValidCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPositiveNumberAfterTransaction(t *testing.T) {
	type args struct {
		transaction *types.Transaction
		balance     float64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"IsPositiveNumberAfterTransaction pass",
			args{
				&types.Transaction{
					Amount: float64(9),
				},
				float64(999),
			},
			true,
			false,
		}, {
			"IsPositiveNumberAfterTransaction fail in balance not enough",
			args{
				&types.Transaction{
					Amount: float64(99999),
				},
				float64(9),
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPositiveNumberAfterTransaction(tt.args.transaction, tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPositiveNumberAfterTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPositiveNumberAfterTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOwnBySameCustomer(t *testing.T) {
	type args struct {
		senderAccount   *types.Account
		receiverAccount *types.Account
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"TestIsOwnBySameCustomer pass",
			args{
				&types.Account{
					Customer: &types.Customer{
						Id: int64(1),
					},
				},
				&types.Account{
					Customer: &types.Customer{
						Id: int64(1),
					},
				},
			},
			true,
		}, {
			"TestIsOwnBySameCustomer fail",
			args{
				&types.Account{
					Customer: &types.Customer{
						Id: int64(1),
					},
				},
				&types.Account{
					Customer: &types.Customer{
						Id: int64(2),
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOwnBySameCustomer(tt.args.senderAccount, tt.args.receiverAccount); got != tt.want {
				t.Errorf("IsOwnBySameCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
