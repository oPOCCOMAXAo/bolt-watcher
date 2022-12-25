package api

import "errors"

var (
	DefaultPayment = PaymentMethod{
		Type: "default",
		ID:   "cash",
	}

	AllowedGroups = []string{
		"standard",
		"comfort",
		"electric",
		"animal_friendly",
		"premium",
		"child_seat",
		"xl",
		"economy",
	}
)

var (
	ErrFailed = errors.New("failed")
)
