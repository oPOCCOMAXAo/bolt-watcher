package api

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type FindCategoriesOverviewRequest struct {
	From Point
	To   Point
}

type FindCategoriesOverviewResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SearchToken      string           `json:"search_token"`
		SearchCategories []SearchCategory `json:"search_categories"`
	} `json:"data"`
}

func (a *API) FindCategoriesOverview(req FindCategoriesOverviewRequest) FindCategoriesOverviewResponse {
	query := url.Values{
		"version":             []string{"CB.3.58"},
		"deviceId":            []string{"ljxfl"},
		"deviceType":          []string{"web"},
		"device_name":         []string{"Linux x86_64"},
		"device_os_version":   []string{"5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36"},
		"country":             []string{"ua"},
		"language":            []string{"ru"},
		"cache":               []string{""},
		"lat":                 []string{strconv.FormatFloat(req.From.Latitude, 'f', -1, 64)},
		"lng":                 []string{strconv.FormatFloat(req.From.Longitude, 'f', -1, 64)},
		"destination_lat":     []string{strconv.FormatFloat(req.To.Latitude, 'f', -1, 64)},
		"destination_lng":     []string{strconv.FormatFloat(req.To.Longitude, 'f', -1, 64)},
		"gps_accuracy":        []string{"17.216"},
		"payment_method_type": []string{"default"},
		"payment_method_id":   []string{"cash"},
		"interaction_method":  []string{"move_map"},
		"initiated_by":        []string{"user"},
	}

	res := a.client.Get("https://search.bolt.eu/findCategoriesOverview?"+query.Encode(), defaultTimeout, a.headers)

	var resp FindCategoriesOverviewResponse
	_ = json.Unmarshal(res.Response, &resp)

	return resp
}
