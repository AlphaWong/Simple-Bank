package models

import (
	"gitlab.com/Simple-Bank/types"
	"gitlab.com/Simple-Bank/utils"

	"golang.org/x/crypto/bcrypt"
)

var c *Customer

type Customer struct{}

func NewCustomerModel() *Customer {
	if nil == c {
		c = new(Customer)
	}
	return c
}

func (*Customer) Create(customer *types.Customer) (int64, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	customer.Password = string(hash)
	return utils.OrmInstance.Insert(customer)
}

func (*Customer) Get(id int64) (*types.Customer, error) {
	customer := new(types.Customer)
	customer.Id = id

	err := utils.OrmInstance.Read(customer)
	if nil != err {
		return nil, err
	}
	return customer, nil
}

func (*Customer) Update(customer *types.Customer) (int64, error) {
	return utils.OrmInstance.Update(customer)
}
