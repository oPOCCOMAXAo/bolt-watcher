package bolt

import (
	"context"
	"net/http"

	"github.com/opoccomaxao-go/xhttp"
	"github.com/opoccomaxao-go/xhttp/request"
	"github.com/opoccomaxao-go/xhttp/response"
	"github.com/pkg/errors"
)

type RideOptionsRequest struct {
	PickupStop       Point         `json:"pickup_stop"`
	DestinationStops []Point       `json:"destination_stops"`
	PaymentMethod    PaymentMethod `json:"payment_method"`
	Timezone         string        `json:"timezone"`
	CampaignCode     struct{}      `json:"campaign_code"`
}

type RideOptionsResponse struct {
	CommonResponse
	Data struct {
		CategoriesList []Category                            `json:"categories_list"`
		RideOptions    map[OrderSystem]OrderSystemCategories `json:"ride_options"`
	} `json:"data"`
}

type OrderSystemCategories struct {
	Categories map[string]RideOption `json:"categories"`
}

type RideOption struct {
	ETAInfo struct {
		PickupETAStr string `json:"pickup_eta_str"`
	} `json:"eta_info"`
	Group string `json:"group"`
	Price struct {
		SurgeStr       string `json:"surge_str"`
		ActualStr      string `json:"actual_str"` // with discount
		FirstLineHTML  string `json:"first_line_html"`
		SecondLineHTML string `json:"second_line_html"`
	} `json:"price"`
}

func (api *Client) GetRideOptions(
	ctx context.Context,
	req *RideOptionsRequest,
) (*RideOptionsResponse, error) {
	var res RideOptionsResponse

	err := api.client.Do(ctx, xhttp.Params{
		Method:       http.MethodPost,
		URL:          "https://user.live.boltsvc.net/rides/search/getRideOptions",
		Headers:      api.headers,
		Query:        api.query,
		RequestBody:  request.JSON(req),
		ResponseBody: response.JSON(&res),
	})
	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.WithMessage(ErrFailed, res.Message)
	}

	return &res, nil
}
