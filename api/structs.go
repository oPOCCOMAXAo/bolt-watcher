package api

type Point struct {
	Latitude  float64
	Longitude float64
}

type SearchCategory struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	UpfrontPriceStr string  `json:"upfront_price_str"`
	SurgeMultiplier float64 `json:"surge_multiplier"`
	PickupEta       int     `json:"pickup_eta"`
}
