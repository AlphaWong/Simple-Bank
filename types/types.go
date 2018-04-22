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
		Sender   *Account  `json:"sender" orm:"rel(fk)"`
		Receiver *Account  `json:"receiver" orm:"null;rel(fk)"`
		Type     string    `json:"type"`
		Currency string    `json:"currency"`
		Amount   float64   `json:"amount"`
		Created  time.Time `orm:"auto_now_add;type(datetime)"`
		Updated  time.Time `orm:"auto_now;type(datetime)"`
	}

	TransactionRequest struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	}
)
