package types

import "time"

type (
	Customer struct {
		Id       int64     `json:"id" orm:"auto"`
		Name     string    `json:"name"`
		Password string    `json:"password"`
		Created  time.Time `orm:"auto_now_add;type(datetime)"`
		Updated  time.Time `orm:"auto_now;type(datetime)"`
	}

	Account struct {
		Id       int64     `json:"id" orm:"auto"`
		Active   bool      `json:"active"`
		Customer *Customer `json:"customer,omitempty" orm:"rel(fk)"`
		Created  time.Time `orm:"auto_now_add;type(datetime)"`
		Updated  time.Time `orm:"auto_now;type(datetime)"`
	}

	Transaction struct {
		Id       int64     `json:"id" orm:"auto"`
		Amount   float64   `json:"amount"`
		Currency string    `json:"currency"`
		Type     string    `json:"type"`
		From     *Account  `json:"from" orm:"rel(fk)"`
		To       *Account  `json:"to" orm:"rel(fk)"`
		Created  time.Time `orm:"auto_now_add;type(datetime)"`
		Updated  time.Time `orm:"auto_now;type(datetime)"`
	}

	Operation struct {
		Account  *Account  `json:"account" orm:"rel(fk)"`
		Id       int64     `json:"id" orm:"auto"`
		Amount   float64   `json:"amount"`
		Currency string    `json:"currency"`
		Type     string    `json:"type"`
		Created  time.Time `orm:"auto_now_add;type(datetime)"`
		Updated  time.Time `orm:"auto_now;type(datetime)"`
	}

	OperationRequest struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	}
)
