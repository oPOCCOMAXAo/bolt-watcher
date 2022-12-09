package api

import "errors"

var (
	DefaultPayment = PaymentMethod{
		Type: "default",
		ID:   "cash",
	}
)

var (
	ErrFailed = errors.New("failed")
)
