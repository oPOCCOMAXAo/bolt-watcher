package api

import (
	"bolt-watcher/utils"
	"context"
	"net/url"

	"github.com/opoccomaxao-go/generic-collection/slice"
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
	DefaultResponse
	Data struct {
		CategoriesList []Category                            `json:"categories_list"`
		RideOptions    map[OrderSystem]OrderSystemRideOption `json:"ride_options"`
	} `json:"data"`
}

type OrderSystemRideOption struct {
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

type RideOptionParsed struct {
	ETA        int64   // minutes.
	Price      int64   // hryvna.
	Multiplier float64 // X.
}

func (api *API) GetRideOptions(ctx context.Context, route Route) ([]RideOptionParsed, error) {
	var res RideOptionsResponse

	err := api.request(
		ctx,
		"https://user.live.boltsvc.net/rides/search/getRideOptions",
		url.Values{
			"version":           []string{"CB.3.74"},
			"deviceId":          []string{"1"},
			"deviceType":        []string{"web"},
			"device_name":       []string{"Linux x86_64"},
			"device_os_version": []string{"5.0"},
			"country":           []string{"ua"},
			"language":          []string{"en-US"},
		},
		&RideOptionsRequest{
			PickupStop:       route[0],
			DestinationStops: route[1:],
			PaymentMethod:    DefaultPayment,
		},
		&res,
		false,
	)

	if err != nil {
		return nil, err
	}

	if res.Code != 0 {
		return nil, errors.WithMessage(ErrFailed, res.Message)
	}

	resList := []RideOptionParsed{}

	taxis := res.Data.RideOptions[OrderSystemTaxi].Categories

	if taxis != nil {
		for _, category := range res.Data.CategoriesList {
			if category.OrderSystem != OrderSystemTaxi {
				continue
			}

			option, ok := taxis[category.CategoryID]
			if !ok {
				continue
			}

			if slice.IndexOf(AllowedGroups, option.Group) == -1 {
				continue
			}

			resList = append(resList, RideOptionParsed{
				ETA: utils.TryParseInt(option.ETAInfo.PickupETAStr),
				Price: slice.FirstNonEmpty(
					utils.TryParseInt(option.Price.SecondLineHTML),
					utils.TryParseInt(option.Price.FirstLineHTML),
					utils.TryParseInt(option.Price.ActualStr),
				),
				Multiplier: utils.TryParseFloat(option.Price.SurgeStr),
			})
		}
	}

	return resList, nil
}
