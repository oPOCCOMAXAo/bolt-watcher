package service

import (
	"bolt-watcher/bolt"
)

var (
	DefaultPayment = bolt.PaymentMethod{
		Type: "default",
		ID:   "cash",
	}

	AllowedCostGroups = []string{
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
