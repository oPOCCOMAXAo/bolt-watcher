package bolt

type CommonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PaymentMethod struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}

type Category struct {
	OrderSystem OrderSystem `json:"order_system"`
	CategoryID  string      `json:"category_id"`
}
