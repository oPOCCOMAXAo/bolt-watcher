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

type Viewport struct {
	NorthEast Point `json:"north_east"`
	SouthWest Point `json:"south_west"`
}

type Vehicle struct {
	ID      string  `json:"id"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lng"`
	Bearing float64 `json:"bearing"`
	IconID  string  `json:"icon_id"`
}
